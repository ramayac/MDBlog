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

func TestDefaults(t *testing.T) {
	toml := `blog_name = "X"`
	tmp := filepath.Join(t.TempDir(), "config.toml")
	_ = os.WriteFile(tmp, []byte(toml), 0644)
	cfg, _ := Load(tmp)
	if cfg.Lang != "en" {
		t.Errorf("default lang = %q, want 'en'", cfg.Lang)
	}
	if cfg.PostsPerPage != 10 {
		t.Errorf("default posts_per_page = %d, want 10", cfg.PostsPerPage)
	}
}
