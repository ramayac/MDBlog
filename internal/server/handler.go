package server

import (
	"compress/gzip"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
)

// TemplateFS is the filesystem used to load templates.
// In production this is the real FS; tests may override it.
var TemplateFS fs.FS = os.DirFS("templates")

// AssetsFS is the filesystem used to serve /assets/* files.
// Override in embed builds to use an embedded FS.
var AssetsFS fs.FS = os.DirFS("assets")

// templateData is the top-level data passed to every page template.
type templateData struct {
	Config          *config.Config
	PageTitle       string
	PageDescription string
	OGType          string
	Canonical       string
	CSSVersion      string
	Menu            []blog.MenuLink
	FooterHTML      template.HTML
	Version         blog.VersionInfo
	RenderTime      string
	JSFiles         []string
	Content         template.HTML // rendered inner content
	// page-specific fields
	Posts           []blog.Post
	Post            *blog.Post
	Tags            []string
	CategorySlug    string
	CurrentCategory *blog.CategoryInfo
	Pagination      blog.Pagination
	Categories      []blog.CategoryInfo
	IndexBlurb      template.HTML
	Query           string
}

// postPreviewData is passed to the post_preview template.
type postPreviewData struct {
	Post    blog.Post
	PostURL string
	Config  *config.Config
}

// Handler is the main HTTP handler for the blog.
type Handler struct {
	cfg     *config.Config
	b       *blog.Blog
	tmpl    *template.Template
	mimeMap map[string]string
}

// New creates and returns the blog HTTP handler.
// It panics if templates cannot be loaded.
func New(cfg *config.Config, b *blog.Blog) *Handler {
	tmpl := mustLoadTemplates()
	h := &Handler{
		cfg:  cfg,
		b:    b,
		tmpl: tmpl,
		mimeMap: map[string]string{
			"css":   "text/css",
			"js":    "application/javascript",
			"png":   "image/png",
			"jpg":   "image/jpeg",
			"jpeg":  "image/jpeg",
			"gif":   "image/gif",
			"svg":   "image/svg+xml",
			"ico":   "image/x-icon",
			"woff":  "font/woff",
			"woff2": "font/woff2",
			"ttf":   "font/ttf",
		},
	}
	return h
}

// ServeHTTP routes all requests.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// ── Static assets ──────────────────────────────────────────────────────
	if strings.HasPrefix(path, "/assets/") {
		h.serveAsset(w, r, path)
		return
	}

	// ── CSP header —————————————————————————————————————————————————————————
	if h.cfg.CSP.Enabled {
		w.Header().Set(splitCSPHeader(h.cfg.CSP.Header))
	}

	// ── Route: RSS feed (/feed.xml, /feed) ────────────────────────────────
	if path == "/feed.xml" {
		h.serveFeedXML(w, r)
		return
	}
	if path == "/feed" {
		h.serveFeedPage(w, r)
		return
	}

	// ── Route: single post (/post) ─────────────────────────────────────────
	if path == "/post" || strings.HasSuffix(path, "/post") {
		h.servPost(w, r)
		return
	}

	// ── Route: index (/) ──────────────────────────────────────────────────
	h.serveIndex(w, r)
}

