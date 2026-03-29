package render

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/ramayac/mdblog/internal/blog"
	"github.com/ramayac/mdblog/internal/buildindex"
	"github.com/ramayac/mdblog/internal/config"
	"github.com/ramayac/mdblog/internal/server"
)

// Run renders a single post to an HTML file in the render/ directory.
// Mirrors scripts/render.php.
//
// Supported argument patterns (args = os.Args[1:] after "render"):
//   - (empty)               → print usage
//   - random                → pick a random post from the full index
//   - <category> random     → pick a random post from a given category
//   - <filename.md>         → render the named post
func Run(cfg *config.Config, args []string) error {
	// Ensure the post index exists; build it if not.
	if _, err := os.Stat(cfg.PostIndexFile); err != nil {
		fmt.Println("Post index not found. Building it...")
		if err := buildindex.Build(cfg); err != nil {
			return fmt.Errorf("render: build index: %w", err)
		}
	}

	index, err := loadIndex(cfg.PostIndexFile)
	if err != nil {
		return fmt.Errorf("render: load index: %w", err)
	}
	if len(index) == 0 {
		return fmt.Errorf("render: no posts in index")
	}

	if len(args) == 0 {
		printUsage()
		return nil
	}

	// Resolve the target post
	var target *buildindex.IndexPost
	switch {
	case len(args) == 1 && args[0] == "random":
		p := index[rand.Intn(len(index))]
		target = &p

	case len(args) == 2 && args[1] == "random":
		catSlug := args[0]
		var catPosts []buildindex.IndexPost
		for _, ip := range index {
			if ip.CategorySlug == catSlug {
				catPosts = append(catPosts, ip)
			}
		}
		if len(catPosts) == 0 {
			return fmt.Errorf("render: no posts found for category %q", catSlug)
		}
		p := catPosts[rand.Intn(len(catPosts))]
		target = &p

	default:
		search := args[0]
		searchSlug := strings.TrimSuffix(search, ".md")
		for i, ip := range index {
			if ip.Filename == search || ip.Slug == searchSlug || strings.Contains(ip.Filename, search) {
				target = &index[i]
				break
			}
		}
	}

	if target == nil {
		return fmt.Errorf("render: could not find a matching post for %q", strings.Join(args, " "))
	}

	fmt.Printf("Target Post details:\n")
	fmt.Printf("  Title:     %s\n", target.Title)
	fmt.Printf("  Date:      %s\n", target.Date)
	fmt.Printf("  Slug:      %s\n", target.Slug)
	fmt.Printf("  Category:  %s\n", orUncategorized(target.CategorySlug))
	fmt.Printf("  File:      %s\n\n", target.SourcePath)

	b := blog.New(cfg)

	// Build the HTML by rendering post template through the HTTP handler.
	// We reuse the server's template machinery via a fake HTTP round-trip.
	h := server.New(cfg, b)

	url := "/post.php?slug=" + target.Slug
	if target.CategorySlug != "" {
		url += "&category=" + target.CategorySlug
	}

	req := httptest.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	resp := rr.Result()
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("render: read response body: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("render: post not found (slug=%s category=%s)", target.Slug, target.CategorySlug)
	}

	html := string(htmlBytes)
	// Inject a base tag so local assets resolve correctly
	html = strings.Replace(html, "<head>", "<head>\n    <base href=\"../\">", 1)

	outDir := "render"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("render: create output dir: %w", err)
	}
	outFile := filepath.Join(outDir, target.Slug+".html")
	if err := os.WriteFile(outFile, []byte(html), 0644); err != nil {
		return fmt.Errorf("render: write output file: %w", err)
	}

	fmt.Printf("SUCCESS! Wrote rendered HTML file to: %s\n", outFile)
	return nil
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  mdblog render random")
	fmt.Println("  mdblog render [category] random")
	fmt.Println("  mdblog render [filename.md]")
}

func orUncategorized(s string) string {
	if s == "" {
		return "(uncategorized)"
	}
	return s
}

func loadIndex(path string) ([]buildindex.IndexPost, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var posts []buildindex.IndexPost
	if err := json.Unmarshal(data, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}
