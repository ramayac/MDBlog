package markdown

import (
	"bytes"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// FrontMatter holds the parsed YAML-like header of a Markdown post.
type FrontMatter struct {
	Title       string
	Date        string
	Author      string
	Tags        string // raw comma-separated or single value
	Description string
	JS          string // optional JS filename
	Extra       map[string]string
}

// ParsedDoc is the result of parsing a complete Markdown post file.
type ParsedDoc struct {
	FrontMatter FrontMatter
	HTML        string
}

// MetaDoc is the result of extracting metadata without rendering HTML.
type MetaDoc struct {
	FrontMatter FrontMatter
	Body        string // raw Markdown body, not rendered
}

var md goldmark.Markdown

func init() {
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,        // tables, strikethrough, linkify, task lists
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			// Do NOT set html.WithUnsafe() — keeps safe mode equivalent to Parsedown setSafeMode(true)
			html.WithHardWraps(), // equivalent to Parsedown setBreaksEnabled(true)
		),
	)
}

// frontMatterRegex matches the YAML front matter block.
var frontMatterRegex = regexp.MustCompile(`(?s)^---\n(.*?)\n---\n(.*)$`)

// ParseMetaOnly extracts front matter and returns the raw Markdown body without rendering HTML.
// This is used during build-index generation — Goldmark is never invoked.
func ParseMetaOnly(content string) MetaDoc {
	content = sanitizeInput(content)
	fm, body := splitFrontMatter(content)
	return MetaDoc{FrontMatter: fm, Body: body}
}

// Parse extracts front matter and renders the Markdown body to HTML.
func Parse(content string) ParsedDoc {
	content = sanitizeInput(content)
	fm, body := splitFrontMatter(content)

	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		// Fallback: escape the raw body so the page at least shows something
		buf.Reset()
		buf.WriteString("<pre>")
		buf.WriteString(htmlEscape(body))
		buf.WriteString("</pre>")
	}
	return ParsedDoc{FrontMatter: fm, HTML: buf.String()}
}

// sanitizeInput ensures valid UTF-8 and normalises line endings.
func sanitizeInput(s string) string {
	if !utf8.ValidString(s) {
		// Replace invalid sequences with the Unicode replacement character
		s = strings.ToValidUTF8(s, "\uFFFD")
	}
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return s
}

// splitFrontMatter separates the YAML front matter from the Markdown body.
func splitFrontMatter(content string) (FrontMatter, string) {
	matches := frontMatterRegex.FindStringSubmatch(content)
	if matches == nil {
		return FrontMatter{}, content
	}
	fm := parseFrontMatterFields(matches[1])
	return fm, matches[2]
}

// parseFrontMatterFields parses the key: value lines of a front matter block.
// This is a port of MarkdownParser::parseFrontMatter() — intentionally simple,
// not a full YAML parser. Supports scalar values and basic inline arrays.
func parseFrontMatterFields(raw string) FrontMatter {
	fm := FrontMatter{Extra: make(map[string]string)}
	lines := strings.Split(raw, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		idx := strings.Index(line, ":")
		if idx < 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		val := strings.TrimSpace(line[idx+1:])

		switch key {
		case "title":
			fm.Title = val
		case "date":
			fm.Date = val
		case "author":
			fm.Author = val
		case "tags":
			fm.Tags = val
		case "description":
			fm.Description = val
		case "js":
			fm.JS = val
		default:
			fm.Extra[key] = val
		}
	}
	return fm
}

// htmlEscape escapes the five XML/HTML special characters.
func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, `"`, "&#34;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}
