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
