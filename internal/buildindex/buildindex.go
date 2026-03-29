package buildindex

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/markdown"
)

// IndexPost is the structure written into posts.index.json.
type IndexPost struct {
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
	slugCleanup = regexp.MustCompile(`[^a-z0-9]+`)
	datePrefix  = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}-`)
)

// Build scans all post directories, extracts metadata, and writes
// posts.index.json to the path specified in cfg.PostIndexFile.
func Build(cfg *config.Config) error {
	var entries []IndexPost

	// Uncategorized posts (root posts dir)
	if cfg.ShowUncategorized {
		root, err := scanFolder(cfg.PostsDir, "", cfg.DateFormat)
		if err != nil {
			log.Printf("buildindex: scan root: %v", err)
		}
		entries = append(entries, root...)
	}

	// Categorized posts
	for slug, cat := range cfg.Categories {
		folder := cat.Folder
		if folder == "" {
			folder = slug
		}
		dir := filepath.Join(cfg.PostsDir, folder)
		catEntries, err := scanFolder(dir, slug, cfg.DateFormat)
		if err != nil {
			log.Printf("buildindex: scan %s: %v", dir, err)
			continue
		}
		entries = append(entries, catEntries...)
	}

	// Sort by date descending (newest first) — pre-sorted so runtime skips sorting
	sort.Slice(entries, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02", entries[i].Date)
		tj, _ := time.Parse("2006-01-02", entries[j].Date)
		return tj.Before(ti)
	})

	// Atomic write: tmp → rename
	indexPath := cfg.PostIndexFile
	if err := os.MkdirAll(filepath.Dir(indexPath), 0755); err != nil {
		return fmt.Errorf("buildindex: mkdir %s: %w", filepath.Dir(indexPath), err)
	}
	tmp := indexPath + ".tmp"
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("buildindex: marshal: %w", err)
	}
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return fmt.Errorf("buildindex: write tmp: %w", err)
	}
	if err := os.Rename(tmp, indexPath); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("buildindex: rename: %w", err)
	}

	fmt.Printf("buildindex: wrote %d posts to %s\n", len(entries), indexPath)
	return nil
}

// scanFolder reads all .md files in dir and extracts post metadata.
func scanFolder(dir, categorySlug, dateFormat string) ([]IndexPost, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}

	var posts []IndexPost
	for _, e := range entries {
		if e.IsDir() || !strings.EqualFold(filepath.Ext(e.Name()), ".md") {
			continue
		}
		// Skip the optional landing-page blurb
		if e.Name() == "index.md" && categorySlug == "" {
			continue
		}

		fullPath := filepath.Join(dir, e.Name())
		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Printf("buildindex: cannot read %s, skipping: %v", fullPath, err)
			continue
		}

		// Repair invalid UTF-8 before parsing
		if !utf8.ValidString(string(content)) {
			content = []byte(strings.ToValidUTF8(string(content), "\uFFFD"))
		}

		meta := markdown.ParseMetaOnly(string(content))
		fm := meta.FrontMatter

		slug := slugFromFilename(e.Name())
		title := fm.Title
		if title == "" {
			title = titleFromFilename(e.Name())
		}

		date := fm.Date
		if date == "" {
			info, err := e.Info()
			if err == nil {
				date = info.ModTime().Format(dateFormat)
			} else {
				date = time.Now().Format(dateFormat)
			}
		}

		excerpt := fm.Description
		if excerpt == "" {
			excerpt = rawMarkdownExcerpt(meta.Body, 200)
		}

		// Source path relative to cwd
		cwd, _ := os.Getwd()
		sourcePath := fullPath
		if cwd != "" {
			if rel, err := filepath.Rel(cwd, fullPath); err == nil {
				sourcePath = rel
			}
		}

		posts = append(posts, IndexPost{
			Slug:         slug,
			Title:        title,
			Date:         date,
			Author:       fm.Author,
			Tags:         fm.Tags,
			Description:  fm.Description,
			Excerpt:      excerpt,
			CategorySlug: categorySlug,
			SourcePath:   sourcePath,
			Filename:     e.Name(),
		})
	}
	return posts, nil
}

// slugFromFilename mirrors Blog.generateSlug.
func slugFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	slug := strings.ToLower(name)
	slug = slugCleanup.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

// titleFromFilename mirrors Blog.getTitleFromFilename.
func titleFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	name = datePrefix.ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "_", " ")
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

// rawMarkdownExcerpt strips Markdown syntax to produce a plain-text excerpt.
func rawMarkdownExcerpt(raw string, length int) string {
	// Strip fenced code blocks
	text := regexp.MustCompile("(?s)```[\\s\\S]*?```").ReplaceAllString(raw, "")
	// Strip inline code
	text = regexp.MustCompile("`[^`]+`").ReplaceAllString(text, "")
	// Strip headings
	text = regexp.MustCompile(`(?m)^#{1,6}\s+`).ReplaceAllString(text, "")
	// Strip images
	text = regexp.MustCompile(`!\[[^\]]*\]\([^\)]+\)`).ReplaceAllString(text, "")
	// Strip links (keep label)
	text = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`).ReplaceAllString(text, "$1")
	// Strip bold/italic
	text = regexp.MustCompile(`\*{1,2}([^*]+)\*{1,2}`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`_{1,2}([^_]+)_{1,2}`).ReplaceAllString(text, "$1")
	// Collapse whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	runes := []rune(text)
	if len(runes) <= length {
		return text
	}
	cut := string(runes[:length])
	if idx := strings.LastIndex(cut, " "); idx > 0 {
		cut = cut[:idx]
	}
	return cut + "..."
}
