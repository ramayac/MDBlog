package buildindex

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ramayac/mdblog/internal/config"
)

func makeTestConfig(t *testing.T) *config.Config {
	t.Helper()
	dir := t.TempDir()
	return &config.Config{
		PostsDir:      dir,
		PostIndexFile: filepath.Join(dir, "posts.index.json"),
		DateFormat:    "2006-01-02",
		Categories: map[string]config.Category{
			"tech": {BlogName: "Tech", Folder: "tech", Index: true, Menu: true},
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

func loadIndex(t *testing.T, path string) []IndexPost {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	var posts []IndexPost
	if err := json.Unmarshal(data, &posts); err != nil {
		t.Fatalf("parse index: %v", err)
	}
	return posts
}

func TestBuild_Basic(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-hello-world.md",
		"---\ntitle: Hello World\ndate: 2024-01-15\nauthor: Alice\ntags: go, test\ndescription: A test post\n---\n\nBody content here.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(posts))
	}
	p := posts[0]

	if p.Title != "Hello World" {
		t.Errorf("Title = %q, want 'Hello World'", p.Title)
	}
	if p.Date != "2024-01-15" {
		t.Errorf("Date = %q, want '2024-01-15'", p.Date)
	}
	if p.Author != "Alice" {
		t.Errorf("Author = %q, want 'Alice'", p.Author)
	}
	if p.Tags != "go, test" {
		t.Errorf("Tags = %q, want 'go, test'", p.Tags)
	}
	if p.Description != "A test post" {
		t.Errorf("Description = %q, want 'A test post'", p.Description)
	}
	if p.Slug != "2024-01-15-hello-world" {
		t.Errorf("Slug = %q, want '2024-01-15-hello-world'", p.Slug)
	}
	if p.Filename != "2024-01-15-hello-world.md" {
		t.Errorf("Filename = %q, want '2024-01-15-hello-world.md'", p.Filename)
	}
	if p.CategorySlug != "tech" {
		t.Errorf("CategorySlug = %q, want 'tech'", p.CategorySlug)
	}
}

func TestBuild_Excerpt_PlainText(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	// Description is empty — excerpt should come from raw Markdown body (no HTML tags)
	writePost(t, techDir, "2024-02-01-excerpted.md",
		"---\ntitle: Excerpted\ndate: 2024-02-01\nauthor: Bob\n---\n\n**Bold text** and _italic_ and normal words.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 1 {
		t.Fatalf("expected 1 post")
	}
	if posts[0].Excerpt == "" {
		t.Error("excerpt should not be empty")
	}
	// Excerpt should be plain text — no HTML or Markdown syntax
	if len(posts[0].Excerpt) > 0 {
		exc := posts[0].Excerpt
		if exc[0] == '<' {
			t.Errorf("excerpt should not start with HTML tag: %q", exc)
		}
	}
}

func TestBuild_Excerpt_UsesDescription(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-03-01-with-desc.md",
		"---\ntitle: With Desc\ndate: 2024-03-01\ndescription: My custom description\n---\n\nBody text that should not appear in excerpt.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 1 {
		t.Fatalf("expected 1 post")
	}
	if posts[0].Excerpt != "My custom description" {
		t.Errorf("Excerpt = %q, want 'My custom description'", posts[0].Excerpt)
	}
}

func TestBuild_SortedNewestFirst(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "older.md", "---\ntitle: Older\ndate: 2023-01-01\n---\nOld.")
	writePost(t, techDir, "newer.md", "---\ntitle: Newer\ndate: 2025-06-01\n---\nNew.")
	writePost(t, techDir, "middle.md", "---\ntitle: Middle\ndate: 2024-03-15\n---\nMid.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 3 {
		t.Fatalf("expected 3 posts, got %d", len(posts))
	}
	if posts[0].Date != "2025-06-01" {
		t.Errorf("first post should be newest, got %q", posts[0].Date)
	}
	if posts[2].Date != "2023-01-01" {
		t.Errorf("last post should be oldest, got %q", posts[2].Date)
	}
}

func TestBuild_MultipleCategoriesAndUncategorized(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.Config{
		PostsDir:          dir,
		PostIndexFile:     filepath.Join(dir, "posts.index.json"),
		DateFormat:        "2006-01-02",
		ShowUncategorized: true,
		Categories: map[string]config.Category{
			"alpha": {BlogName: "Alpha", Folder: "alpha"},
			"beta":  {BlogName: "Beta", Folder: "beta"},
		},
	}

	writePost(t, dir, "root-post.md", "---\ntitle: Root\ndate: 2024-01-01\n---\nRoot.")
	writePost(t, filepath.Join(dir, "alpha"), "a.md", "---\ntitle: A\ndate: 2024-02-01\n---\nA.")
	writePost(t, filepath.Join(dir, "beta"), "b.md", "---\ntitle: B\ndate: 2024-03-01\n---\nB.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 3 {
		t.Fatalf("expected 3 posts (1 root + 2 categorized), got %d", len(posts))
	}

	slugs := map[string]bool{}
	for _, p := range posts {
		slugs[p.Slug] = true
	}
	for _, s := range []string{"root-post", "a", "b"} {
		if !slugs[s] {
			t.Errorf("expected slug %q in index", s)
		}
	}
}

func TestBuild_EmptyDir(t *testing.T) {
	cfg := makeTestConfig(t)
	// No post files — should produce an empty index without error
	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 0 {
		t.Errorf("expected 0 posts, got %d", len(posts))
	}
}

func TestBuild_AtomicWrite(t *testing.T) {
	// A previous .tmp file should not interfere with a clean build
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "p.md", "---\ntitle: P\ndate: 2024-01-01\n---\nBody.")

	// Write stale .tmp file
	_ = os.WriteFile(cfg.PostIndexFile+".tmp", []byte("stale"), 0644)

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 1 {
		t.Errorf("expected 1 post, got %d", len(posts))
	}
}

func TestBuild_TitleFallback(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	// No title in front matter — should derive from filename
	writePost(t, techDir, "my-awesome-post.md", "---\ndate: 2024-01-01\n---\nContent.")

	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}

	posts := loadIndex(t, cfg.PostIndexFile)
	if len(posts) != 1 {
		t.Fatalf("expected 1 post")
	}
	if posts[0].Title == "" {
		t.Error("Title should be derived from filename when not provided")
	}
}
