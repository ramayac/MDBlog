package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/render"
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

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg := config.MustLoad("config.toml")

	switch os.Args[1] {
	case "serve":
		runServe(cfg)
	case "build-index":
		runBuildIndex(cfg)
	case "render":
		runRender(cfg, os.Args[2:])
	case "version":
		fmt.Printf("mdblog %s (%s) built %s\n", version, commit, date)
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %q\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: mdblog <subcommand> [args]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("  serve         Start HTTP server (default :8080, set PORT env to override)")
	fmt.Println("  build-index   Generate posts/posts.index.json")
	fmt.Println("  render        Render a post to a standalone HTML file")
	fmt.Println("  version       Print version information")
}

func runServe(cfg *config.Config) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	b := blog.New(cfg)
	h := server.New(cfg, b)

	fmt.Printf("Listening on http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, h); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func runBuildIndex(cfg *config.Config) {
	if err := buildindex.Build(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "build-index: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Post index built successfully.")
}

func runRender(cfg *config.Config, args []string) {
	if err := render.Run(cfg, args); err != nil {
		fmt.Fprintf(os.Stderr, "render: %v\n", err)
		os.Exit(1)
	}
}
