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
			"tech": {BlogName: "Tech", Folder: "tech", Index: true},
		},
		Menu: config.MenuConfig{
			Categories: config.MenuDropdown{
				Item: []config.MenuCategoryRef{
					{Category: "tech", Order: 1},
				},
			},
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
	cfg.MenuLinks = []config.MenuLink{{Label: "Home", URL: "/"}}
	cfg.Menu.Categories.Label = "Writings"
	b := New(cfg)

	menu := b.GetMenu()

	// categories are now grouped into a single dropdown item
	var dropdown *MenuLink
	for i := range menu {
		if menu[i].Label == "Writings" {
			dropdown = &menu[i]
			break
		}
	}
	if dropdown == nil {
		t.Fatal("expected a 'Writings' dropdown item in menu")
	}
	if len(dropdown.SubItems) == 0 {
		t.Fatal("expected SubItems in Writings dropdown")
	}
	var techURL string
	for _, sub := range dropdown.SubItems {
		if sub.Label == "Tech" {
			techURL = sub.URL
		}
	}
	if techURL != "/?category=tech" {
		t.Errorf("Tech URL = %q, want /?category=tech", techURL)
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

// ─────────────────────────────────────────────────────────────────────────────
// GetPage tests
// ─────────────────────────────────────────────────────────────────────────────

func makeTestConfigWithPages(postsDir, pagesDir string) *config.Config {
	cfg := makeTestConfig(postsDir)
	cfg.PagesDir = pagesDir
	return cfg
}

func writePage(t *testing.T, dir, filename, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, filename), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGetPage_Found(t *testing.T) {
	postsDir := t.TempDir()
	pagesDir := t.TempDir()
	writePage(t, pagesDir, "about.md", "---\ntitle: About Me\ndescription: Learn about me.\n---\n\nHello, I'm a person.")

	cfg := makeTestConfigWithPages(postsDir, pagesDir)
	b := New(cfg)

	page := b.GetPage("about")
	if page == nil {
		t.Fatal("expected page, got nil")
	}
	if page.Title != "About Me" {
		t.Errorf("Title = %q, want 'About Me'", page.Title)
	}
	if page.Slug != "about" {
		t.Errorf("Slug = %q, want 'about'", page.Slug)
	}
	if page.Content == "" {
		t.Error("Content should not be empty")
	}
	if !strings.Contains(page.Content, "Hello") {
		t.Errorf("Content should contain rendered body, got: %s", page.Content)
	}
}

func TestGetPage_FrontMatterDescription(t *testing.T) {
	pagesDir := t.TempDir()
	writePage(t, pagesDir, "contact.md", "---\ntitle: Contact\ndescription: Get in touch.\n---\nBody.")

	cfg := makeTestConfigWithPages(t.TempDir(), pagesDir)
	b := New(cfg)

	page := b.GetPage("contact")
	if page == nil {
		t.Fatal("expected page, got nil")
	}
	if page.FrontMatter.Description != "Get in touch." {
		t.Errorf("Description = %q, want 'Get in touch.'", page.FrontMatter.Description)
	}
}

func TestGetPage_TitleFallsBackToSlug(t *testing.T) {
	pagesDir := t.TempDir()
	// No title in front matter
	writePage(t, pagesDir, "mypage.md", "---\n---\n\nContent here.")

	cfg := makeTestConfigWithPages(t.TempDir(), pagesDir)
	b := New(cfg)

	page := b.GetPage("mypage")
	if page == nil {
		t.Fatal("expected page, got nil")
	}
	if page.Title != "mypage" {
		t.Errorf("Title = %q, want slug fallback 'mypage'", page.Title)
	}
}

func TestGetPage_NotFound(t *testing.T) {
	cfg := makeTestConfigWithPages(t.TempDir(), t.TempDir())
	b := New(cfg)

	if b.GetPage("does-not-exist") != nil {
		t.Error("expected nil for missing page")
	}
}

func TestGetPage_PathTraversal(t *testing.T) {
	cfg := makeTestConfigWithPages(t.TempDir(), t.TempDir())
	b := New(cfg)

	for _, slug := range []string{"../etc/passwd", "foo/bar", `foo\bar`, "../../secret"} {
		if b.GetPage(slug) != nil {
			t.Errorf("slug %q should be rejected for path traversal", slug)
		}
	}
}

func TestGetPage_RendersMarkdown(t *testing.T) {
	pagesDir := t.TempDir()
	writePage(t, pagesDir, "projects.md", "---\ntitle: Projects\n---\n\n**Bold text** and a [link](https://example.com).")

	cfg := makeTestConfigWithPages(t.TempDir(), pagesDir)
	b := New(cfg)

	page := b.GetPage("projects")
	if page == nil {
		t.Fatal("expected page, got nil")
	}
	if !strings.Contains(page.Content, "<strong>Bold text</strong>") {
		t.Errorf("expected rendered bold, got: %s", page.Content)
	}
	if !strings.Contains(page.Content, "<a ") {
		t.Errorf("expected rendered link, got: %s", page.Content)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// GetMenu dropdown tests
// ─────────────────────────────────────────────────────────────────────────────

func TestGetMenu_PinnedLinksAreInline(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.Menu.Pinned = []config.MenuCategoryRef{
		{Category: "tech", Order: 1},
	}
	b := New(cfg)

	menu := b.GetMenu()
	// Pinned items must appear as top-level links (no SubItems)
	for _, item := range menu {
		if item.Label == "Tech" {
			if len(item.SubItems) != 0 {
				t.Error("pinned item should not have SubItems")
			}
			if item.URL != "/?category=tech" {
				t.Errorf("pinned URL = %q, want /?category=tech", item.URL)
			}
			return
		}
	}
	t.Error("expected 'Tech' pinned item in top-level menu")
}

func TestGetMenu_NoCategoriesNoDropdown(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.Menu.Categories.Item = nil // no dropdown categories
	b := New(cfg)

	menu := b.GetMenu()
	for _, item := range menu {
		if len(item.SubItems) > 0 {
			t.Errorf("expected no dropdown when categories are empty, found item %q with SubItems", item.Label)
		}
	}
}

func TestGetMenu_DropdownLabelFallback(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.Menu.Categories.Label = "" // empty — should fall back to "More"
	b := New(cfg)

	menu := b.GetMenu()
	for _, item := range menu {
		if len(item.SubItems) > 0 {
			if item.Label != "More" {
				t.Errorf("dropdown label = %q, want 'More' fallback", item.Label)
			}
			return
		}
	}
	t.Error("expected a dropdown item in menu")
}

func TestGetMenu_StaticLinksOrderedFirst(t *testing.T) {
	cfg := makeTestConfig(t.TempDir())
	cfg.MenuLinks = []config.MenuLink{
		{Label: "Home", URL: "/"},
		{Label: "About", URL: "/page?slug=about"},
	}
	b := New(cfg)

	menu := b.GetMenu()
	if len(menu) < 2 {
		t.Fatal("expected at least 2 menu items")
	}
	if menu[0].Label != "Home" {
		t.Errorf("first item = %q, want 'Home'", menu[0].Label)
	}
	if menu[1].Label != "About" {
		t.Errorf("second item = %q, want 'About'", menu[1].Label)
	}
}

func TestGetMenu_DropdownSubItemsOrdered(t *testing.T) {
	cfg := &config.Config{
		PostsDir:      t.TempDir(),
		PostsPerPage:  10,
		ExcerptLength: 200,
		DateFormat:    "2006-01-02",
		Categories: map[string]config.Category{
			"aaa": {BlogName: "AAA", Folder: "aaa"},
			"bbb": {BlogName: "BBB", Folder: "bbb"},
			"ccc": {BlogName: "CCC", Folder: "ccc"},
		},
		Menu: config.MenuConfig{
			Categories: config.MenuDropdown{
				Label: "Writings",
				Item: []config.MenuCategoryRef{
					{Category: "ccc", Order: 3},
					{Category: "aaa", Order: 1},
					{Category: "bbb", Order: 2},
				},
			},
		},
	}
	b := New(cfg)

	menu := b.GetMenu()
	var dropdown *MenuLink
	for i := range menu {
		if menu[i].Label == "Writings" {
			dropdown = &menu[i]
			break
		}
	}
	if dropdown == nil {
		t.Fatal("expected 'Writings' dropdown")
	}
	if len(dropdown.SubItems) != 3 {
		t.Fatalf("expected 3 sub-items, got %d", len(dropdown.SubItems))
	}
	want := []string{"AAA", "BBB", "CCC"}
	for i, w := range want {
		if dropdown.SubItems[i].Label != w {
			t.Errorf("SubItems[%d].Label = %q, want %q", i, dropdown.SubItems[i].Label, w)
		}
	}
}
