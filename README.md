# Markdown Blog

A simple Markdown-based blog engine written in **Go**, deployed as a Docker container on AWS Lambda. The idea behind it is this workflow: Write → Commit → Push → Deploy.

This used to be a generic project, but I ended up keeping my posts and configuration here.
Eventually, I will work on a general release :)

Author: [@ramayac](https://x.com/ramayac).

## Features

- Write posts in Markdown with YAML-style front matter
- **Standalone pages** (`pages/` directory, `/page?slug=`) — About, Projects, or any single page without dates or categories
- Built-in pagination powered by a **build-time metadata index** (handles 300+ posts without Lambda timeouts)
- Standalone search page using the pre-built metadata index
- **RSS 2.0 feed** (`/feed.xml`) generated at build time; human-readable feed page at `/feed`
- Fallback post routing for clean URLs (resolves posts missing a category path)
- GitHub Flavoured Markdown + footnotes via [Goldmark](https://github.com/yuin/goldmark)
- Custom JavaScript support per post
- Responsive design with Dark/Light theme based on OS preference
- Gzip compression (when supported by the client)
- Dynamic navigation menu driven by `config.toml`, with category **dropdown** grouping
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
make build-feed      # Generate RSS feed (feed.xml — requires build-index first)
make build-sitemap   # Generate sitemap.xml and robots.txt (requires build-index first)
make serve           # Start HTTP dev server at http://localhost:8080
make lint            # Run go vet on all packages
make test            # Build index + feed + sitemap, then run the Go test suite
make render random   # Render a random post to a standalone HTML file
make wiki-refresh    # Show wiki files, recent log, diff-driven inputs, and lint results
```

> **Tip:** Run `make build-index`, `make build-feed`, and `make build-sitemap` whenever you add or edit posts
> locally so that paginated listings, the RSS feed, and the sitemap reflect your changes immediately.

## Repository Wiki

This repo now includes a persistent `wiki/` directory for architecture notes, operating procedures, and durable answers that should survive beyond a single chat session.

- Start with `wiki/index.md`.
- Check recent history in `wiki/log.md`.
- Use `wiki/operations/ingest.md`, `wiki/operations/query.md`, and `wiki/operations/lint.md` for the repeatable workflow.
- Treat `posts/` as excluded from routine wiki maintenance because it is user-authored content rather than the primary architecture surface.

The wiki is intentionally shell-friendly. Typical inspection commands:

```bash
find wiki -maxdepth 2 -type f | sort
grep -R "^#\|^##" wiki
grep "^## \[" wiki/log.md | tail -10
```

Equivalent `make` entrypoints are available for repeatable use:

```bash
make wiki-list
make wiki-headings
make wiki-log-tail WIKI_LOG_N=5
make wiki-search WIKI_Q=category
make wiki-changed WIKI_DIFF=master...HEAD
make wiki-ingest-candidates WIKI_DIFF=master...HEAD
make wiki-lint
make wiki-refresh
```

All wiki helper targets are shell-first and call plain `sh` scripts under `scripts/`.

There are also on-demand chat prompts under `.github/prompts/` for wiki maintenance:

```text
/wiki-refresh
/wiki-ingest
/wiki-query
```

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

The nav bar is driven entirely by `config.toml` and renders in three layers:

1. **`[[menu_links]]`** — static links, always shown first (e.g. Home, About).
2. **`[[menu.pinned]]`** — category links shown directly in the bar (no dropdown).
3. **`[menu.categories]`** — category links grouped into a single **dropdown**, labelled by `menu.categories.label`.

```toml
# Static links
[[menu_links]]
label = "Home"
url   = "/"

[[menu_links]]
label = "About"
url   = "/page?slug=about"

# Pinned categories — shown inline in the nav bar
[[menu.pinned]]
category = "guides"
order    = 2

# Dropdown section — all items appear under a single "Writings" button
[menu.categories]
label = "Writings"

[[menu.categories.item]]
category = "personal"
order    = 1

[[menu.categories.item]]
category = "srbyte"
order    = 2
```

Result: `HOME | ABOUT | GUIDES | WRITINGS ▾ | 🔍 | 🌓`  
Hovering **WRITINGS ▾** opens the dropdown with all `menu.categories.item` links.

## Standalone Pages

Pages are Markdown files that live in `pages/` (configurable via `pages_dir`). Unlike posts they have no date, category, author, or tags — just a title and rendered content.

```bash
# Create a page
echo '---\ntitle: About\n---\n# About Me\n...' > pages/about.md
```

Add a nav link in `config.toml`:

```toml
[[menu_links]]
label = "About"
url   = "/page?slug=about"
```

Routes: `/page?slug=<slug>` renders `pages/<slug>.md` using `templates/page.html`.

## Landing Page and Search

The home page (`/` with no query params) shows a static landing page with category cards.
To add an optional intro blurb above the cards, create `posts/index.md`.

Browsing posts: `/?category=slug`
Searching posts: `/?q=keyword` (requires the post metadata index)

## RSS Feed

MDBlog generates a valid RSS 2.0 feed at build time — no XML is generated on every request.

- **`/feed.xml`** — Machine-readable RSS 2.0 feed served on request from the pre-built file
- **`/feed`** — Human-readable feed page listing all recent posts (date, category, title)

Configure the feed in `config.toml`:

```toml
[feed]
enabled     = true
title       = "Your Blog Name"
description = "Blog description."
base_url    = "https://your-domain.com"   # required — used for absolute URLs in the XML
max_items   = 50
output_file = "feed.xml"
```

`base_url` is **required** when `feed.enabled = true`. `make build-index` must run before `make build-feed`.

To add a "Feed" link to the nav bar, add to `config.toml`:

```toml
[[menu_links]]
label = "Feed"
url   = "/feed"
```

### Generating the feed

```bash
make build-index   # must run first
make build-feed    # writes feed.xml
```

`make docker-build` runs both steps automatically inside the build stage.

## SEO — Sitemap and robots.txt

MDBlog generates `sitemap.xml` and `robots.txt` at build time. In production, these pre-built files are served directly from disk; during local development (`make serve`), the server generates them on-the-fly so no extra build step is required.

- **`/sitemap.xml`** — Valid XML sitemap listing the homepage, all category pages, and all individual posts
- **`/robots.txt`** — Allows all crawlers and points to the sitemap

Configure in `config.toml`:

```toml
[sitemap]
enabled              = true
output_file          = "sitemap.xml"    # path to write
robots_file          = "robots.txt"     # path to write
changefreq_home      = "weekly"         # valid: always, hourly, daily, weekly, monthly, yearly, never
changefreq_category  = "weekly"
changefreq_post      = "monthly"
priority_home        = "1.0"            # 0.0–1.0
priority_category    = "0.8"
priority_post        = "0.6"
```

`sitemap.enabled = true` requires `feed.base_url` to be set (used for absolute URLs).

### Generating sitemap and robots.txt

```bash
make build-index    # must run first
make build-sitemap  # writes sitemap.xml and robots.txt
```

`make docker-build` runs all three steps (`build-index`, `build-feed`, `build-sitemap`) automatically.

**After deploying**, submit `https://your-domain.com/sitemap.xml` in [Google Search Console](https://search.google.com/search-console) to accelerate indexing.

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
| `pages_dir` | Standalone pages directory (default `"pages"`) |
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
  blog/             # Core domain: posts, pages, pagination, menu, search
  buildindex/       # Build-time index generator
  config/           # TOML config loader
  markdown/         # Front matter parser + Goldmark renderer
  render/           # CLI render subcommand
  server/           # net/http handler: routing, templates, gzip, CSP
templates/          # Go html/template files (*.html)
assets/css/         # CSS themes
assets/js/          # Per-post JavaScript files
posts/              # All blog post content (organized in category subfolders)
pages/              # Standalone pages (about.md, etc.) — no category, no date
embed.go            # go:embed declarations
config.toml         # Runtime configuration
Dockerfile          # Standard Lambda image (FROM scratch)
Dockerfile.embed    # Embed variant (binary + posts + pages + config only)
```

## Requirements

**Production (Lambda):**
- Docker (to build the image)
- AWS account with Lambda + API Gateway (+ optionally CloudFront)
- Container registry (e.g. ghcr.io or ECR)

**Local development:**
- Go 1.24+
- `make`

No database. All data is read from the `posts/` directory and the pre-built JSON index.

## License

MIT License
