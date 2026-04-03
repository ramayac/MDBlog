package buildfeed

import (
"encoding/xml"
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
		Lang:          "en",
		PostsDir:      dir,
		PostIndexFile: filepath.Join(dir, "posts.index.json"),
		DateFormat:    "2006-01-02",
		Categories: map[string]config.Category{
			"tech": {BlogName: "Tech Posts", Folder: "tech", Index: true},
		},
		Feed: config.FeedConfig{
			Enabled:     true,
			Title:       "Test Blog",
			Description: "A test blog.",
			BaseURL:     "https://example.com",
			MaxItems:    50,
			OutputFile:  filepath.Join(dir, "feed.xml"),
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

// testBuildFeed builds index then feed and returns the raw feed.xml as a string.
func testBuildFeed(t *testing.T, cfg *config.Config) string {
	t.Helper()
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex.Build: %v", err)
	}
	if err := Build(cfg); err != nil {
		t.Fatalf("buildfeed.Build: %v", err)
	}
	data, err := os.ReadFile(cfg.Feed.OutputFile)
	if err != nil {
		t.Fatalf("read feed.xml: %v", err)
	}
	return string(data)
}

func TestBuild_XMLDeclaration(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-post.md",
"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr := testBuildFeed(t, cfg)
	if !strings.HasPrefix(xmlStr, "<?xml") {
		t.Errorf("feed.xml should start with XML declaration, got: %q", xmlStr[:min(50, len(xmlStr))])
	}
}

func TestBuild_ChannelMetadata(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-hello-world.md",
"---\ntitle: Hello World\ndate: 2024-01-15\nauthor: Alice\ntags: go, test\ndescription: A test post about Go\n---\n\nBody content here.")

	xmlStr := testBuildFeed(t, cfg)

	checks := []struct{ name, want string }{
		{"title", "<title>Test Blog</title>"},
		{"link", "<link>https://example.com</link>"},
		{"description", "<description>A test blog.</description>"},
		{"language", "<language>en</language>"},
		{"atom:link href", `href="https://example.com/feed.xml"`},
		{"rss version", `version="2.0"`},
		{"atom namespace", `xmlns:atom="http://www.w3.org/2005/Atom"`},
	}
	for _, c := range checks {
		if !strings.Contains(xmlStr, c.want) {
			t.Errorf("%s: feed.xml should contain %q", c.name, c.want)
		}
	}
}

func TestBuild_Items(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-hello-world.md",
"---\ntitle: Hello World\ndate: 2024-01-15\nauthor: Alice\ntags: go, test\ndescription: A test post about Go\n---\n\nBody content here.")
	writePost(t, techDir, "2024-03-01-second-post.md",
"---\ntitle: Second Post\ndate: 2024-03-01\nauthor: Alice\ntags: go\ndescription: The second post\n---\n\nMore content.")

	xmlStr := testBuildFeed(t, cfg)

	// Parse items — item elements have no namespace issues
	type itemsOnly struct {
		XMLName xml.Name `xml:"rss"`
		Channel struct {
			Items []rssItem `xml:"channel>item"`
		}
	}
	// Use a simpler parse: just count <item> occurrences
	count := strings.Count(xmlStr, "<item>")
	if count != 2 {
		t.Errorf("expected 2 <item> elements, got %d", count)
	}

	// Newest post first (index is pre-sorted)
	idx1 := strings.Index(xmlStr, "Second Post")
	idx2 := strings.Index(xmlStr, "Hello World")
	if idx1 == -1 || idx2 == -1 {
		t.Error("both post titles should appear in feed.xml")
	} else if idx1 > idx2 {
		t.Error("newer post (Second Post) should appear before older post (Hello World)")
	}

	// Item link format
	expectedLink := `https://example.com/post?slug=2024-01-15-hello-world&amp;category=tech`
	if !strings.Contains(xmlStr, expectedLink) {
		t.Errorf("expected item link %q not found in feed.xml", expectedLink)
	}

	// Category name
	if !strings.Contains(xmlStr, "<category>Tech Posts</category>") {
		t.Error("expected category name 'Tech Posts' in feed.xml")
	}
}

func TestBuild_GUIDIsPermaLink(t *testing.T) {
	cfg := makeTestConfig(t)
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-post.md",
"---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")

	xmlStr := testBuildFeed(t, cfg)
	if !strings.Contains(xmlStr, `isPermaLink="true"`) {
		t.Error("guid should have isPermaLink=\"true\"")
	}
}

func TestBuild_MaxItems(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Feed.MaxItems = 1
	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-first.md",
"---\ntitle: First\ndate: 2024-01-15\n---\nContent.")
	writePost(t, techDir, "2024-02-01-second.md",
"---\ntitle: Second\ndate: 2024-02-01\n---\nContent.")

	xmlStr := testBuildFeed(t, cfg)
	count := strings.Count(xmlStr, "<item>")
	if count != 1 {
		t.Errorf("expected 1 item (max_items=1), got %d", count)
	}
}

func TestBuild_Disabled(t *testing.T) {
	cfg := makeTestConfig(t)
	cfg.Feed.Enabled = false

	techDir := filepath.Join(cfg.PostsDir, "tech")
	writePost(t, techDir, "2024-01-15-post.md", "---\ntitle: Post\ndate: 2024-01-15\n---\nContent.")
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex: %v", err)
	}
	if err := Build(cfg); err != nil {
		t.Fatalf("Build error: %v", err)
	}
	if _, err := os.Stat(cfg.Feed.OutputFile); !os.IsNotExist(err) {
		t.Errorf("expected feed.xml to not exist when disabled, but it was created")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
