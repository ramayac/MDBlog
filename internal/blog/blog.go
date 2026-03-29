package blog

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/markdown"
)

// Version info injected at build time via -ldflags. Populated by cmd packages.
var (
	BuildVersion = ""
	BuildCommit  = ""
	BuildDate    = ""
)

// Post represents a single blog post with all its metadata and rendered content.
type Post struct {
	Filename     string
	Slug         string
	Title        string
	Date         string
	Content      string // rendered HTML, populated by getPostBySlug only
	Excerpt      string
	FrontMatter  markdown.FrontMatter
	CategorySlug string
	Category     *CategoryInfo // resolved category, may be nil
}

// CategoryInfo is the displayable info for a category, used in templates.
type CategoryInfo struct {
	BlogName      string
	HeaderContent string
	Folder        string
	Slug          string
	Count         int
}

// Pagination holds data for building prev/next page links in templates.
type Pagination struct {
	Current int
	Total   int
	HasNext bool
	HasPrev bool
	Next    int
	Prev    int
}

// PostList is returned by GetPosts and SearchPosts.
type PostList struct {
	Posts         []Post
	Pagination    Pagination
	TotalMatches  int // populated by SearchPosts
}

// MenuLink is a navigation link item.
type MenuLink struct {
	Label string
	URL   string
}

// VersionInfo holds build version metadata.
type VersionInfo struct {
	Version string
	Commit  string
	Date    string
}

// indexPost mirrors the JSON fields in posts.index.json.
type indexPost struct {
	Slug         string `json:"slug"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	Author       string `json:"author"`
	Tags         string `json:"tags"`
	Description  string `json:"description"`
	Excerpt      string `json:"excerpt"`
	CategorySlug string `json:"category_slug"`
	SourcePath   string `json:"source_path"`
	Filename     string `json:"filename"`
}

var (
	slugCleanupRegex  = regexp.MustCompile(`[^a-z0-9]+`)
	datePrefixRegex   = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}-`)
	whitespaceRegex   = regexp.MustCompile(`\s+`)
)

// Blog is the core service for reading and serving blog posts.
type Blog struct {
	cfg             *config.Config
	categoriesCache map[string]*CategoryInfo
}

// New creates a Blog backed by the given config.
func New(cfg *config.Config) *Blog {
	return &Blog{cfg: cfg}
}

// ─────────────────────────────────────────────────────────────────────────────
// Public API
// ─────────────────────────────────────────────────────────────────────────────

// GetPosts returns a page of posts, optionally filtered to a category.
func (b *Blog) GetPosts(page int, categorySlug string) PostList {
	if page < 1 {
		page = 1
	}
	index := b.loadPostIndex()
	if index != nil {
		return b.getPostsFromIndex(index, page, categorySlug)
	}
	// Fallback: filesystem scan
	return b.scanAndPaginate(page, categorySlug)
}

// SearchPosts searches the post index for posts matching query.
func (b *Blog) SearchPosts(query string, page int) PostList {
	if page < 1 {
		page = 1
	}
	index := b.loadPostIndex()
	if index == nil {
		return PostList{Pagination: b.buildPagination(1, 1)}
	}
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return PostList{Pagination: b.buildPagination(1, 1)}
	}
	var results []Post
	for _, ip := range index {
		haystack := strings.ToLower(ip.Title + " " + ip.Excerpt + " " + ip.Tags)
		if strings.Contains(haystack, q) {
			results = append(results, indexPostToPost(ip))
		}
	}
	sortPostsByDate(results)
	total := len(results)
	perPage := b.cfg.PostsPerPage
	totalPages := max1(int(math.Ceil(float64(total) / float64(perPage))))
	offset := (page - 1) * perPage
	if offset > total {
		offset = total
	}
	end := offset + perPage
	if end > total {
		end = total
	}
	posts := results[offset:end]
	b.attachCategories(posts)
	return PostList{
		Posts:        posts,
		Pagination:   b.buildPagination(page, totalPages),
		TotalMatches: total,
	}
}

