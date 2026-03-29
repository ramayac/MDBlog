package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func TestGetMenu_CategoryLinks(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.MenuLinks = []config.MenuLink{{Label: "Home", URL: "index.php"}}
	b := New(cfg)

	menu := b.GetMenu()
	labels := make(map[string]string)
	for _, m := range menu {
		labels[m.Label] = m.URL
	}
	if _, ok := labels["Tech"]; !ok {
		t.Error("expected 'Tech' category link in menu")
	}
	if labels["Tech"] != "index.php?category=tech" {
		t.Errorf("Tech URL = %q, want index.php?category=tech", labels["Tech"])
	}
}

func TestGetCategories_WithPosts(t *testing.T) {
	dir := t.TempDir()
	techDir := filepath.Join(dir, "tech")
	writePost(t, techDir, "my-post.md", "---\ntitle: A\n---\nContent.")

	cfg := makeTestConfig(dir)
	b := New(cfg)

	cats := b.GetCategories()
	if _, ok := cats["tech"]; !ok {
		t.Fatal("expected 'tech' in categories")
	}
	if cats["tech"].Count != 1 {
		t.Errorf("Count = %d, want 1", cats["tech"].Count)
	}
}

func TestGetCategories_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	// tech folder exists but has no .md files
	if err := os.MkdirAll(filepath.Join(dir, "tech"), 0755); err != nil {
		t.Fatal(err)
	}
	cfg := makeTestConfig(dir)
	b := New(cfg)

	cats := b.GetCategories()
	if _, ok := cats["tech"]; ok {
		t.Error("expected empty category to be excluded")
	}
}

func TestGetCategories_CachingIsIdempotent(t *testing.T) {
	dir := t.TempDir()
	writePost(t, filepath.Join(dir, "tech"), "p.md", "---\ntitle: T\n---\nBody.")
	cfg := makeTestConfig(dir)
	b := New(cfg)

	c1 := b.GetCategories()
	c2 := b.GetCategories()
	if len(c1) != len(c2) {
		t.Error("GetCategories should return same result on second call")
	}
}

func TestGetCategoryBySlug_Found(t *testing.T) {
	dir := t.TempDir()
	writePost(t, filepath.Join(dir, "tech"), "p.md", "---\ntitle: T\n---\nBody.")
	cfg := makeTestConfig(dir)
	b := New(cfg)

	cat := b.GetCategoryBySlug("tech")
	if cat == nil {
		t.Fatal("expected category, got nil")
	}
	if cat.Slug != "tech" {
		t.Errorf("Slug = %q, want 'tech'", cat.Slug)
	}
}

func TestGetCategoryBySlug_NotFound(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)

	if b.GetCategoryBySlug("nonexistent") != nil {
		t.Error("expected nil for unknown slug")
	}
}

func TestParseMarkdown(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)
	html := b.ParseMarkdown("**bold** text")
	if html == "" {
		t.Error("expected non-empty HTML")
	}
	if !strings.Contains(html, "<strong>bold</strong>") {
		t.Errorf("expected <strong>, got: %s", html)
	}
}

func TestGetVersionInfo(t *testing.T) {
	BuildVersion = "v1.2.3"
	BuildCommit = "abc123"
	BuildDate = "2024-01-01"
	t.Cleanup(func() { BuildVersion = ""; BuildCommit = ""; BuildDate = "" })

	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)
	vi := b.GetVersionInfo()

	if vi.Version != "v1.2.3" {
		t.Errorf("Version = %q, want v1.2.3", vi.Version)
	}
	if vi.Commit != "abc123" {
		t.Errorf("Commit = %q, want abc123", vi.Commit)
	}
}

func TestGetConfig(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)
	if b.GetConfig() != cfg {
		t.Error("GetConfig should return the same pointer")
	}
}

func TestBuildPagination(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)

	p := b.buildPagination(2, 5)
	if p.Current != 2 || p.Total != 5 {
		t.Errorf("Current/Total = %d/%d, want 2/5", p.Current, p.Total)
	}
	if !p.HasNext || !p.HasPrev {
		t.Error("page 2 of 5 should have both prev and next")
	}
	if p.Next != 3 || p.Prev != 1 {
		t.Errorf("Next/Prev = %d/%d, want 3/1", p.Next, p.Prev)
	}

	p1 := b.buildPagination(1, 1)
	if p1.HasNext || p1.HasPrev {
		t.Error("single page should have no next/prev")
	}
}

