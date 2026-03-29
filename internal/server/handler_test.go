package server

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
)

// testSetup creates a minimal temp dir with posts and returns a ready Handler.
func testSetup(t *testing.T) *Handler {
	t.Helper()

	dir := t.TempDir()

	// Create srbyte category
	srbyteDir := dir + "/srbyte"
	if err := os.MkdirAll(srbyteDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Write a known post
	postContent := `---
title: 12:34:56 7 8 9 y el tiempo
date: 2008-01-01
author: Rodrigo Amaya
tags: tiempo, linux
description: A post about time
---

# 12:34:56 7 8 9 y el tiempo

El tiempo es relativo.
`
	if err := os.WriteFile(srbyteDir+"/srbyte-12-34-56-7-8-9-y-el-tiempo.md", []byte(postContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Write a second post
	post2 := `---
title: Linux Commands
date: 2009-01-01
author: Rodrigo Amaya
tags: linux
description: Useful linux commands
---

# Linux Commands

Many useful commands.
`
	if err := os.WriteFile(srbyteDir+"/linux-commands.md", []byte(post2), 0644); err != nil {
		t.Fatal(err)
	}

	cfg := &config.Config{
		BlogName:               "Rodrigo A.",
		AuthorName:             "Rodrigo Amaya",
		Lang:                   "en",
		HeaderContent:          "Wholesome Software Development.",
		PostsPerPage:           25,
		ExcerptLength:          200,
		ShowUncategorized:      false,
		PostsDir:               dir,
		PostIndexFile:          dir + "/posts.index.json",
		DateFormat:             "2006-01-02",
		DefaultMetaDescription: "Test blog.",
		CSSTheme:               "assets/css/default.style.css",
		CSP: config.CSPConfig{
			Enabled: true,
			Header:  "Content-Security-Policy: default-src 'self';",
		},
		MenuLinks: []config.MenuLink{{Label: "Home", URL: "index.php"}},
		Categories: map[string]config.Category{
			"srbyte": {BlogName: "Sr. Byte 👨‍💻", Folder: "srbyte", Index: false, Menu: true},
		},
		Labels: config.Labels{
			ReadMore:             "Read more →",
			BackToAll:            "← Back to all posts",
			BackToCategory:       "← Back to %s",
			NotFoundTitle:        "404 — Post Not Found",
			NotFoundMessage:      "The post you're looking for doesn't exist.",
			NoPostsInCategory:    "No posts found in this category.",
			PaginationPrev:       "← Newer Posts",
			PaginationNext:       "Older Posts →",
			PageIndicator:        "Page %d of %d",
			AuthorBy:             "By %s",
			SearchTitle:          "Search",
			SearchPlaceholder:    "What are you looking for?",
			SearchButton:         "🔍 Search",
			SearchShowingResults: `Showing results for "%s"`,
			SearchEmptyQuery:     "Enter a keyword above to search through posts.",
			SearchNoResults:      "No posts found matching your query.",
			SearchResultsTitle:   `Search Results for "%s"`,
		},
	}

	// Build the post index
	if err := buildindex.Build(cfg); err != nil {
		t.Fatalf("buildindex: %v", err)
	}

	b := blog.New(cfg)

	// Point template FS at our real templates directory
	TemplateFS = os.DirFS("../../templates")
	// Verify it's readable
	if _, err := fs.Stat(TemplateFS, "layout.html"); err != nil {
		t.Skipf("templates not found (running outside repo root): %v", err)
	}

	return New(cfg, b)
}

func get(h *Handler, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func TestHome(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Rodrigo A.") {
		t.Error("should contain blog name")
	}
	if !strings.Contains(body, "srbyte") {
		t.Error("should contain category card for srbyte")
	}
	if strings.Contains(body, "What are you looking for?") {
		t.Error("should NOT contain search form on home page")
	}
}

func TestCategory(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?category=srbyte")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Read more") {
		t.Error("should show post list with read-more links")
	}
}

func TestSearchLayout(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?q=linux")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Search Results") {
		t.Errorf("should contain Search Results heading, got excerpt: %s", body[:500])
	}
	if !strings.Contains(body, "What are you looking for?") {
		t.Error("should contain search form")
	}
}

func TestSearchWorking(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/?q=tiempo")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Read more") {
		t.Error("should find matching post")
	}
	if !strings.Contains(body, "12:34:56") {
		t.Errorf("should find the specific post, body excerpt: %s", body[:500])
	}
}

func TestPost(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/post.php?slug=srbyte-12-34-56-7-8-9-y-el-tiempo&category=srbyte")
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "12:34:56") {
		t.Error("post title should be displayed")
	}
	if !strings.Contains(body, "Rodrigo Amaya") {
		t.Error("author should be displayed")
	}
}

func Test404(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/post.php?slug=does-not-exist")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "404") {
		t.Error("should contain 404 in page")
	}
	if !strings.Contains(body, "Not Found") {
		t.Error("should contain 'Not Found' in page")
	}
}

func TestAssetPathTraversal(t *testing.T) {
	h := testSetup(t)
	for _, path := range []string{
		"/assets/../config.toml",
		"/assets/%2e%2e/config.toml",
	} {
		w := get(h, path)
		if w.Code == http.StatusOK {
			t.Errorf("path %q should be rejected, got 200", path)
		}
	}
}

func TestCSPHeader(t *testing.T) {
	h := testSetup(t)
	w := get(h, "/")
	if csp := w.Header().Get("Content-Security-Policy"); csp == "" {
		t.Error("CSP header should be set")
	}
}
