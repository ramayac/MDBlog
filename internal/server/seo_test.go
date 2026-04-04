package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
)

// seoTestSetup creates a minimal Handler with Sitemap config for SEO-specific tests.
func seoTestSetup(t *testing.T) (*Handler, *config.Config) {
	t.Helper()
	base := testSetup(t)

	// Patch in Sitemap config pointing at the handler's temp dir
	dir := base.cfg.PostsDir
	base.cfg.Sitemap = config.SitemapConfig{
		Enabled:            true,
		OutputFile:         filepath.Join(dir, "sitemap.xml"),
		RobotsFile:         filepath.Join(dir, "robots.txt"),
		ChangeFreqHome:     "weekly",
		ChangeFreqCategory: "weekly",
		ChangeFreqPost:     "monthly",
		PriorityHome:       "1.0",
		PriorityCategory:   "0.8",
		PriorityPost:       "0.6",
	}
	return base, base.cfg
}

// ─────────────────────────────────────────────────────────────────────────────
// siteBaseURL
// ─────────────────────────────────────────────────────────────────────────────

func TestSiteBaseURL_TrimsTrailingSlash(t *testing.T) {
	cfg := &config.Config{Feed: config.FeedConfig{BaseURL: "https://example.com/"}}
	if got := siteBaseURL(cfg); got != "https://example.com" {
		t.Errorf("siteBaseURL = %q, want %q", got, "https://example.com")
	}
}

func TestSiteBaseURL_NoTrailingSlash(t *testing.T) {
	cfg := &config.Config{Feed: config.FeedConfig{BaseURL: "https://example.com"}}
	if got := siteBaseURL(cfg); got != "https://example.com" {
		t.Errorf("siteBaseURL = %q, want %q", got, "https://example.com")
	}
}

