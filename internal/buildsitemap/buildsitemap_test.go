package buildsitemap

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
)

func makeTestConfig(t *testing.T) *config.Config {
	t.Helper()
	dir := t.TempDir()
	return &config.Config{
		BlogName:      "Test Blog",
		Lang:          "en",
		PostsDir:      dir,
		PostIndexFile: filepath.Join(dir, "posts.index.json"),
		DateFormat:    "2006-01-02",
		Categories: map[string]config.Category{
			"tech": {BlogName: "Tech Posts", Folder: "tech", Index: true},
		},
		Feed: config.FeedConfig{
			BaseURL: "https://example.com",
		},
		Sitemap: config.SitemapConfig{
			Enabled:            true,
			OutputFile:         filepath.Join(dir, "sitemap.xml"),
			RobotsFile:         filepath.Join(dir, "robots.txt"),
			ChangeFreqHome:     "weekly",
			ChangeFreqCategory: "weekly",
			ChangeFreqPost:     "monthly",
			PriorityHome:       "1.0",
			PriorityCategory:   "0.8",
			PriorityPost:       "0.6",
		},
	}
}

func writePost(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func testBuildSitemap(t *testing.T, cfg *config.Config) (string, string) {
	t.Helper()
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex.Build: %v", err)
	}
	if err := Build(cfg); err != nil {
		t.Fatalf("buildsitemap.Build: %v", err)
	}
	sitemapData, err := os.ReadFile(cfg.Sitemap.OutputFile)
	if err != nil {
		t.Fatalf("read sitemap.xml: %v", err)
	}
	robotsData, err := os.ReadFile(cfg.Sitemap.RobotsFile)
	if err != nil {
		t.Fatalf("read robots.txt: %v", err)
	}
	return string(sitemapData), string(robotsData)
}

func TestBuildSitemap_XMLDeclaration(t *testing.T) {
	cfg := makeTestConfig(t)
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.HasPrefix(xmlStr, "<?xml") {
		t.Errorf("sitemap.xml should start with XML declaration, got: %q", xmlStr[:min(50, len(xmlStr))])
	}
}

func TestBuildSitemap_URLSetNamespace(t *testing.T) {
	cfg := makeTestConfig(t)
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.Contains(xmlStr, `xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"`) {
		t.Errorf("sitemap.xml should contain the sitemap namespace")
	}
}

func TestBuildSitemap_HomepageURL(t *testing.T) {
	cfg := makeTestConfig(t)
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.Contains(xmlStr, "<loc>https://example.com/</loc>") {
		t.Errorf("sitemap.xml should contain homepage URL <loc>https://example.com/</loc>")
	}
}

func TestBuildSitemap_CategoryURLs(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-post-one.md",
		"---\ntitle: Post One\ndate: 2024-01-15\n---\nContent.")
	writePost(t, techDir, "2024-02-01-post-two.md",
		"---\ntitle: Post Two\ndate: 2024-02-01\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	count := strings.Count(xmlStr, "?category=tech")
	if count != 1 {
		t.Errorf("expected category URL exactly once (deduped), got %d occurrences", count)
	}
}

func TestBuildSitemap_PostURLs(t *testing.T) {
	cfg := makeTestConfig(t)
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-my-post.md",
		"---\ntitle: My Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.Contains(xmlStr, "slug=2024-01-15-my-post") {
		t.Errorf("expected post slug in sitemap URL")
	}
	if !strings.Contains(xmlStr, "category=tech") {
		t.Errorf("expected category param in post URL")
	}
	if !strings.Contains(xmlStr, "<lastmod>2024-01-15</lastmod>") {
		t.Errorf("expected <lastmod> with post date")
	}
}

func TestBuildSitemap_ConfigurablePriority(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Sitemap.PriorityHome = "0.9"
	cfg.Sitemap.PriorityCategory = "0.7"
	cfg.Sitemap.PriorityPost = "0.5"
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	for _, want := range []string{"<priority>0.9</priority>", "<priority>0.7</priority>", "<priority>0.5</priority>"} {
		if !strings.Contains(xmlStr, want) {
			t.Errorf("expected %q in sitemap.xml", want)
		}
	}
}

func TestBuildSitemap_ConfigurableChangeFreq(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Sitemap.ChangeFreqHome = "daily"
	cfg.Sitemap.ChangeFreqCategory = "daily"
	cfg.Sitemap.ChangeFreqPost = "yearly"
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.Contains(xmlStr, "<changefreq>daily</changefreq>") {
		t.Errorf("expected <changefreq>daily</changefreq> in sitemap.xml")
	}
	if !strings.Contains(xmlStr, "<changefreq>yearly</changefreq>") {
		t.Errorf("expected <changefreq>yearly</changefreq> in sitemap.xml")
	}
}

func TestBuildSitemap_RobotsFile(t *testing.T) {
	cfg := makeTestConfig(t)
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	_, robots := testBuildSitemap(t, cfg)
	for _, want := range []string{"User-agent: *", "Allow: /", "Sitemap: https://example.com/"} {
		if !strings.Contains(robots, want) {
			t.Errorf("robots.txt should contain %q, got:\n%s", want, robots)
		}
	}
}

func TestBuildSitemap_Disabled(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Sitemap.Enabled = false
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex: %v", err)
	}
	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}
	if _, err := os.Stat(cfg.Sitemap.OutputFile); !os.IsNotExist(err) {
		t.Errorf("expected sitemap.xml to not exist when disabled, but it was created")
	}
	if _, err := os.Stat(cfg.Sitemap.RobotsFile); !os.IsNotExist(err) {
		t.Errorf("expected robots.txt to not exist when disabled, but it was created")
	}
}

func TestBuildSitemap_MissingIndex(t *testing.T) {
	cfg := makeTestConfig(t)
	err := Build(cfg)
	if err == nil {
		t.Error("expected error when post index is missing, got nil")
	}
}

func TestBuildSitemap_MissingBaseURL(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Feed.BaseURL = ""
	writePost(t, filepath.Join(cfg.PostsDir, "tech"), "2024-01-15-post.md",
		"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex: %v", err)
	}
	err := Build(cfg)
	if err == nil {
		t.Error("expected error when feed.base_url is empty, got nil")
	}
}

func TestBuildSitemap_PostWithoutCategory(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.ShowUncategorized = true
	writePost(t, cfg.PostsDir, "2024-01-15-root-post.md",
		"---\ntitle: Root Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr, _ := testBuildSitemap(t, cfg)
	if !strings.Contains(xmlStr, "slug=2024-01-15-root-post") {
		t.Errorf("expected root post slug in sitemap")
	}
	if strings.Contains(xmlStr, "slug=2024-01-15-root-post&amp;category=") {
		t.Errorf("uncategorized post URL should not include category param")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
