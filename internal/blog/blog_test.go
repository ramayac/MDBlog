package blog

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ramayac/mdblog/internal/config"
)

func TestGenerateSlug(t *testing.T) {
	cases := []struct{ in, want string }{
		{"2024-01-15-hello-world.md", "2024-01-15-hello-world"},
		{"srbyte-12-34-56-7-8-9-y-el-tiempo.md", "srbyte-12-34-56-7-8-9-y-el-tiempo"},
		{"My_Post File.md", "my-post-file"},
	}
	for _, c := range cases {
		got := generateSlug(c.in)
		if got != c.want {
			t.Errorf("generateSlug(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTitleFromFilename(t *testing.T) {
	cases := []struct{ in, want string }{
		{"2024-01-15-hello-world.md", "Hello world"},
		{"my-post.md", "My post"},
	}
	for _, c := range cases {
		got := titleFromFilename(c.in)
		if got != c.want {
			t.Errorf("titleFromFilename(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGenerateExcerpt(t *testing.T) {
	html := "<p>Hello <strong>world</strong>, this is a test of the excerpt generator function.</p>"
	got := generateExcerpt(html, 20)
	if len(got) > 23 { // 20 chars + "..."
		t.Errorf("excerpt too long: %q", got)
	}
	if got == "" {
		t.Error("excerpt should not be empty")
	}
}

func makeTestConfig(postsDir string) *config.Config {
	return &config.Config{
		BlogName:          "Test",
		PostsDir:          postsDir,
		PostIndexFile:     filepath.Join(postsDir, "posts.index.json"),
		PostsPerPage:      10,
		ExcerptLength:     200,
		DateFormat:        "2006-01-02",
		ShowUncategorized: true,
		Categories: map[string]config.Category{
			"tech": {BlogName: "Tech", Folder: "tech", Index: true, Menu: true},
		},
	}
}

func writePost(t *testing.T, dir, filename, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, filename), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGetPostBySlug(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "2024-01-01-test-post.md", "---\ntitle: Test Post\ndate: 2024-01-01\nauthor: Alice\n---\n\nHello world.")

	cfg := makeTestConfig(dir)
	b := New(cfg)

	post := b.GetPostBySlug("2024-01-01-test-post", "")
	if post == nil {
		t.Fatal("expected post, got nil")
	}
	if post.Title != "Test Post" {
		t.Errorf("Title = %q, want %q", post.Title, "Test Post")
	}
	if post.Content == "" {
		t.Error("Content should not be empty")
	}
}

func TestGetPostBySlug_PathTraversal(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)

	for _, slug := range []string{"../etc/passwd", "foo/bar", `foo\bar`} {
		if b.GetPostBySlug(slug, "") != nil {
			t.Errorf("slug %q should be rejected for path traversal", slug)
		}
	}
}

func TestGetPosts_Scan(t *testing.T) {
	dir := t.TempDir()
	writePost(t, dir, "2024-01-01-alpha.md", "---\ntitle: Alpha\ndate: 2024-01-01\n---\n\nAlpha content.")
	writePost(t, dir, "2024-01-02-beta.md", "---\ntitle: Beta\ndate: 2024-01-02\n---\n\nBeta content.")

	cfg := makeTestConfig(dir)
	cfg.PostIndexFile = filepath.Join(dir, "nonexistent.json") // force filesystem scan
	b := New(cfg)

	list := b.GetPosts(1, "")
	if len(list.Posts) != 2 {
		t.Errorf("expected 2 posts, got %d", len(list.Posts))
	}
	// Newest first
	if list.Posts[0].Date != "2024-01-02" {
		t.Errorf("first post date = %q, want newest 2024-01-02", list.Posts[0].Date)
	}
}

func TestSearchPosts_NoIndex(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.PostIndexFile = filepath.Join(t.TempDir(), "nonexistent.json")
	b := New(cfg)

	list := b.SearchPosts("anything", 1)
	if len(list.Posts) != 0 {
		t.Error("expected no results without index")
	}
}

func TestGetMenu(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.MenuLinks = []config.MenuLink{{Label: "Home", URL: "index.php"}}
	b := New(cfg)

	menu := b.GetMenu()
	if len(menu) < 2 { // at least Home + theme toggle
		t.Errorf("expected at least 2 menu items, got %d", len(menu))
	}
	if menu[0].Label != "Home" {
		t.Errorf("first menu item = %q, want 'Home'", menu[0].Label)
	}
	last := menu[len(menu)-1]
	if last.Label != "🌓" {
		t.Errorf("last menu item = %q, want theme toggle", last.Label)
	}
}