func TestSiteBaseURL_Empty(t *testing.T) {
	cfg := &config.Config{}
	if got := siteBaseURL(cfg); got != "" {
		t.Errorf("siteBaseURL empty input = %q, want empty string", got)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// marshalJSONLD
// ─────────────────────────────────────────────────────────────────────────────

func TestMarshalJSONLD_WrapsInScriptTag(t *testing.T) {
	data := map[string]any{"@type": "WebSite", "name": "Test"}
	result := string(marshalJSONLD(data))
	if !strings.HasPrefix(result, `<script type="application/ld+json">`) {
		t.Errorf("marshalJSONLD should start with script tag, got: %q", result[:min(60, len(result))])
	}
	if !strings.HasSuffix(result, `</script>`) {
		t.Errorf("marshalJSONLD should end with </script>, got: %q", result[max(0, len(result)-20):])
	}
}

func TestMarshalJSONLD_ValidJSON(t *testing.T) {
	data := map[string]any{"@context": "https://schema.org", "@type": "WebSite"}
	result := string(marshalJSONLD(data))
	// Extract JSON between the tags
	inner := strings.TrimPrefix(result, `<script type="application/ld+json">`)
	inner = strings.TrimSuffix(inner, `</script>`)
	var parsed map[string]any
	if err := json.Unmarshal([]byte(inner), &parsed); err != nil {
		t.Errorf("marshalJSONLD inner content is not valid JSON: %v", err)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// buildWebSiteJSONLD
// ─────────────────────────────────────────────────────────────────────────────

func TestBuildWebSiteJSONLD_RequiredFields(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	result := string(buildWebSiteJSONLD(cfg))
	for _, want := range []string{`"@type":"WebSite"`, `"name":"My Blog"`, `"url":"https://example.com"`, `"SearchAction"`} {
		if !strings.Contains(result, want) {
			t.Errorf("buildWebSiteJSONLD: expected %q in output", want)
		}
	}
}

func TestBuildWebSiteJSONLD_EmptyBaseURL(t *testing.T) {
	cfg := &config.Config{BlogName: "My Blog"}
	result := string(buildWebSiteJSONLD(cfg))
	if result != "" {
		t.Errorf("buildWebSiteJSONLD with empty base_url should return empty, got: %q", result)
	}
}

func TestBuildWebSiteJSONLD_IncludesDescription(t *testing.T) {
	cfg := &config.Config{
		BlogName:        "My Blog",
		BlogDescription: "A fun blog.",
		Feed:            config.FeedConfig{BaseURL: "https://example.com"},
	}
	result := string(buildWebSiteJSONLD(cfg))
	if !strings.Contains(result, `"description":"A fun blog."`) {
		t.Errorf("buildWebSiteJSONLD should include description when set")
	}
}

func TestBuildWebSiteJSONLD_OmitsDescriptionWhenEmpty(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	result := string(buildWebSiteJSONLD(cfg))
	if strings.Contains(result, `"description"`) {
		t.Errorf("buildWebSiteJSONLD should omit description when empty")
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// buildWebPageJSONLD
// ─────────────────────────────────────────────────────────────────────────────

func TestBuildWebPageJSONLD_RequiredFields(t *testing.T) {
	cfg := &config.Config{Feed: config.FeedConfig{BaseURL: "https://example.com"}}
	result := string(buildWebPageJSONLD(cfg, "https://example.com/?category=tech", "Tech Posts", "Tech description"))
	for _, want := range []string{`"@type":"WebPage"`, `"name":"Tech Posts"`, `"url":"https://example.com/?category=tech"`, `"description":"Tech description"`} {
		if !strings.Contains(result, want) {
			t.Errorf("buildWebPageJSONLD: expected %q in output", want)
		}
	}
}

func TestBuildWebPageJSONLD_BothEmpty(t *testing.T) {
	cfg := &config.Config{}
	result := string(buildWebPageJSONLD(cfg, "", "", ""))
	if result != "" {
		t.Errorf("buildWebPageJSONLD with empty base and canonical should return empty, got: %q", result)
	}
}

func TestBuildWebPageJSONLD_OmitsDescriptionWhenEmpty(t *testing.T) {
	cfg := &config.Config{Feed: config.FeedConfig{BaseURL: "https://example.com"}}
	result := string(buildWebPageJSONLD(cfg, "https://example.com/", "Home", ""))
	if strings.Contains(result, `"description"`) {
		t.Errorf("buildWebPageJSONLD should omit description when empty")
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// buildArticleJSONLD
// ─────────────────────────────────────────────────────────────────────────────

func TestBuildArticleJSONLD_NilPost(t *testing.T) {
	cfg := &config.Config{Feed: config.FeedConfig{BaseURL: "https://example.com"}}
	result := string(buildArticleJSONLD(cfg, nil, "", ""))
	if result != "" {
		t.Errorf("buildArticleJSONLD with nil post should return empty, got: %q", result)
	}
}

func TestBuildArticleJSONLD_TwoScriptBlocks(t *testing.T) {
	cfg := &config.Config{
		BlogName:   "My Blog",
		AuthorName: "Alice",
		Feed:       config.FeedConfig{BaseURL: "https://example.com"},
	}
	post := &blog.Post{
		Slug:  "my-post",
		Title: "My Post",
		Date:  "2024-01-15",
	}
	result := string(buildArticleJSONLD(cfg, post, "", "tech"))
	count := strings.Count(result, `<script type="application/ld+json">`)
	if count != 2 {
		t.Errorf("buildArticleJSONLD should produce 2 <script> blocks, got %d", count)
	}
}

func TestBuildArticleJSONLD_BlogPostingType(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	post := &blog.Post{
		Slug:  "my-post",
		Title: "My Post",
		Date:  "2024-01-15",
	}
	result := string(buildArticleJSONLD(cfg, post, "", ""))
	if !strings.Contains(result, `"BlogPosting"`) {
		t.Errorf("buildArticleJSONLD should contain BlogPosting type")
	}
	if !strings.Contains(result, `"BreadcrumbList"`) {
		t.Errorf("buildArticleJSONLD should contain BreadcrumbList type")
	}
}

func TestBuildArticleJSONLD_AuthorFallback(t *testing.T) {
	cfg := &config.Config{
		BlogName:   "My Blog",
		AuthorName: "DefaultAuthor",
		Feed:       config.FeedConfig{BaseURL: "https://example.com"},
	}
	post := &blog.Post{
		Slug:  "my-post",
		Title: "My Post",
		Date:  "2024-01-15",
		// No FrontMatter.Author set
	}
	result := string(buildArticleJSONLD(cfg, post, "", ""))
	if !strings.Contains(result, `"DefaultAuthor"`) {
		t.Errorf("buildArticleJSONLD should fall back to cfg.AuthorName when post has no author")
	}
}

func TestBuildArticleJSONLD_DescriptionFallback(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	post := &blog.Post{
		Slug:    "my-post",
		Title:   "My Post",
		Date:    "2024-01-15",
		Excerpt: "This is the excerpt.",
		// FrontMatter.Description is empty
	}
	result := string(buildArticleJSONLD(cfg, post, "", ""))
	if !strings.Contains(result, `"This is the excerpt."`) {
		t.Errorf("buildArticleJSONLD should fall back to post.Excerpt when description is empty")
	}
}

func TestBuildArticleJSONLD_Breadcrumb_NoCategory(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	post := &blog.Post{
		Slug:  "my-post",
		Title: "My Post",
		Date:  "2024-01-15",
		// Category is nil
	}
	result := string(buildArticleJSONLD(cfg, post, "", ""))
	// Without category there should be exactly 2 breadcrumb items (home + post).
	// We check by counting "ListItem" occurrences which appear once per item.
	count := strings.Count(result, `"ListItem"`)
	if count != 2 {
		t.Errorf("breadcrumb without category should have 2 ListItems, got %d", count)
	}
}

func TestBuildArticleJSONLD_Breadcrumb_WithCategory(t *testing.T) {
	cfg := &config.Config{
		BlogName: "My Blog",
		Feed:     config.FeedConfig{BaseURL: "https://example.com"},
	}
	cat := &blog.CategoryInfo{BlogName: "Tech Posts", Slug: "tech"}
	post := &blog.Post{
		Slug:         "my-post",
		Title:        "My Post",
		Date:         "2024-01-15",
		Category:     cat,
		CategorySlug: "tech",
	}
	result := string(buildArticleJSONLD(cfg, post, "", "tech"))
	// 3 items: home + category + post
	count := strings.Count(result, `"ListItem"`)
	if count != 3 {
		t.Errorf("breadcrumb with category should have 3 ListItems, got %d", count)
	}
	if !strings.Contains(result, `"Tech Posts"`) {
		t.Errorf("breadcrumb should include category name")
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// serveSitemap HTTP handler
// ─────────────────────────────────────────────────────────────────────────────

func TestServeSitemap_FromPrebuiltFile(t *testing.T) {
	h, cfg := seoTestSetup(t)
	sitemapContent := `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>`
	if err := os.WriteFile(cfg.Sitemap.OutputFile, []byte(sitemapContent), 0644); err != nil {
		t.Fatal(err)
	}

	w := get(h, "/sitemap.xml")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "application/xml") {
		t.Errorf("Content-Type = %q, want application/xml", ct)
	}
	if !strings.Contains(w.Body.String(), `xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"`) {
		t.Errorf("expected pre-built sitemap content in response")
	}
}

func TestServeSitemap_DynamicFallback(t *testing.T) {
	h, cfg := seoTestSetup(t)
	// Ensure the pre-built file does NOT exist
	_ = os.Remove(cfg.Sitemap.OutputFile)

	w := get(h, "/sitemap.xml")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "<?xml") {
		t.Errorf("dynamic sitemap fallback should return XML, got: %q", body[:min(100, len(body))])
	}
	if !strings.Contains(body, "https://example.com/") {
		t.Errorf("dynamic sitemap should include base URL")
	}
}

func TestServeSitemap_Disabled(t *testing.T) {
	h, cfg := seoTestSetup(t)
	cfg.Sitemap.Enabled = false

	w := get(h, "/sitemap.xml")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404 when sitemap disabled", w.Code)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// serveRobots HTTP handler
// ─────────────────────────────────────────────────────────────────────────────

func TestServeRobots_FromPrebuiltFile(t *testing.T) {
	h, cfg := seoTestSetup(t)
	robotsContent := "User-agent: *\nAllow: /\nSitemap: https://example.com/sitemap.xml\n"
	if err := os.WriteFile(cfg.Sitemap.RobotsFile, []byte(robotsContent), 0644); err != nil {
		t.Fatal(err)
	}

	w := get(h, "/robots.txt")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.Contains(ct, "text/plain") {
		t.Errorf("Content-Type = %q, want text/plain", ct)
	}
	if !strings.Contains(w.Body.String(), "User-agent: *") {
		t.Errorf("expected pre-built robots.txt content in response")
	}
}

func TestServeRobots_DynamicFallback(t *testing.T) {
	h, cfg := seoTestSetup(t)
	// Ensure pre-built file does NOT exist
	_ = os.Remove(cfg.Sitemap.RobotsFile)

	w := get(h, "/robots.txt")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "User-agent: *") {
		t.Errorf("dynamic robots.txt should contain 'User-agent: *'")
	}
	if !strings.Contains(body, "Allow: /") {
		t.Errorf("dynamic robots.txt should contain 'Allow: /'")
	}
}

func TestServeRobots_DynamicFallback_IncludesSitemap(t *testing.T) {
	h, cfg := seoTestSetup(t)
	_ = os.Remove(cfg.Sitemap.RobotsFile)

	w := get(h, "/robots.txt")
	body := w.Body.String()
	if !strings.Contains(body, "Sitemap: https://example.com/") {
		t.Errorf("dynamic robots.txt should include Sitemap line, got:\n%s", body)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
