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
url   = "index.php"

[categories.foo]
blog_name = "Foo"
folder    = "foo"
index     = true
menu      = true

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
