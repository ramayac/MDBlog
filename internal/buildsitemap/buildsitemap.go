package buildsitemap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ramayac/mdblog/internal/config"
)

// indexPost mirrors the JSON fields in posts.index.json.
type indexPost struct {
	Slug         string `json:"slug"`
	Title        string `json:"title"`
	Date         string `json:"date"`
	CategorySlug string `json:"category_slug"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Sitemap XML structures
// ─────────────────────────────────────────────────────────────────────────────

type sitemapEntry struct {
	Loc        string `xml:"loc"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
	LastMod    string `xml:"lastmod,omitempty"`
}

type sitemapURLSet struct {
	XMLName xml.Name       `xml:"urlset"`
	XMLNS   string         `xml:"xmlns,attr"`
	URLs    []sitemapEntry `xml:"url"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Build
// ─────────────────────────────────────────────────────────────────────────────

// Build reads posts.index.json, writes sitemap.xml and robots.txt atomically.
func Build(cfg *config.Config) error {
	if !cfg.Sitemap.Enabled {
		fmt.Println("buildsitemap: sitemap disabled, skipping")
		return nil
	}

	base := strings.TrimRight(cfg.Feed.BaseURL, "/")
	if base == "" {
		return fmt.Errorf("buildsitemap: feed.base_url must be set when sitemap.enabled is true")
	}

	// Load post index
	indexData, err := os.ReadFile(cfg.PostIndexFile)
	if err != nil {
		return fmt.Errorf("buildsitemap: cannot read index %s: %w", cfg.PostIndexFile, err)
	}
	var posts []indexPost
	if err := json.Unmarshal(indexData, &posts); err != nil {
		return fmt.Errorf("buildsitemap: cannot parse index: %w", err)
	}

	// Build URL list
	urls := []sitemapEntry{
		{
			Loc:        base + "/",
			ChangeFreq: cfg.Sitemap.ChangeFreqHome,
			Priority:   cfg.Sitemap.PriorityHome,
		},
	}

	// Category pages — collect unique slugs in index order
	seen := map[string]bool{}
	for _, p := range posts {
		if p.CategorySlug != "" && !seen[p.CategorySlug] {
			seen[p.CategorySlug] = true
			urls = append(urls, sitemapEntry{
				Loc:        base + "/?category=" + url.QueryEscape(p.CategorySlug),
				ChangeFreq: cfg.Sitemap.ChangeFreqCategory,
				Priority:   cfg.Sitemap.PriorityCategory,
			})
		}
	}

	// Individual post pages
	for _, p := range posts {
		u := base + "/post?slug=" + url.QueryEscape(p.Slug)
		if p.CategorySlug != "" {
			u += "&category=" + url.QueryEscape(p.CategorySlug)
		}
		entry := sitemapEntry{
			Loc:        u,
			ChangeFreq: cfg.Sitemap.ChangeFreqPost,
			Priority:   cfg.Sitemap.PriorityPost,
		}
		if p.Date != "" {
			entry.LastMod = p.Date
		}
		urls = append(urls, entry)
	}

	urlset := sitemapURLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	xmlBytes, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		return fmt.Errorf("buildsitemap: marshal xml: %w", err)
	}
	output := []byte(xml.Header + string(xmlBytes) + "\n")

	if err := atomicWrite(cfg.Sitemap.OutputFile, output); err != nil {
		return fmt.Errorf("buildsitemap: write sitemap: %w", err)
	}
	fmt.Printf("buildsitemap: wrote %d URLs to %s\n", len(urls), cfg.Sitemap.OutputFile)

	// robots.txt
	robots := "User-agent: *\nAllow: /\nSitemap: " + base + "/" + cfg.Sitemap.OutputFile + "\n"
	if err := atomicWrite(cfg.Sitemap.RobotsFile, []byte(robots)); err != nil {
		return fmt.Errorf("buildsitemap: write robots: %w", err)
	}
	fmt.Printf("buildsitemap: wrote %s\n", cfg.Sitemap.RobotsFile)

	return nil
}

func atomicWrite(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}
