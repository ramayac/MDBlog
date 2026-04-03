package buildfeed

import (
"encoding/json"
"encoding/xml"
"fmt"
"net/url"
"os"
"path/filepath"
"time"

"github.com/ramayac/mdblog/internal/config"
)

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

// ─────────────────────────────────────────────────────────────────────────────
// RSS 2.0 XML structures
// ─────────────────────────────────────────────────────────────────────────────

type rssRoot struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Atom    string     `xml:"xmlns:atom,attr"`
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	LastBuildDate string    `xml:"lastBuildDate"`
	AtomLink      atomLink  `xml:"atom:link"`
	Items         []rssItem `xml:"item"`
}

type atomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type rssItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	GUID        rssGUID `xml:"guid"`
	PubDate     string  `xml:"pubDate"`
	Description string  `xml:"description"`
	Category    string  `xml:"category,omitempty"`
}

type rssGUID struct {
	IsPermaLink string `xml:"isPermaLink,attr"`
	Value       string `xml:",chardata"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Build
// ─────────────────────────────────────────────────────────────────────────────

// Build reads posts.index.json and writes feed.xml atomically.
func Build(cfg *config.Config) error {
	if !cfg.Feed.Enabled {
		fmt.Println("buildfeed: feed disabled, skipping")
		return nil
	}

	// Load post index
	indexData, err := os.ReadFile(cfg.PostIndexFile)
	if err != nil {
		return fmt.Errorf("buildfeed: cannot read index %s: %w", cfg.PostIndexFile, err)
	}
	var posts []indexPost
	if err := json.Unmarshal(indexData, &posts); err != nil {
		return fmt.Errorf("buildfeed: cannot parse index: %w", err)
	}

	// Cap at max_items (index is already sorted newest-first)
	if len(posts) > cfg.Feed.MaxItems {
		posts = posts[:cfg.Feed.MaxItems]
	}

	baseURL := cfg.Feed.BaseURL

	// Build lastBuildDate from the newest post
	var lastBuildDate string
	if len(posts) > 0 {
		if t, err := time.Parse("2006-01-02", posts[0].Date); err == nil {
			lastBuildDate = t.UTC().Format(time.RFC1123Z)
		}
	}
	if lastBuildDate == "" {
		lastBuildDate = time.Now().UTC().Format(time.RFC1123Z)
	}

	// Build items
	items := make([]rssItem, 0, len(posts))
	for _, p := range posts {
		link := buildPostURL(baseURL, p.Slug, p.CategorySlug)
		pubDate := ""
		if t, err := time.Parse("2006-01-02", p.Date); err == nil {
			pubDate = t.UTC().Format(time.RFC1123Z)
		}

		categoryName := ""
		if p.CategorySlug != "" {
			if cat, ok := cfg.Categories[p.CategorySlug]; ok {
				categoryName = cat.BlogName
			}
		}

		items = append(items, rssItem{
Title:       p.Title,
Link:        link,
GUID:        rssGUID{IsPermaLink: "true", Value: link},
PubDate:     pubDate,
Description: p.Excerpt,
Category:    categoryName,
})
	}

	feedURL := baseURL + "/feed.xml"

	root := rssRoot{
		Version: "2.0",
		Atom:    "http://www.w3.org/2005/Atom",
		Channel: rssChannel{
			Title:         cfg.Feed.Title,
			Link:          baseURL,
			Description:   cfg.Feed.Description,
			Language:      cfg.Lang,
			LastBuildDate: lastBuildDate,
			AtomLink: atomLink{
				Href: feedURL,
				Rel:  "self",
				Type: "application/rss+xml",
			},
			Items: items,
		},
	}

	// Marshal XML
	xmlBytes, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		return fmt.Errorf("buildfeed: marshal xml: %w", err)
	}
	output := []byte(xml.Header + string(xmlBytes) + "\n")

	// Atomic write: tmp -> rename
	outPath := cfg.Feed.OutputFile
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("buildfeed: mkdir: %w", err)
	}
	tmp := outPath + ".tmp"
	if err := os.WriteFile(tmp, output, 0644); err != nil {
		return fmt.Errorf("buildfeed: write tmp: %w", err)
	}
	if err := os.Rename(tmp, outPath); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("buildfeed: rename: %w", err)
	}

	fmt.Printf("buildfeed: wrote %d items to %s\n", len(items), outPath)
	return nil
}

// buildPostURL constructs an absolute post URL.
func buildPostURL(baseURL, slug, categorySlug string) string {
	u := baseURL + "/post?slug=" + url.QueryEscape(slug)
	if categorySlug != "" {
		u += "&category=" + url.QueryEscape(categorySlug)
	}
	return u
}
