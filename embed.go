// Package mdblog exposes static assets embedded at compile time.
// This file is intentionally at the repository root so that the //go:embed
// directives can reference the sibling templates/ and assets/ directories
// (go:embed paths cannot traverse ".." out of the package directory).
package mdblog

import (
	"embed"
	"io/fs"
)

//go:embed all:templates all:assets
var embeddedFiles embed.FS

// EmbeddedTemplates returns an fs.FS rooted at the embedded templates/ tree.
func EmbeddedTemplates() fs.FS {
	sub, err := fs.Sub(embeddedFiles, "templates")
	if err != nil {
		panic("mdblog: sub templates: " + err.Error())
	}
	return sub
}

// EmbeddedAssets returns an fs.FS rooted at the embedded assets/ tree.
func EmbeddedAssets() fs.FS {
	sub, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		panic("mdblog: sub assets: " + err.Error())
	}
	return sub
}