// ─────────────────────────────────────────────────────────────────────────────
// Route handlers
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) serveIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	q := r.URL.Query()
	categorySlug := q.Get("category")
	searchQuery := q.Get("q")
	isSearch := q.Has("search") || searchQuery != ""
	page := max1(intParam(q.Get("page"), 1))

	menu := h.b.GetMenu()
	canonical := buildCanonical(r)
	cssV := cssVersion(h.cfg.CSSTheme)
	versionInfo := h.b.GetVersionInfo()
	footerHTML := template.HTML(h.b.ParseMarkdown(h.cfg.FooterContent))

	base := templateData{
		Config:     h.cfg,
		OGType:     "website",
		Canonical:  canonical,
		CSSVersion: cssV,
		Menu:       menu,
		FooterHTML: footerHTML,
		Version:    versionInfo,
	}

	if isSearch {
		title := h.cfg.BlogName
		if searchQuery != "" {
			title = fmt.Sprintf(h.cfg.Labels.SearchResultsTitle, searchQuery) + " - " + h.cfg.BlogName
		} else {
			title = h.cfg.Labels.SearchTitle + " - " + h.cfg.BlogName
		}
		list := h.b.SearchPosts(searchQuery, page)
		data := base
		data.PageTitle = title
		data.Posts = list.Posts
		data.Pagination = list.Pagination
		data.Query = searchQuery
		h.renderPage(w, r, start, "search.html", &data)
		return
	}

	if categorySlug != "" {
		cat := h.b.GetCategoryBySlug(categorySlug)
		if cat == nil {
			h.serve404(w, r, start, menu, canonical, cssV)
			return
		}
		list := h.b.GetPosts(page, categorySlug)
		data := base
		data.PageTitle = h.cfg.BlogName + " - " + cat.BlogName
		data.CurrentCategory = cat
		data.CategorySlug = categorySlug
		data.Posts = list.Posts
		data.Pagination = list.Pagination
		h.renderPage(w, r, start, "category.html", &data)
		return
	}

	// Landing page — no post scanning
	cats := h.b.GetCategoriesSorted()
	blurb := ""
	blurbFile := filepath.Join(h.cfg.PostsDir, "index.md")
	if fileExists(blurbFile) {
		raw, err := os.ReadFile(blurbFile)
		if err == nil {
			blurb = h.b.ParseMarkdown(string(raw))
		}
	}
	data := base
	data.PageTitle = h.cfg.BlogName
	data.Categories = cats
	data.IndexBlurb = template.HTML(blurb)
	h.renderPage(w, r, start, "home.html", &data)
}

