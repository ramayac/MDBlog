package server

// seo.go — JSON-LD structured data helpers and static-file serving for
// sitemap.xml and robots.txt. These are separated from handler.go to keep
// SEO concerns in one place.

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
)

// ─────────────────────────────────────────────────────────────────────────────
// JSON-LD structured data
// ─────────────────────────────────────────────────────────────────────────────

// siteBaseURL returns the configured base URL (from feed config) with no trailing slash.
func siteBaseURL(cfg *config.Config) string {
	return strings.TrimRight(cfg.Feed.BaseURL, "/")
}

// marshalJSONLD encodes v as JSON and wraps it in a <script type="application/ld+json"> tag.
func marshalJSONLD(v any) template.HTML {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return template.HTML(`<script type="application/ld+json">` + string(b) + `</script>`)
}

// buildWebSiteJSONLD returns the WebSite schema for the homepage, including a SearchAction.
func buildWebSiteJSONLD(cfg *config.Config) template.HTML {
	base := siteBaseURL(cfg)
	if base == "" {
		return ""
	}
	data := map[string]any{
		"@context": "https://schema.org",
		"@type":    "WebSite",
		"name":     cfg.BlogName,
		"url":      base,
		"potentialAction": map[string]any{
			"@type": "SearchAction",
			"target": map[string]any{
				"@type":       "EntryPoint",
				"urlTemplate": base + "/?q={search_term_string}",
			},
			"query-input": "required name=search_term_string",
		},
	}
	if cfg.BlogDescription != "" {
		data["description"] = cfg.BlogDescription
	}
	return marshalJSONLD(data)
}

// buildWebPageJSONLD returns a WebPage schema for category/listing pages.
func buildWebPageJSONLD(cfg *config.Config, canonical, name, description string) template.HTML {
	base := siteBaseURL(cfg)
	if base == "" && canonical == "" {
		return ""
	}
	url := canonical
	if url == "" {
		url = base
	}
	data := map[string]any{
		"@context": "https://schema.org",
		"@type":    "WebPage",
		"name":     name,
		"url":      url,
	}
	if description != "" {
		data["description"] = description
	}
	return marshalJSONLD(data)
}

// buildArticleJSONLD returns BlogPosting + BreadcrumbList schemas for a post page.
func buildArticleJSONLD(cfg *config.Config, post *blog.Post, canonical, categorySlug string) template.HTML {
	base := siteBaseURL(cfg)
	if post == nil {
		return ""
	}
	url := canonical
	if url == "" && base != "" {
		url = base + "/post?slug=" + post.Slug
		if categorySlug != "" {
			url += "&category=" + categorySlug
		}
	}

	description := post.FrontMatter.Description
	if description == "" {
		description = post.Excerpt
	}

	author := post.FrontMatter.Author
	if author == "" {
		author = cfg.AuthorName
	}

	article := map[string]any{
		"@context":      "https://schema.org",
		"@type":         "BlogPosting",
		"headline":      post.Title,
		"url":           url,
		"datePublished": post.Date,
		"author": map[string]any{
			"@type": "Person",
			"name":  author,
		},
		"mainEntityOfPage": map[string]any{
			"@type": "WebPage",
			"@id":   url,
		},
	}
	if description != "" {
		article["description"] = description
	}

	breadcrumb := map[string]any{
		"@context":        "https://schema.org",
		"@type":           "BreadcrumbList",
		"itemListElement": buildBreadcrumbItems(cfg, post, base, url),
	}

	a, err1 := json.Marshal(article)
	bc, err2 := json.Marshal(breadcrumb)
	if err1 != nil || err2 != nil {
		return ""
	}
	return template.HTML(
		`<script type="application/ld+json">` + string(a) + `</script>` +
			"\n" +
			`<script type="application/ld+json">` + string(bc) + `</script>`,
	)
}

func buildBreadcrumbItems(cfg *config.Config, post *blog.Post, base, postURL string) []map[string]any {
	items := []map[string]any{
		{
			"@type":    "ListItem",
			"position": 1,
			"name":     cfg.BlogName,
			"item":     base + "/",
		},
	}
	pos := 2
	if post.Category != nil {
		items = append(items, map[string]any{
			"@type":    "ListItem",
			"position": pos,
			"name":     post.Category.BlogName,
			"item":     base + "/?category=" + post.CategorySlug,
		})
		pos++
	}
	items = append(items, map[string]any{
		"@type":    "ListItem",
		"position": pos,
		"name":     post.Title,
		"item":     postURL,
	})
	return items
}

// ─────────────────────────────────────────────────────────────────────────────
// sitemap.xml and robots.txt — serve pre-built files, fall back to dynamic
// ─────────────────────────────────────────────────────────────────────────────

type sitemapURL struct {
	Loc        string `xml:"loc"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
	LastMod    string `xml:"lastmod,omitempty"`
}

type sitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

// serveSitemap serves sitemap.xml. It prefers the pre-built file written by
// `make build-sitemap`; if the file is absent it falls back to a dynamically
// generated response so the dev server always works out of the box.
func (h *Handler) serveSitemap(w http.ResponseWriter, r *http.Request) {
	if !h.cfg.Sitemap.Enabled {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Header().Set("Cache-Control", h.cacheControlPages())

	// Prefer pre-built file (production path)
	if data, err := os.ReadFile(h.cfg.Sitemap.OutputFile); err == nil {
		_, _ = w.Write(data)
		return
	}

	// Fallback: generate dynamically (dev convenience)
	base := siteBaseURL(h.cfg)
	if base == "" {
		http.Error(w, "sitemap unavailable: feed.base_url not configured", http.StatusNotFound)
		return
	}

	urls := []sitemapURL{
		{Loc: base + "/", ChangeFreq: h.cfg.Sitemap.ChangeFreqHome, Priority: h.cfg.Sitemap.PriorityHome},
	}
	for _, cat := range h.b.GetCategoriesSorted() {
		urls = append(urls, sitemapURL{
			Loc:        base + "/?category=" + cat.Slug,
			ChangeFreq: h.cfg.Sitemap.ChangeFreqCategory,
			Priority:   h.cfg.Sitemap.PriorityCategory,
		})
	}
	for _, p := range h.b.GetFeedPosts(10000) {
		u := base + "/post?slug=" + p.Slug
		if p.CategorySlug != "" {
			u += "&category=" + p.CategorySlug
		}
		entry := sitemapURL{Loc: u, ChangeFreq: h.cfg.Sitemap.ChangeFreqPost, Priority: h.cfg.Sitemap.PriorityPost}
		if p.Date != "" {
			entry.LastMod = p.Date
		}
		urls = append(urls, entry)
	}

	urlset := sitemapURLSet{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9", URLs: urls}
	out, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		http.Error(w, "sitemap generation error", http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>%s`, "\n")
	_, _ = w.Write(out)
}

// serveRobots serves robots.txt. It prefers the pre-built file; falls back to
// a minimal dynamic response pointing at the sitemap.
func (h *Handler) serveRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", h.cacheControlPages())

	// Prefer pre-built file (production path)
	if data, err := os.ReadFile(h.cfg.Sitemap.RobotsFile); err == nil {
		_, _ = w.Write(data)
		return
	}

	// Fallback: minimal dynamic response
	body := "User-agent: *\nAllow: /\n"
	if base := siteBaseURL(h.cfg); base != "" && h.cfg.Sitemap.Enabled {
		body += "Sitemap: " + base + "/" + h.cfg.Sitemap.OutputFile + "\n"
	}
	_, _ = fmt.Fprint(w, body)
}
