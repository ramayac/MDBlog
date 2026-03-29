package markdown

import (
	"strings"
	"testing"
)

func TestParseFrontMatter(t *testing.T) {
	input := `---
title: Hello World
date: 2024-01-15
author: Jane Doe
tags: go, blog
description: A test post
---

# Hello

Some content here.
`
	doc := Parse(input)
	if doc.FrontMatter.Title != "Hello World" {
		t.Errorf("Title = %q, want %q", doc.FrontMatter.Title, "Hello World")
	}
	if doc.FrontMatter.Date != "2024-01-15" {
		t.Errorf("Date = %q, want %q", doc.FrontMatter.Date, "2024-01-15")
	}
	if doc.FrontMatter.Author != "Jane Doe" {
		t.Errorf("Author = %q, want %q", doc.FrontMatter.Author, "Jane Doe")
	}
	if doc.FrontMatter.Tags != "go, blog" {
		t.Errorf("Tags = %q, want %q", doc.FrontMatter.Tags, "go, blog")
	}
	if !strings.Contains(doc.HTML, "<h1") {
		t.Errorf("HTML should contain h1, got: %s", doc.HTML)
	}
	if !strings.Contains(doc.HTML, "Some content here") {
		t.Errorf("HTML should contain body text")
	}
}

func TestNoFrontMatter(t *testing.T) {
	input := "# Just a title\n\nSome content."
	doc := Parse(input)
	if doc.FrontMatter.Title != "" {
		t.Errorf("expected empty title, got %q", doc.FrontMatter.Title)
	}
	if !strings.Contains(doc.HTML, "Just a title") {
		t.Errorf("HTML should contain heading text")
	}
}

func TestParseMetaOnly(t *testing.T) {
	input := `---
title: Meta Only
date: 2024-05-01
---

Body text here.
`
	meta := ParseMetaOnly(input)
	if meta.FrontMatter.Title != "Meta Only" {
		t.Errorf("Title = %q, want %q", meta.FrontMatter.Title, "Meta Only")
	}
	if meta.Body == "" {
		t.Error("Body should not be empty")
	}
	if strings.Contains(meta.Body, "---") {
		t.Error("Body should not contain front matter delimiters")
	}
}

func TestSafeMode_NoRawHTML(t *testing.T) {
	// Goldmark's safe mode (default, no WithUnsafe) should strip raw HTML from Markdown
	input := `---
title: Safe
---

Hello <script>alert(1)</script> world.
`
	doc := Parse(input)
	if strings.Contains(doc.HTML, "<script>") {
		t.Error("raw HTML <script> should be escaped in safe mode")
	}
}

func TestLineBreaks(t *testing.T) {
	input := `---
title: Breaks
---

Line one
Line two
`
	doc := Parse(input)
	// WithHardWraps should produce <br> for single newlines within paragraphs
	if !strings.Contains(doc.HTML, "<br") {
		t.Errorf("expected <br> for line breaks, got HTML: %s", doc.HTML)
	}
}

func TestInvalidUTF8(t *testing.T) {
	// Should not panic on invalid byte sequences
	input := "---\ntitle: Bad\n---\n\nHello \xff world"
	doc := Parse(input)
	if doc.HTML == "" {
		t.Error("expected non-empty HTML even with invalid bytes")
	}
}

func TestParse_JSField(t *testing.T) {
	input := `---
title: With Script
js: myfile.js
---

Content.
`
	doc := Parse(input)
	if doc.FrontMatter.JS != "myfile.js" {
		t.Errorf("JS = %q, want 'myfile.js'", doc.FrontMatter.JS)
	}
}

func TestParse_ExtraFields(t *testing.T) {
	input := `---
title: Extra
custom_field: hello
another: world
---

Body.
`
	doc := ParseMetaOnly(input)
	if doc.FrontMatter.Extra["custom_field"] != "hello" {
		t.Errorf("Extra[custom_field] = %q, want 'hello'", doc.FrontMatter.Extra["custom_field"])
	}
	if doc.FrontMatter.Extra["another"] != "world" {
		t.Errorf("Extra[another] = %q, want 'world'", doc.FrontMatter.Extra["another"])
	}
}

func TestParse_MissingTitle(t *testing.T) {
	input := `---
date: 2024-01-01
---

Content.
`
	doc := Parse(input)
	if doc.FrontMatter.Title != "" {
		t.Errorf("expected empty title, got %q", doc.FrontMatter.Title)
	}
	if doc.FrontMatter.Date != "2024-01-01" {
		t.Errorf("Date = %q, want '2024-01-01'", doc.FrontMatter.Date)
	}
}

func TestParse_GFMCodeFence(t *testing.T) {
	input := "---\ntitle: Code\n---\n\n```go\nfmt.Println(\"hello\")\n```\n"
	doc := Parse(input)
	if !strings.Contains(doc.HTML, "<code") {
		t.Errorf("expected <code> block in HTML, got: %s", doc.HTML)
	}
	if !strings.Contains(doc.HTML, "fmt.Println") {
		t.Errorf("expected code content in HTML, got: %s", doc.HTML)
	}
}

func TestParse_GFMStrikethrough(t *testing.T) {
	input := "---\ntitle: Strike\n---\n\n~~deleted~~\n"
	doc := Parse(input)
	if !strings.Contains(doc.HTML, "<del>") {
		t.Errorf("expected <del> for strikethrough, got: %s", doc.HTML)
	}
}

func TestParse_XSSInBody(t *testing.T) {
	// Raw HTML tags should be suppressed in safe mode (Goldmark without WithUnsafe)
	input := "---\ntitle: XSS\n---\n\n<iframe src=\"evil.com\"></iframe>"
	doc := Parse(input)
	if strings.Contains(doc.HTML, "<iframe") {
		t.Error("iframe should be stripped in safe mode")
	}
}

func TestSplitFrontMatter_NonePresent(t *testing.T) {
	// Documents with no front matter block render as plain Markdown
	input := "Just some plain text\nwith no delimiters."
	doc := Parse(input)
	if !strings.Contains(doc.HTML, "plain text") {
		t.Errorf("expected body text in HTML, got: %s", doc.HTML)
	}
}