func (h *Handler) servPost(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	q := r.URL.Query()
	slug := q.Get("slug")
	categorySlug := q.Get("category")

	menu := h.b.GetMenu()
	canonical := buildCanonical(r)
	cssV := cssVersion(h.cfg.CSSTheme)
	footerHTML := template.HTML(h.b.ParseMarkdown(h.cfg.FooterContent))
	versionInfo := h.b.GetVersionInfo()

	base := templateData{
		Config:     h.cfg,
		OGType:     "article",
		Canonical:  canonical,
		CSSVersion: cssV,
		Menu:       menu,
		FooterHTML: footerHTML,
		Version:    versionInfo,
	}

	if slug == "" || !validSlug(slug) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	post := h.b.GetPostBySlug(slug, categorySlug)
	if post == nil {
		w.WriteHeader(http.StatusNotFound)
		data := base
		data.PageTitle = h.cfg.Labels.NotFoundTitle + " — " + h.cfg.BlogName
		h.renderPage(w, r, start, "404.html", &data)
		return
	}

	// ETag / 304
	postMtime := int64(0)
	if t, err := time.Parse("2006-01-02", post.Date); err == nil {
		postMtime = t.Unix()
	}
	etag := fmt.Sprintf(`"%x"`, md5sum(post.Slug+post.Date))
	lastMod := time.Unix(postMtime, 0).UTC().Format(http.TimeFormat)
	w.Header().Set("Last-Modified", lastMod)
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", h.cacheControlPages())

	ifNoneMatch := strings.TrimSpace(r.Header.Get("If-None-Match"))
	ifModSince := r.Header.Get("If-Modified-Since")
	if ifNoneMatch == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	if ifModSince != "" {
		if t, err := http.ParseTime(ifModSince); err == nil && t.Unix() >= postMtime {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	// Parse JS files from front matter
	var jsFiles []string
	if post.FrontMatter.JS != "" {
		jsFile := strings.TrimSpace(post.FrontMatter.JS)
		if jsFile != "" && fileExists("assets/js/"+jsFile) {
			jsFiles = []string{jsFile}
		}
	}

	// Parse tags
	var tags []string
	if post.FrontMatter.Tags != "" {
		for _, t := range strings.Split(post.FrontMatter.Tags, ",") {
			if trimmed := strings.TrimSpace(t); trimmed != "" {
				tags = append(tags, trimmed)
			}
		}
	}

	data := base
	data.PageTitle = post.Title + " - " + h.cfg.BlogName
	data.PageDescription = post.FrontMatter.Description
	data.Post = post
	data.CategorySlug = categorySlug
	data.JSFiles = jsFiles
	data.Tags = tags
	h.renderPage(w, r, start, "post.html", &data)
}

func (h *Handler) cacheControlPages() string {
	if !h.cfg.Cache.Enabled {
		return "no-store"
	}
	return fmt.Sprintf("public, max-age=%d", h.cfg.Cache.MaxAgePages)
}

func (h *Handler) serveFeedXML(w http.ResponseWriter, r *http.Request) {
	if !h.cfg.Feed.Enabled {
		http.NotFound(w, r)
		return
	}
	data, err := os.ReadFile(h.cfg.Feed.OutputFile)
	if err != nil {
		http.Error(w, "feed not available", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.Header().Set("Cache-Control", h.cacheControlPages())

	var out io.Writer = w
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		out = gz
	}
	_, _ = out.Write(data)
}

func (h *Handler) serveFeedPage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	menu := h.b.GetMenu()
	canonical := buildCanonical(r)
	cssV := cssVersion(h.cfg.CSSTheme)
	footerHTML := template.HTML(h.b.ParseMarkdown(h.cfg.FooterContent))
	versionInfo := h.b.GetVersionInfo()

	posts := h.b.GetFeedPosts(h.cfg.Feed.MaxItems)

	data := &templateData{
		Config:     h.cfg,
		PageTitle:  "Feed — " + h.cfg.BlogName,
		OGType:     "website",
		Canonical:  canonical,
		CSSVersion: cssV,
		Menu:       menu,
		FooterHTML: footerHTML,
		Version:    versionInfo,
		Posts:      posts,
	}
	h.renderPage(w, r, start, "feed.html", data)
}

func (h *Handler) cacheControlAssets() string {
	if !h.cfg.Cache.Enabled {
		return "no-store"
	}
	return fmt.Sprintf("public, max-age=%d", h.cfg.Cache.MaxAgeAssets)
}

func (h *Handler) serve404(w http.ResponseWriter, r *http.Request, start time.Time, menu []blog.MenuLink, canonical, cssV string) {
	w.WriteHeader(http.StatusNotFound)
	h.renderPage(w, r, start, "404.html", &templateData{
		Config:     h.cfg,
		PageTitle:  h.cfg.Labels.NotFoundTitle + " — " + h.cfg.BlogName,
		Canonical:  canonical,
		CSSVersion: cssV,
		Menu:       menu,
	})
}

func (h *Handler) serveAsset(w http.ResponseWriter, r *http.Request, urlPath string) {
	// Strip leading /assets/ prefix to get the fs-relative path.
	rel := strings.TrimPrefix(urlPath, "/assets/")

	// Validate: reject traversal sequences and null bytes.
	if strings.Contains(rel, "..") || strings.Contains(rel, "\x00") {
		http.NotFound(w, r)
		return
	}

	f, err := AssetsFS.Open(rel)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil || st.IsDir() {
		http.NotFound(w, r)
		return
	}

	rs, ok := f.(io.ReadSeeker)
	if !ok {
		http.Error(w, "asset not seekable", http.StatusInternalServerError)
		return
	}

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(rel), "."))
	contentType, ok := h.mimeMap[ext]
	if !ok {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", h.cacheControlAssets())
	http.ServeContent(w, r, rel, st.ModTime(), rs)
}

// ─────────────────────────────────────────────────────────────────────────────
// Template rendering
// ─────────────────────────────────────────────────────────────────────────────

func (h *Handler) renderPage(w http.ResponseWriter, r *http.Request, start time.Time, tmplName string, data *templateData) {
	// Render the inner page template into a buffer, then wrap with layout.
	var innerBuf strings.Builder
	if err := h.tmpl.ExecuteTemplate(&innerBuf, tmplName, data); err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data.Content = template.HTML(innerBuf.String())

	if h.cfg.ShowRenderTime {
		ms := math.Round(float64(time.Since(start).Microseconds()) / 1000.0)
		data.RenderTime = fmt.Sprintf("%.2f", ms)
		log.Printf("render %s %s %.2fms", r.Method, r.URL.RequestURI(), ms)
	}

	// Gzip when supported (skip on Lambda — API Gateway / CloudFront handles it)
	var out io.Writer = w
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		out = gz
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := h.tmpl.ExecuteTemplate(out, "layout.html", data); err != nil {
		// Headers already sent — can't send a proper error response
		fmt.Fprintf(out, "\n<!-- render error: %v -->", err)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Template loading
// ─────────────────────────────────────────────────────────────────────────────

func mustLoadTemplates() *template.Template {
	funcMap := template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		"formatDate": func(dateStr string) string {
			t, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return dateStr
			}
			return t.Format("January 2, 2006")
		},
		"urlquery": func(s string) string { return s }, // stdlib already does this
		"printf":   fmt.Sprintf,
		"not": func(v interface{}) bool {
			switch x := v.(type) {
			case bool:
				return !x
			case []blog.Post:
				return len(x) == 0
			case string:
				return x == ""
			}
			return false
		},
		"defaultStr": func(dflt, s string) string {
			if s == "" {
				return dflt
			}
			return s
		},
		"postPreviewData": func(p blog.Post, catSlug string, cfg *config.Config) postPreviewData {
			slug := p.CategorySlug
			if slug == "" {
				slug = catSlug
			}
			url := "/post?slug=" + p.Slug
			if slug != "" {
				url += "&category=" + slug
			}
			return postPreviewData{Post: p, PostURL: url, Config: cfg}
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	entries, err := fs.ReadDir(TemplateFS, ".")
	if err != nil {
		panic("server: cannot read templates dir: " + err.Error())
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".html") {
			continue
		}
		data, err := fs.ReadFile(TemplateFS, e.Name())
		if err != nil {
			panic("server: cannot read template " + e.Name() + ": " + err.Error())
		}
		name := e.Name()
		t := tmpl.New(name)
		if _, err := t.Parse(string(data)); err != nil {
			panic("server: cannot parse template " + name + ": " + err.Error())
		}
	}
	return tmpl
}

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

// splitCSPHeader splits "Header-Name: value" into name, value for w.Header().Set.
func splitCSPHeader(h string) (string, string) {
	if idx := strings.Index(h, ":"); idx > 0 {
		return strings.TrimSpace(h[:idx]), strings.TrimSpace(h[idx+1:])
	}
	return h, ""
}

func buildCanonical(r *http.Request) string {
	scheme := "https"
	if r.TLS == nil && r.Header.Get("X-Forwarded-Proto") == "" {
		scheme = "http"
	}
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}
	host := r.Host
	if host == "" {
		return ""
	}
	uri := r.URL.RequestURI()
	// Remove ?page= to avoid duplicate canonicals on paginated pages
	uri = removePaginationParam(uri)
	return scheme + "://" + host + uri
}

func removePaginationParam(uri string) string {
	// Very lightweight: strip &page=N or ?page=N
	idx := strings.Index(uri, "?")
	if idx < 0 {
		return uri
	}
	base := uri[:idx]
	query := uri[idx+1:]
	var parts []string
	for _, kv := range strings.Split(query, "&") {
		if !strings.HasPrefix(kv, "page=") {
			parts = append(parts, kv)
		}
	}
	if len(parts) == 0 {
		return base
	}
	return base + "?" + strings.Join(parts, "&")
}

func cssVersion(cssPath string) string {
	info, err := os.Stat(cssPath)
	if err != nil {
		return "0"
	}
	return strconv.FormatInt(info.ModTime().Unix(), 10)
}

func validSlug(s string) bool {
	if len(s) > 200 {
		return false
	}
	for _, c := range s {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func intParam(s string, def int) int {
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func max1(n int) int {
	if n < 1 {
		return 1
	}
	return n
}

// md5sum returns a hex string of the MD5 of s (used for ETag).
func md5sum(s string) string {
	// Use crypto/md5 via fmt to avoid importing the package at the top level
	// (it's only called from servPost).
	_ = runtime.Version() // suppress unused import if crypto/md5 moved
	h := simpleHash(s)
	return fmt.Sprintf("%x", h)
}

// simpleHash is a fast non-cryptographic hash used only for ETags.
func simpleHash(s string) uint64 {
	var h uint64 = 14695981039346656037
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}
