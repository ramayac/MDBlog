package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	// Write a minimal TOML to a temp file.
	toml := `
blog_name = "Test Blog"
author_name = "Tester"
posts_per_page = 5
posts_dir = "posts"
post_index_file = "posts/posts.index.json"
date_format = "2006-01-02"

[csp]
enabled = true
header = "Content-Security-Policy: default-src 'self';"

[[menu_links]]
label = "Home"
url   = "/"

[categories.foo]
blog_name = "Foo"
folder    = "foo"
index     = true

[labels]
read_more = "Read more"
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	if err := os.WriteFile(tmp, []byte(toml), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(tmp)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.BlogName != "Test Blog" {
		t.Errorf("BlogName = %q, want %q", cfg.BlogName, "Test Blog")
	}
	if cfg.PostsPerPage != 5 {
		t.Errorf("PostsPerPage = %d, want 5", cfg.PostsPerPage)
	}
	if len(cfg.MenuLinks) != 1 || cfg.MenuLinks[0].Label != "Home" {
		t.Errorf("MenuLinks[0] = %+v, want label 'Home'", cfg.MenuLinks)
	}
	if cat, ok := cfg.Categories["foo"]; !ok || cat.BlogName != "Foo" {
		t.Errorf("categories.foo missing or wrong: %+v", cfg.Categories)
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("/nonexistent/path/config.toml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoad_InvalidTOML(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte("not = [valid toml"), 0644)
	_, err := Load(tmp)
	if err == nil {
		t.Error("expected error for invalid TOML, got nil")
	}
}

func TestMustLoad_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustLoad should panic on missing file")
		}
	}()
	MustLoad("/nonexistent/config.toml")
}

func TestMustLoad_OK(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(`blog_name = "OK"`), 0644)
	cfg := MustLoad(tmp)
	if cfg.BlogName != "OK" {
		t.Errorf("BlogName = %q, want 'OK'", cfg.BlogName)
	}
}

func TestLoad_CSP(t *testing.T) {
	toml := `
[csp]
enabled = true
header = "Content-Security-Policy: default-src 'self';"
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.CSP.Enabled {
		t.Error("CSP.Enabled should be true")
	}
	if cfg.CSP.Header == "" {
		t.Error("CSP.Header should not be empty")
	}
}

func TestLoad_Labels(t *testing.T) {
	toml := `
[labels]
read_more = "Continue →"
pagination_prev = "← Newer"
pagination_next = "Older →"
search_title = "Find"
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Labels.ReadMore != "Continue →" {
		t.Errorf("ReadMore = %q, want 'Continue →'", cfg.Labels.ReadMore)
	}
	if cfg.Labels.SearchTitle != "Find" {
		t.Errorf("SearchTitle = %q, want 'Find'", cfg.Labels.SearchTitle)
	}
}

func TestLoad_MultipleMenuLinks(t *testing.T) {
	toml := `
[[menu_links]]
label = "Home"
url   = "/"

[[menu_links]]
label = "About"
url   = "/about"
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.MenuLinks) != 2 {
		t.Fatalf("expected 2 menu links, got %d", len(cfg.MenuLinks))
	}
	if cfg.MenuLinks[1].Label != "About" {
		t.Errorf("MenuLinks[1].Label = %q, want 'About'", cfg.MenuLinks[1].Label)
	}
}

func TestDefaults_ExcerptLength(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(`blog_name = "X"`), 0644)
	cfg, _ := Load(tmp)
	if cfg.ExcerptLength != 200 {
		t.Errorf("default excerpt_length = %d, want 200", cfg.ExcerptLength)
	}
}

func TestCacheDefaults_WhenNotSet(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(`blog_name = "X"`), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Cache.Enabled {
		t.Error("cache should be enabled by default when not configured")
	}
	if cfg.Cache.MaxAgePages != 3600 {
		t.Errorf("MaxAgePages = %d, want 3600", cfg.Cache.MaxAgePages)
	}
	if cfg.Cache.MaxAgeAssets != 86400 {
		t.Errorf("MaxAgeAssets = %d, want 86400", cfg.Cache.MaxAgeAssets)
	}
}

