# Markdown Blog

A simple Markdown-based blog engine written in **Go**, deployed as a Docker container on AWS Lambda. The idea behind it is this workflow: Write → Commit → Push → Deploy.

This used to be a generic project, but I ended up keeping my posts and configuration here.
Eventually, I will work on a general release :)

Author: [@ramayac](https://x.com/ramayac).

## Features

- Write posts in Markdown with YAML-style front matter
- Built-in pagination powered by a **build-time metadata index** (handles 300+ posts without Lambda timeouts)
- Standalone search page using the pre-built metadata index
- Fallback post routing for clean URLs (resolves posts missing a category path)
- GitHub Flavoured Markdown + footnotes via [Goldmark](https://github.com/yuin/goldmark)
- Custom JavaScript support per post
- Responsive design with Dark/Light theme based on OS preference
- Gzip compression (when supported by the client)
- Dynamic navigation menu driven by `config.toml`
- Landing page with category cards (no full post scan on homepage)
- Statically linked Go binary — no runtime dependencies
- Minimal `FROM scratch` Docker image, read-only filesystem, all capabilities dropped
- AWS Lambda ready via [algnhsa](https://github.com/akrylysov/algnhsa)

## Quick Start

MDBlog is deployed as a Docker container image on AWS Lambda behind API Gateway.

1. Clone the repo and configure `config.toml` (blog name, author, categories)
2. Add `.md` posts to `posts/<category>/` directories
3. Build and push the Docker image: `make docker-build && make docker-push`
4. Deploy the image to AWS Lambda as a container image function behind API Gateway

For local development, see [Running Locally](#running-locally) below.

## Running Locally

Requires **Go 1.24+** and `make`. No other runtime dependencies.

```bash
make build-index     # Generate post metadata index (posts/posts.index.json)
make serve           # Start HTTP dev server at http://localhost:8080
make lint            # Run go vet on all packages
make test            # Build index then run the Go test suite
make render random   # Render a random post to a standalone HTML file
```

> **Tip:** Run `make build-index` whenever you add or edit posts locally so that paginated
> listings reflect your changes immediately. Without the index the blog falls back to a
> full filesystem scan — correct but slower for large categories.

## Deployment (AWS Lambda)

The production image uses a **multi-stage Docker build**: a `golang:1.24` stage compiles the Go binary and generates the post index; the final stage copies only the binary, posts, templates, assets and config into a minimal `FROM scratch` image.

```bash
make docker-build                        # Build production image (FROM scratch, Lambda-ready)
make docker-push REGISTRY=ghcr.io/...   # Tag and push to a container registry
make docker-pull TAG=1.2.3              # Pull a release image and tag as latest
```

After pushing the image, update the Lambda function to use the new image URI.

### Embed Variant

`Dockerfile.embed` builds `cmd/lambda-embed`, which has `templates/` and `assets/` baked into the binary via `go:embed`. The resulting image only needs the binary, `posts/`, and `config.toml`.

```bash
make docker-build-embed   # Build the embed-variant image
```

### Continuous Deployment (CI/CD)

MDBlog includes GitHub Actions for automated deployment. Pushing any `.md` file in `posts/` to `master` triggers a Docker build. Once pushed to GHCR, a secondary workflow propagates the image to Amazon ECR and updates the Lambda function.

### Caching on Lambda

There is no file-based cache. The container filesystem is read-only; the pre-built JSON index is baked into the image.

**Recommended caching strategy:** place **CloudFront** in front of the Lambda function. Since posts change only on redeploy, a CloudFront TTL of hours or days is safe. Invalidate the distribution after each `make docker-push`.

## Creating a New Post

```bash
make new-post TITLE="My Post Title" CATEGORY=my-category TAGS="tag1, tag2"
```

Creates a pre-filled `.md` file at `posts/<category>/YYYY-MM-DD-my-post-title.md`.
Author is read automatically from `config.toml`. `CATEGORY` and `TAGS` are optional.

## Writing Posts

Create a `.md` file in a category subfolder under `posts/` with front matter:

```yaml
---
title: My Post
date: 2024-01-15
author: Your Name
tags: tag1, tag2
description: Optional meta description
js: optional-script.js   # loaded from assets/js/
---

Your markdown content here (GFM + footnotes).
```

## Navigation Menu

The nav bar is generated automatically from `config.toml`.

```toml
# Static custom links (always shown, in order)
[[menu_links]]
label = "Home"
url   = "index.php"

# Per-category: set menu = true to include it in the nav
[categories.srbyte]
blog_name = "Sr. Byte"
menu      = true
```

## Landing Page and Search

The home page (`/` with no query params) shows a static landing page with category cards.
To add an optional intro blurb above the cards, create `posts/index.md`.

Browsing posts: `/?category=slug`
Searching posts: `/?q=keyword` (requires the post metadata index)

## Categories

Register categories in `config.toml`:

```toml
[categories.my-category]
blog_name      = "Display Name"
header_content = "Subtitle."
folder         = "my-category"   # subfolder under posts/
index          = true            # include in legacy aggregated index
menu           = true            # show in nav bar
```

Then add `.md` files to `posts/my-category/`.

## Post Metadata Index

Listing and pagination pages are powered by a **pre-built metadata index** (`posts/posts.index.json`) that avoids scanning and parsing all Markdown files on every request.

### How it works

1. `make build-index` (`internal/buildindex.Build()`) scans all posts, extracts front-matter metadata, and writes `posts/posts.index.json`. **Goldmark is never called** — full post bodies are not rendered during this step.
2. `make docker-build` runs `make build-index` automatically inside the Docker build stage, so the index is baked into the image.
3. At request time, `blog.GetPosts()` reads the index for filtering and pagination, and `blog.SearchPosts()` uses it for full-text search — no `.md` files are opened.
4. Individual post pages still parse the full Markdown body, but only for the single requested post. The index is also used as a fallback to resolve a post's parent category when it is missing from the URL.

### Fallback

If `posts/posts.index.json` is absent, the blog falls back to a live filesystem scan with a performance warning logged. Search and slug-only URL resolution will not work without the index.

### Keeping the index fresh

```bash
make build-index   # regenerate posts/posts.index.json
```

## Configuration

Edit `config.toml` to customize all settings. Key fields:

| Key | Purpose |
|-----|---------|
| `blog_name` | Site title |
| `author_name` | Default author for new posts |
| `header_content` | Landing page subtitle |
| `footer_content` | Footer text (Markdown supported) |
| `posts_per_page` | Pagination size |
| `excerpt_length` | Max characters in post excerpt |
| `show_uncategorized` | Show root-level posts in listings |
| `post_index_file` | Path to pre-built metadata index |
| `css_theme` | Active CSS theme path |
| `[[menu_links]]` | Static nav links |
| `[categories.<slug>]` | Category definitions |
| `[csp]` | Content Security Policy (`enabled`, `header`) |
| `[labels]` | All user-visible UI strings |

## Architecture

```
cmd/
  mdblog/           # CLI: serve | build-index | render | version
  lambda/           # AWS Lambda entry point (disk-based templates/assets)
  lambda-embed/     # AWS Lambda entry point (templates+assets in binary)
internal/
  blog/             # Core domain: posts, pagination, menu, search
  buildindex/       # Build-time index generator
  config/           # TOML config loader
  markdown/         # Front matter parser + Goldmark renderer
  render/           # CLI render subcommand
  server/           # net/http handler: routing, templates, gzip, CSP
templates/          # Go html/template files (*.html)
assets/css/         # CSS themes
assets/js/          # Per-post JavaScript files
posts/              # All Markdown content
embed.go            # go:embed declarations
config.toml         # Runtime configuration
Dockerfile          # Standard Lambda image (FROM scratch)
Dockerfile.embed    # Embed variant (binary + posts + config only)
```

## Requirements

**Production (Lambda):**
- Docker (to build the image)
- AWS account with Lambda + API Gateway (+ optionally CloudFront)
- Container registry (e.g. ghcr.io or ECR)

**Local development:**
- Go 1.24+
- `make`

No database. No PHP. No Node.js.

## License

MIT License
