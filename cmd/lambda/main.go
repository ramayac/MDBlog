package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/server"
)

// Injected via -ldflags at build time:
//
//	-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	blog.BuildVersion = version
	blog.BuildCommit = commit
	blog.BuildDate = date

	cfg := config.MustLoad("config.toml")
	b := blog.New(cfg)
	h := server.New(cfg, b)

	algnhsa.ListenAndServe(h, nil)
}