func TestGetPosts_WithIndex(t *testing.T) {
	dir := t.TempDir()
	techDir := filepath.Join(dir, "tech")
	writePost(t, techDir, "2024-01-01-first.md", "---\ntitle: First\ndate: 2024-01-01\nauthor: A\ntags: go\ndescription: desc\n---\nText.")
	writePost(t, techDir, "2024-06-01-second.md", "---\ntitle: Second\ndate: 2024-06-01\nauthor: A\ntags: go\ndescription: desc2\n---\nText2.")

	cfg := makeTestConfig(dir)

	// Build the index so GetPosts can use it
	import_buildindex := func() {
		// Write a minimal hand-crafted index to avoid importing buildindex (cycle risk)
		idx := `[
			{"slug":"2024-06-01-second","title":"Second","date":"2024-06-01","author":"A","tags":"go","description":"desc2","excerpt":"Text2.","category_slug":"tech","source_path":"tech/2024-06-01-second.md","filename":"2024-06-01-second.md"},
			{"slug":"2024-01-01-first","title":"First","date":"2024-01-01","author":"A","tags":"go","description":"desc","excerpt":"Text.","category_slug":"tech","source_path":"tech/2024-01-01-first.md","filename":"2024-01-01-first.md"}
		]`
		_ = os.WriteFile(cfg.PostIndexFile, []byte(idx), 0644)
	}
	import_buildindex()

	b := New(cfg)
	list := b.GetPosts(1, "tech")

	if len(list.Posts) != 2 {
		t.Fatalf("expected 2 posts, got %d", len(list.Posts))
	}
	if list.Posts[0].Date != "2024-06-01" {
		t.Errorf("first post date = %q, want newest first", list.Posts[0].Date)
	}
}

func TestGetPosts_Pagination(t *testing.T) {
	dir := t.TempDir()
	cfg := makeTestConfig(dir)
	cfg.PostsPerPage = 1
	cfg.PostIndexFile = filepath.Join(dir, "nonexistent.json") // force scan

	techDir := filepath.Join(dir, "tech")
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("2024-01-%02d-post-%d.md", i+1, i)
		content := fmt.Sprintf("---\ntitle: Post %d\ndate: 2024-01-%02d\n---\nContent.", i, i+1)
		writePost(t, techDir, name, content)
	}

	b := New(cfg)
	page1 := b.GetPosts(1, "tech")
	if len(page1.Posts) != 1 {
		t.Errorf("page 1: expected 1 post, got %d", len(page1.Posts))
	}
	if !page1.Pagination.HasNext {
		t.Error("page 1 should have next page")
	}
	if page1.Pagination.HasPrev {
		t.Error("page 1 should not have prev page")
	}
}

func TestSearchPosts_Matches(t *testing.T) {
	dir := t.TempDir()
	cfg := makeTestConfig(dir)
	idx := `[
		{"slug":"a","title":"Go Programming","date":"2024-01-01","tags":"go","description":"","excerpt":"Learn Go today.","category_slug":"tech","filename":"a.md","source_path":"tech/a.md"},
		{"slug":"b","title":"Python Basics","date":"2024-01-02","tags":"python","description":"","excerpt":"Learn Python.","category_slug":"tech","filename":"b.md","source_path":"tech/b.md"}
	]`
	_ = os.WriteFile(cfg.PostIndexFile, []byte(idx), 0644)

	b := New(cfg)
	list := b.SearchPosts("go", 1)

	if len(list.Posts) != 1 {
		t.Fatalf("expected 1 result for 'go', got %d", len(list.Posts))
	}
	if list.Posts[0].Title != "Go Programming" {
		t.Errorf("unexpected match: %q", list.Posts[0].Title)
	}
	if list.TotalMatches != 1 {
		t.Errorf("TotalMatches = %d, want 1", list.TotalMatches)
	}
}

func TestSearchPosts_EmptyQuery(t *testing.T) {
	dir := t.TempDir()
	cfg := makeTestConfig(dir)
	_ = os.WriteFile(cfg.PostIndexFile, []byte(`[{"slug":"a","title":"A","date":"2024-01-01","filename":"a.md"}]`), 0644)
	b := New(cfg)

	list := b.SearchPosts("", 1)
	if len(list.Posts) != 0 {
		t.Error("empty query should return no results")
	}
}

func TestSearchPosts_CaseInsensitive(t *testing.T) {
	dir := t.TempDir()
	cfg := makeTestConfig(dir)
	idx := `[{"slug":"a","title":"Golang Rocks","date":"2024-01-01","tags":"","description":"","excerpt":"","category_slug":"tech","filename":"a.md","source_path":"tech/a.md"}]`
	_ = os.WriteFile(cfg.PostIndexFile, []byte(idx), 0644)
	b := New(cfg)

	list := b.SearchPosts("GOLANG", 1)
	if len(list.Posts) != 1 {
		t.Errorf("case-insensitive search for 'GOLANG' should match, got %d results", len(list.Posts))
	}
}