func TestCacheExplicitValues_ArePreserved(t *testing.T) {
	toml := `
[cache]
enabled        = true
max_age_pages  = 600
max_age_assets = 7200
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Cache.Enabled {
		t.Error("Cache.Enabled should be true")
	}
	if cfg.Cache.MaxAgePages != 600 {
		t.Errorf("MaxAgePages = %d, want 600", cfg.Cache.MaxAgePages)
	}
	if cfg.Cache.MaxAgeAssets != 7200 {
		t.Errorf("MaxAgeAssets = %d, want 7200", cfg.Cache.MaxAgeAssets)
	}
}

func TestCacheDisabled_WhenExplicitlySet(t *testing.T) {
	// Non-zero TTLs prevent the "not configured" heuristic from firing.
	toml := `
[cache]
enabled        = false
max_age_pages  = 3600
max_age_assets = 86400
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Cache.Enabled {
		t.Error("Cache.Enabled should be false when explicitly disabled")
	}
}

func TestLoad_SitemapDefaults(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(`blog_name = "X"`), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Sitemap.OutputFile != "sitemap.xml" {
		t.Errorf("default output_file = %q, want 'sitemap.xml'", cfg.Sitemap.OutputFile)
	}
	if cfg.Sitemap.RobotsFile != "robots.txt" {
		t.Errorf("default robots_file = %q, want 'robots.txt'", cfg.Sitemap.RobotsFile)
	}
	if cfg.Sitemap.ChangeFreqHome != "weekly" {
		t.Errorf("default changefreq_home = %q, want 'weekly'", cfg.Sitemap.ChangeFreqHome)
	}
	if cfg.Sitemap.ChangeFreqCategory != "weekly" {
		t.Errorf("default changefreq_category = %q, want 'weekly'", cfg.Sitemap.ChangeFreqCategory)
	}
	if cfg.Sitemap.ChangeFreqPost != "monthly" {
		t.Errorf("default changefreq_post = %q, want 'monthly'", cfg.Sitemap.ChangeFreqPost)
	}
	if cfg.Sitemap.PriorityHome != "1.0" {
		t.Errorf("default priority_home = %q, want '1.0'", cfg.Sitemap.PriorityHome)
	}
	if cfg.Sitemap.PriorityCategory != "0.8" {
		t.Errorf("default priority_category = %q, want '0.8'", cfg.Sitemap.PriorityCategory)
	}
	if cfg.Sitemap.PriorityPost != "0.6" {
		t.Errorf("default priority_post = %q, want '0.6'", cfg.Sitemap.PriorityPost)
	}
}

func TestLoad_SitemapCustomValues(t *testing.T) {
	toml := `
[sitemap]
enabled              = true
output_file          = "mysitemap.xml"
robots_file          = "myrobots.txt"
changefreq_home      = "daily"
changefreq_category  = "daily"
changefreq_post      = "yearly"
priority_home        = "0.9"
priority_category    = "0.7"
priority_post        = "0.5"

[feed]
base_url = "https://example.com"
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, err := Load(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Sitemap.Enabled {
		t.Error("Sitemap.Enabled should be true")
	}
	if cfg.Sitemap.OutputFile != "mysitemap.xml" {
		t.Errorf("OutputFile = %q, want 'mysitemap.xml'", cfg.Sitemap.OutputFile)
	}
	if cfg.Sitemap.ChangeFreqPost != "yearly" {
		t.Errorf("ChangeFreqPost = %q, want 'yearly'", cfg.Sitemap.ChangeFreqPost)
	}
	if cfg.Sitemap.PriorityPost != "0.5" {
		t.Errorf("PriorityPost = %q, want '0.5'", cfg.Sitemap.PriorityPost)
	}
}

func TestLoad_SitemapEnabled_RequiresBaseURL(t *testing.T) {
	toml := `
[sitemap]
enabled = true
`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	_, err := Load(tmp)
	if err == nil {
		t.Error("expected error when sitemap.enabled = true without feed.base_url, got nil")
	}
}