// GetPostBySlug returns a fully-rendered post by its slug, optionally scoped
// to a category. Returns nil when not found.
func (b *Blog) GetPostBySlug(slug, categorySlug string) *Post {
	// Path traversal guard — mirrors Blog.php validation
	if strings.Contains(slug, "..") || strings.Contains(slug, "/") || strings.Contains(slug, `\`) {
		return nil
	}

	postsDir := b.cfg.PostsDir
	if categorySlug != "" {
		cat := b.GetCategoryBySlug(categorySlug)
		if cat == nil {
			return nil
		}
		postsDir = filepath.Join(b.cfg.PostsDir, cat.Folder)
	}

	fullPath := filepath.Join(postsDir, slug+".md")
	if !fileExists(fullPath) {
		// Index fallback routing: find the post in another category folder
		if categorySlug == "" {
			if resolved := b.resolveSlugViaIndex(slug); resolved != "" {
				fullPath = resolved
				// Extract category from resolved path
				rel, _ := filepath.Rel(b.cfg.PostsDir, fullPath)
				parts := strings.SplitN(rel, string(filepath.Separator), 2)
				if len(parts) == 2 {
					categorySlug = parts[0]
				}
			}
		}
		if !fileExists(fullPath) {
			return nil
		}
	}

	post, err := b.parsePost(fullPath)
	if err != nil {
		log.Printf("blog: parsePost %s: %v", fullPath, err)
		return nil
	}
	post.CategorySlug = categorySlug
	if categorySlug != "" {
		post.Category = b.GetCategoryBySlug(categorySlug)
	}
	return post
}

// categoryMenuEntry is used internally to sort category nav links.
type categoryMenuEntry struct {
	slug  string
	cat   config.Category
}

// GetMenu returns the ordered list of navigation links.
// Static [[menu_links]] come first (in config order), followed by categories
// with menu=true sorted by menu_order (ascending), then slug for ties.
// The theme toggle and search links are rendered directly in the layout template.
func (b *Blog) GetMenu() []MenuLink {
	var links []MenuLink
	for _, ml := range b.cfg.MenuLinks {
		links = append(links, MenuLink{Label: ml.Label, URL: ml.URL})
	}

	// Collect and sort category entries for deterministic order.
	var catEntries []categoryMenuEntry
	for slug, cat := range b.cfg.Categories {
		if cat.Menu {
			catEntries = append(catEntries, categoryMenuEntry{slug: slug, cat: cat})
		}
	}
	sort.Slice(catEntries, func(i, j int) bool {
		if catEntries[i].cat.MenuOrder != catEntries[j].cat.MenuOrder {
			return catEntries[i].cat.MenuOrder < catEntries[j].cat.MenuOrder
		}
		return catEntries[i].slug < catEntries[j].slug
	})
	for _, e := range catEntries {
		links = append(links, MenuLink{
			Label: e.cat.BlogName,
			URL:   "/?category=" + e.slug,
		})
	}
	return links
}

// GetCategories returns all categories that have at least one post.
func (b *Blog) GetCategories() map[string]*CategoryInfo {
	if b.categoriesCache != nil {
		return b.categoriesCache
	}
	cats := make(map[string]*CategoryInfo)
	for slug, cat := range b.cfg.Categories {
		dir := filepath.Join(b.cfg.PostsDir, cat.Folder)
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		count := 0
		for _, e := range entries {
			if !e.IsDir() && strings.EqualFold(filepath.Ext(e.Name()), ".md") {
				count++
			}
		}
		if count > 0 {
			cats[slug] = &CategoryInfo{
				BlogName:      cat.BlogName,
				HeaderContent: cat.HeaderContent,
				Folder:        cat.Folder,
				Slug:          slug,
				Count:         count,
			}
		}
	}
	b.categoriesCache = cats
	return cats
}

// GetCategoryBySlug returns category info by slug, or nil if not found.
func (b *Blog) GetCategoryBySlug(slug string) *CategoryInfo {
	return b.GetCategories()[slug]
}

// ParseMarkdown renders a Markdown string to HTML.
func (b *Blog) ParseMarkdown(content string) string {
	return markdown.Parse(content).HTML
}

// GetVersionInfo returns build version metadata.
func (b *Blog) GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version: BuildVersion,
		Commit:  BuildCommit,
		Date:    BuildDate,
	}
}

// GetConfig returns the underlying config, useful for templates.
func (b *Blog) GetConfig() *config.Config {
	return b.cfg
}

// ─────────────────────────────────────────────────────────────────────────────
// Index loading
// ─────────────────────────────────────────────────────────────────────────────

func (b *Blog) loadPostIndex() []indexPost {
	path := b.cfg.PostIndexFile
	if path == "" || !fileExists(path) {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("blog: loadPostIndex read error: %v", err)
		return nil
	}
	var posts []indexPost
	if err := json.Unmarshal(data, &posts); err != nil {
		log.Printf("blog: loadPostIndex parse error: %v — falling back to filesystem scan", err)
		return nil
	}
	return posts
}

func (b *Blog) getPostsFromIndex(index []indexPost, page int, categorySlug string) PostList {
	var filtered []indexPost
	if categorySlug != "" {
		for _, ip := range index {
			if ip.CategorySlug == categorySlug {
				filtered = append(filtered, ip)
			}
		}
	} else {
		// Aggregated: uncategorized + categories with index=true
		indexedCats := b.indexedCategorySlugs()
		for _, ip := range index {
			if ip.CategorySlug == "" {
				if b.cfg.ShowUncategorized {
					filtered = append(filtered, ip)
				}
			} else if indexedCats[ip.CategorySlug] {
				filtered = append(filtered, ip)
			}
		}
	}

	total := len(filtered)
	perPage := b.cfg.PostsPerPage
	totalPages := max1(int(math.Ceil(float64(total) / float64(perPage))))
	offset := (page - 1) * perPage
	if offset > total {
		offset = total
	}
	end := offset + perPage
	if end > total {
		end = total
	}

	posts := make([]Post, 0, end-offset)
	for _, ip := range filtered[offset:end] {
		posts = append(posts, indexPostToPost(ip))
	}
	b.attachCategories(posts)
	return PostList{
		Posts:      posts,
		Pagination: b.buildPagination(page, totalPages),
	}
}

func (b *Blog) indexedCategorySlugs() map[string]bool {
	m := make(map[string]bool)
	for slug, cat := range b.cfg.Categories {
		if cat.Index {
			m[slug] = true
		}
	}
	return m
}

// ─────────────────────────────────────────────────────────────────────────────
// Filesystem fallback scan
// ─────────────────────────────────────────────────────────────────────────────

func (b *Blog) scanAndPaginate(page int, categorySlug string) PostList {
	var posts []Post
	if categorySlug != "" {
		cat := b.GetCategoryBySlug(categorySlug)
		if cat == nil {
			return PostList{Pagination: b.buildPagination(1, 1)}
		}
		posts = b.scanFolder(filepath.Join(b.cfg.PostsDir, cat.Folder), categorySlug)
	} else {
		if b.cfg.ShowUncategorized {
			posts = append(posts, b.scanFolder(b.cfg.PostsDir, "")...)
		}
		for slug, cat := range b.cfg.Categories {
			if cat.Index {
				dir := filepath.Join(b.cfg.PostsDir, cat.Folder)
				posts = append(posts, b.scanFolder(dir, slug)...)
			}
		}
	}
	sortPostsByDate(posts)
	total := len(posts)
	perPage := b.cfg.PostsPerPage
	totalPages := max1(int(math.Ceil(float64(total) / float64(perPage))))
	offset := (page - 1) * perPage
	if offset > total {
		offset = total
	}
	end := offset + perPage
	if end > total {
		end = total
	}
	return PostList{
		Posts:      posts[offset:end],
		Pagination: b.buildPagination(page, totalPages),
	}
}

func (b *Blog) scanFolder(dir, categorySlug string) []Post {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var posts []Post
	for _, e := range entries {
		if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), ".md") {
			continue
		}
		post, err := b.parsePost(filepath.Join(dir, e.Name()))
		if err != nil {
			log.Printf("blog: scanFolder skip %s: %v", e.Name(), err)
			continue
		}
		post.CategorySlug = categorySlug
		if categorySlug != "" {
			post.Category = b.GetCategoryBySlug(categorySlug)
		}
		posts = append(posts, *post)
	}
	return posts
}

// ─────────────────────────────────────────────────────────────────────────────
// Post parsing
// ─────────────────────────────────────────────────────────────────────────────

func (b *Blog) parsePost(fullPath string) (*Post, error) {
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", fullPath, err)
	}
	doc := markdown.Parse(string(data))
	filename := filepath.Base(fullPath)
	slug := generateSlug(filename)

	title := doc.FrontMatter.Title
	if title == "" {
		title = titleFromFilename(filename)
	}

	date := doc.FrontMatter.Date
	if date == "" {
		info, err := os.Stat(fullPath)
		if err == nil {
			date = info.ModTime().Format(b.cfg.DateFormat)
		} else {
			date = time.Now().Format(b.cfg.DateFormat)
		}
	}

	return &Post{
		Filename:    filename,
		Slug:        slug,
		Title:       title,
		Date:        date,
		Content:     doc.HTML,
		Excerpt:     generateExcerpt(doc.HTML, b.cfg.ExcerptLength),
		FrontMatter: doc.FrontMatter,
	}, nil
}

// resolveSlugViaIndex searches the post index for a slug and returns the full
// filesystem path when found, or empty string when not found.
func (b *Blog) resolveSlugViaIndex(slug string) string {
	index := b.loadPostIndex()
	if index == nil {
		return ""
	}
	for _, ip := range index {
		if ip.Slug == slug {
			catSlug := ip.CategorySlug
			if catSlug == "" {
				continue
			}
			cat, ok := b.cfg.Categories[catSlug]
			if !ok {
				continue
			}
			candidate := filepath.Join(b.cfg.PostsDir, cat.Folder, slug+".md")
			if fileExists(candidate) {
				return candidate
			}
		}
	}
	return ""
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func indexPostToPost(ip indexPost) Post {
	return Post{
		Filename:     ip.Filename,
		Slug:         ip.Slug,
		Title:        ip.Title,
		Date:         ip.Date,
		Excerpt:      ip.Excerpt,
		CategorySlug: ip.CategorySlug,
		FrontMatter: markdown.FrontMatter{
			Title:       ip.Title,
			Date:        ip.Date,
			Author:      ip.Author,
			Tags:        ip.Tags,
			Description: ip.Description,
		},
	}
}

func (b *Blog) attachCategories(posts []Post) {
	for i := range posts {
		if posts[i].Category == nil && posts[i].CategorySlug != "" {
			posts[i].Category = b.GetCategoryBySlug(posts[i].CategorySlug)
		}
	}
}

func sortPostsByDate(posts []Post) {
	sort.Slice(posts, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02", posts[i].Date)
		tj, _ := time.Parse("2006-01-02", posts[j].Date)
		return tj.Before(ti)
	})
}

func (b *Blog) buildPagination(current, total int) Pagination {
	return Pagination{
		Current: current,
		Total:   total,
		HasNext: current < total,
		HasPrev: current > 1,
		Next:    current + 1,
		Prev:    current - 1,
	}
}

func generateSlug(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	slug := strings.ToLower(name)
	slug = slugCleanupRegex.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

func titleFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	name = datePrefixRegex.ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	return strings.Map(func(r rune) rune {
		return unicode.ToTitle(r)
	}, name[:1]) + name[1:]
}

func generateExcerpt(htmlContent string, length int) string {
	// Strip all HTML tags
	text := regexp.MustCompile(`<[^>]+>`).ReplaceAllString(htmlContent, "")
	text = whitespaceRegex.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	if len(text) <= length {
		return text
	}
	text = text[:length]
	if idx := strings.LastIndex(text, " "); idx > 0 {
		text = text[:idx]
	}
	return text + "..."
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func max1(n int) int {
	if n < 1 {
		return 1
	}
	return n
}