func TestGetPostBySlug_WithCategory(t *testing.T) {
	dir := t.TempDir()
	techDir := filepath.Join(dir, "tech")
	writePost(t, techDir, "my-article.md", "---\ntitle: My Article\ndate: 2024-03-01\nauthor: Bob\n---\n\nContent here.")

	cfg := makeTestConfig(dir)
	b := New(cfg)

	post := b.GetPostBySlug("my-article", "tech")
	if post == nil {
		t.Fatal("expected post, got nil")
	}
	if post.Title != "My Article" {
		t.Errorf("Title = %q, want 'My Article'", post.Title)
	}
	if post.CategorySlug != "tech" {
		t.Errorf("CategorySlug = %q, want 'tech'", post.CategorySlug)
	}
	if post.Category == nil {
		t.Error("Category should be populated")
	}
}

func TestGetPostBySlug_UnknownCategory(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)

	if post := b.GetPostBySlug("anything", "unknown-cat"); post != nil {
		t.Error("expected nil for unknown category")
	}
}

func TestGetPostBySlug_NotFound(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	b := New(cfg)

	if post := b.GetPostBySlug("does-not-exist", ""); post != nil {
		t.Error("expected nil for missing post")
	}
}

func TestGetPostBySlug_DateFallback(t *testing.T) {
	// Post with no date in front matter — should fall back to file mtime
	dir := t.TempDir()
	writePost(t, dir, "no-date.md", "---\ntitle: No Date\n---\n\nContent.")

	cfg := makeTestConfig(dir)
	b := New(cfg)

	post := b.GetPostBySlug("no-date", "")
	if post == nil {
		t.Fatal("expected post")
	}
	if post.Date == "" {
		t.Error("Date should be populated from mtime when not in front matter")
	}
}

func TestGetPostBySlug_ResolveViaIndex(t *testing.T) {
	dir := t.TempDir()
	techDir := filepath.Join(dir, "tech")
	writePost(t, techDir, "indexed-post.md", "---\ntitle: Indexed\ndate: 2024-01-01\n---\nBody.")

	cfg := makeTestConfig(dir)
	// Write an index pointing to the post
	idx := `[{"slug":"indexed-post","title":"Indexed","date":"2024-01-01","author":"","tags":"","description":"","excerpt":"Body.","category_slug":"tech","source_path":"tech/indexed-post.md","filename":"indexed-post.md"}]`
	_ = os.WriteFile(cfg.PostIndexFile, []byte(idx), 0644)

	b := New(cfg)
	// Request without a category slug — should resolve via index
	post := b.GetPostBySlug("indexed-post", "")
	if post == nil {
		t.Fatal("expected post resolved via index, got nil")
	}
	if post.Title != "Indexed" {
		t.Errorf("Title = %q, want 'Indexed'", post.Title)
	}
}

func TestSortPostsByDate(t *testing.T) {
	posts := []Post{
		{Date: "2023-01-01"},
		{Date: "2025-06-15"},
		{Date: "2024-03-20"},
	}
	sortPostsByDate(posts)
	if posts[0].Date != "2025-06-15" {
		t.Errorf("expected newest first, got %q", posts[0].Date)
	}
	if posts[2].Date != "2023-01-01" {
		t.Errorf("expected oldest last, got %q", posts[2].Date)
	}
}

func TestGenerateExcerpt_Short(t *testing.T) {
	html := "<p>Short.</p>"
	got := generateExcerpt(html, 200)
	if got != "Short." {
		t.Errorf("short text should be returned as-is, got %q", got)
	}
}

func TestGenerateExcerpt_WordBoundary(t *testing.T) {
	html := "<p>word1 word2 word3 word4 word5</p>"
	got := generateExcerpt(html, 12)
	// 12 chars: "word1 word2 " — truncation should happen at word boundary
	if strings.HasSuffix(got, " ...") {
		t.Errorf("trailing space before ellipsis: %q", got)
	}
	if !strings.HasSuffix(got, "...") {
		t.Errorf("expected trailing ellipsis, got %q", got)
	}
}

func TestMax1(t *testing.T) {
	if max1(0) != 1 {
		t.Error("max1(0) should return 1")
	}
	if max1(-5) != 1 {
		t.Error("max1(-5) should return 1")
	}
	if max1(3) != 3 {
		t.Error("max1(3) should return 3")
	}
}
