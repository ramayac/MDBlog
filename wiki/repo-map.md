# MDBlog Repo Map

## Purpose

MDBlog is a flat-file blog engine written in Go 1.24. It serves Markdown posts and standalone pages, generates a metadata index for listing and search, and can run locally or as an AWS Lambda container image.

## High-Signal Areas

- `cmd/mdblog/` holds the CLI entry point.
- `cmd/lambda/` and `cmd/lambda-embed/` hold Lambda entry points.
- `internal/blog/` holds post, page, category, menu, and search domain logic.
- `internal/server/` holds HTTP routing, templates, gzip, CSP, and SEO serving.
- `internal/buildindex/`, `internal/buildfeed/`, and `internal/buildsitemap/` generate derived artifacts.
- `templates/` and `assets/` define the UI surface.
- `config.toml` is the runtime configuration source of truth.

## Generated Artifacts

- `posts/posts.index.json` is the build-time metadata index.
- `feed.xml` is the build-time RSS feed.
- `sitemap.xml` and `robots.txt` are build-time SEO outputs.

## Feed and SEO Artifacts

- `internal/buildfeed/` builds `feed.xml` from `posts/posts.index.json`, not by re-rendering Markdown posts.
- The RSS feed is enabled through `[feed]` config, uses `feed.base_url` for absolute links, caps output to `feed.max_items`, and writes item descriptions from the prebuilt post excerpts.
- Runtime feed serving has two paths: `/feed.xml` serves the prebuilt XML file directly, while `/feed` renders a human-readable feed page from `blog.GetFeedPosts()`.
- Unlike sitemap and robots, `/feed.xml` does not have a dynamic fallback when the built file is missing; it returns a not-found style error instead.
- `internal/buildsitemap/` builds both `sitemap.xml` and `robots.txt` from the same prebuilt post index.
- The sitemap includes the home page, one URL per discovered category slug, and one URL per indexed post. Category and post URLs use query-string routes, and post entries include `lastmod` when a post date exists.
- Sitemap `changefreq` and `priority` values come from the `[sitemap]` config block.
- The generated `robots.txt` is intentionally simple: it allows all crawlers and advertises the sitemap URL.
- In production, `server/seo.go` serves the prebuilt `sitemap.xml` and `robots.txt` files from disk. For local development, if those files are absent, the handler can generate fallback responses dynamically so `make serve` still works without running the build step first.

## Template Rendering

- MDBlog uses Go `html/template` for all page rendering.
- `internal/server/handler.go` loads every `templates/*.html` file at startup from `TemplateFS`, which defaults to the on-disk `templates/` directory.
- The embed build can override `TemplateFS` and `AssetsFS` so templates and assets come from the repository-root `go:embed` bundle in `embed.go`.
- `embed.go` lives at the repository root because `go:embed` cannot reach sibling directories through `..`; placing it there lets one package embed both `templates/` and `assets/` directly.
- `embed.go` exposes `EmbeddedTemplates()` and `EmbeddedAssets()`, each returning an `fs.FS` rooted to the corresponding subtree via `fs.Sub`.
- `cmd/lambda-embed/main.go` assigns those embedded filesystems to `server.TemplateFS` and `server.AssetsFS` before building the HTTP handler, so the rest of the server code can keep using the same filesystem abstraction without special cases.
- The purpose of the embed variant is deployment simplicity: the compiled Lambda binary already contains the UI templates and static assets, so the embed Docker image does not need `templates/` or `assets/` copied alongside it at runtime.
- Posts, pages, config, and generated feed and sitemap files still stay on disk in the embed image; `embed.go` is only for templates and assets.
- Each route prepares a shared `templateData` struct that carries global page state such as `Config`, menu links, footer HTML, version info, canonical URL, JSON-LD, and page-specific fields like `Post`, `Posts`, `Page`, `Categories`, and pagination.
- Route-specific templates such as `home.html`, `category.html`, `search.html`, `post.html`, `page.html`, `feed.html`, and `404.html` render the inner body first.
- The handler then stores that inner HTML in `templateData.Content` and executes `layout.html` as the outer shell, so navigation, metadata, footer, theme toggles, and optional per-post JS are centralized in one layout.
- Shared template helpers are registered in the server, including `safeHTML`, `formatDate`, `defaultStr`, and `postPreviewData`.
- `_post_preview.html` is a reusable partial defined as the `post_preview` template and is invoked by listing pages like category and search.
- Post and page bodies are already rendered to HTML by the blog/markdown layer before templates receive them, so templates mostly assemble page chrome, metadata, and links around trusted rendered content.

## Footer

- The footer is rendered once in `templates/layout.html`, so every HTML page shares the same footer shell.
- Each route prepares `FooterHTML` by rendering `footer_content` from `config.toml` through `blog.ParseMarkdown()`, which means the footer text supports Markdown formatting.
- The footer template prints that rendered Markdown first when `FooterHTML` is non-empty.
- The footer can also show build metadata from `blog.GetVersionInfo()`: version, commit, and build date.
- When `show_render_time = true`, the handler measures request render time and appends the value in milliseconds to the footer's version line.
- If render timing is enabled but no build commit is available, the footer still shows the render time by itself.
- Non-HTML artifact endpoints such as `/feed.xml` do not use the footer because they do not render through `layout.html`.

## CSS Themes

- Theme selection is global and configuration-driven through `css_theme` in `config.toml`.
- The layout template emits a single stylesheet link using `.Config.CSSTheme`, and the server adds a `?v=` query parameter based on the selected file's modification time so CSS changes bust caches without changing filenames.
- Static CSS files are served through the `/assets/*` route from `AssetsFS`, which points at the on-disk `assets/` directory in the default build and can be swapped to embedded assets in the embed build.
- The current stylesheet set is `assets/css/base.style.css`, `assets/css/default.style.css`, and `assets/css/anthropic.style.css`.
- `base.style.css` is the shared structural layer: reset rules, CSS variables, layout primitives, typography, menus, cards, and other common component styles.
- `anthropic.style.css` is a theme built on top of that shared layer via `@import "base.style.css"`; it overrides design tokens and adds theme-specific component styling.
- `default.style.css` is a standalone theme file with its own full rule set rather than importing `base.style.css`.
- Theme switching between light and dark modes is client-side: `layout.html` stores the selected mode in `localStorage` under `theme-mode`, sets `data-theme` on the root element, and each theme file defines light variables, explicit `:root[data-theme="dark"]` overrides, plus a `prefers-color-scheme: dark` fallback when no explicit choice has been stored.
- This means MDBlog chooses one CSS file at runtime, and that CSS file internally handles both the visual design and the light/dark variant behavior.

## Compression and Security

- HTML and feed responses can be gzip-compressed when the request advertises `Accept-Encoding: gzip`, but the server intentionally skips that step on AWS Lambda because API Gateway or CloudFront is expected to handle compression there.
- Static assets are served through a narrow `/assets/*` route that rejects `..` path traversal sequences and null bytes before opening files from `AssetsFS`.
- A configurable Content Security Policy is injected per request from the `[csp]` section of `config.toml`.
- Markdown rendering is kept in safe mode: the Goldmark configuration does not enable `html.WithUnsafe()`, so raw HTML passthrough is not enabled in post or page content.
- Page templating uses Go `html/template`, so ordinary template values are auto-escaped by default; explicit raw HTML output is limited to trusted server-side content paths such as already-rendered Markdown and JSON-LD blocks.
- Post and standalone page loading both include path-traversal guards before reading from `posts/` or `pages/`.
- The production runtime is also hardened by the deployment model: the app ships as a minimal `FROM scratch` container image with no database, no mutable application state, and prebuilt content artifacts baked into the image.

## Build and Run Path

- Local dev uses `make serve`.
- Local container preview uses `make docker-run`, which delegates to `docker compose up --build` and reads `docker-compose.yml` from the repo root.
- `docker-compose.yml` defines a single `blog` service that builds from the standard root `Dockerfile`, tags the resulting image as `mdblog:latest`, and then overrides the container command to `/mdblog serve` so local Docker runs the plain HTTP server instead of the Lambda entry point.
- The compose service sets `PORT=8080`, publishes host `8080` to container `8080`, and uses `restart: unless-stopped` for local restarts.
- The compose file also applies the same local hardening intent as production by setting `read_only: true`, `security_opt: [no-new-privileges:true]`, and `cap_drop: [ALL]`, and it declares a `64M` memory limit under `deploy.resources.limits`.
- Tests use `make test` and depend on building index, feed, and sitemap first.
- Docker builds compile Go binaries and regenerate the index inside the image build.
- `make docker-run-release` reuses that same compose file but calls `docker compose up --no-build`, which is intended for a previously pulled `mdblog:latest` image instead of rebuilding from local sources.
- Production runs as an AWS Lambda container image.

## Make Targets

- Core development targets: `help`, `serve`, `build`, `build-embed`, `build-index`, `build-feed`, `build-sitemap`, `lint`, `lint-config`, `test`, `render`, and `new-post`.
- Wiki maintenance targets: `wiki-list`, `wiki-headings`, `wiki-log-tail`, `wiki-search`, `wiki-changed`, `wiki-ingest-candidates`, `wiki-lint`, and `wiki-refresh`.
- Docker targets: `docker-build`, `docker-build-embed`, `docker-run`, `docker-run-release`, `docker-stop`, `docker-push`, and `docker-pull`.
- `help` is the default goal and prints the annotated target list from the Makefile.
- `render` is argument-driven and supports forms like `make render random`, `make render [category] random`, and `make render filename.md`.
- `new-post` scaffolds a Markdown post using `TITLE` and optional `CATEGORY` and `TAGS` variables.

## Testing Model

- MDBlog uses Go test files under `internal/` for package-level automated testing.
- The repo currently includes tests for config loading, Markdown parsing, blog domain logic, index generation, feed generation, sitemap generation, handler behavior, and SEO behavior.
- The main test entrypoint is `make test`, which first regenerates the build-time index, feed, and sitemap artifacts, then runs `go test ./...`.
- The tests are package-local white-box style tests rather than a separate external integration test suite.

## AWS Runtime Model

- MDBlog is deployed as a Docker container image to AWS Lambda behind API Gateway.
- The Lambda entry point is `cmd/lambda`, which builds the normal blog handler and passes it to `algnhsa`, letting API Gateway events be served through the same standard `net/http` handler used elsewhere.
- The standard production image is a multi-stage `FROM scratch` build that copies in the Lambda binary plus the content and runtime files the handler needs: `posts/`, `pages/`, `assets/`, `templates/`, `config.toml`, `feed.xml`, `sitemap.xml`, and `robots.txt`.
- The embed image variant keeps the same runtime model but bakes `templates/` and `assets/` into the binary, so only content and config files remain on disk in the final image.
- The image is intended to run read-only, with no mutable application state and no runtime build step.
- In the standard `Dockerfile`, the build stage produces two binaries, `/out/lambda` and `/out/mdblog`, then uses `/out/mdblog` to generate `posts/posts.index.json`, `feed.xml`, `sitemap.xml`, and `robots.txt` before copying the runtime files into the final image.
- The standard final image contains `/lambda`, `/mdblog`, `/posts/`, `/pages/`, `/assets/`, `/templates/`, `/config.toml`, `/feed.xml`, `/sitemap.xml`, `/robots.txt`, and CA certificates.
- In `Dockerfile.embed`, the build stage produces `/out/lambda-embed` and `/out/mdblog`, then generates the same derived content artifacts.
- The embed final image contains `/lambda`, `/posts/`, `/pages/`, `/config.toml`, `/feed.xml`, `/sitemap.xml`, `/robots.txt`, and CA certificates; `templates/` and `assets/` are omitted because they are embedded into the binary.

## Publication Flow

- The production publish path is GitHub Actions based rather than a runtime CMS upload flow.
- A push to `master` that changes `posts/**/*.md` triggers `.github/workflows/ghcr-release.yml`.
- That workflow builds the Docker image, computes version metadata from git, and pushes the resulting image to GHCR with tags including `latest`.
- The Docker build itself regenerates `posts/posts.index.json`, `feed.xml`, `sitemap.xml`, and `robots.txt`, so Markdown post changes become part of the published image through that build stage.
- A second workflow, `.github/workflows/aws-deploy.yml`, listens for successful completion of the GHCR workflow, pulls the `latest` image from GHCR, retags and pushes it to Amazon ECR, and then runs `aws lambda update-function-code` to point the Lambda function at the new container image.
- This means post Markdown files are published by being committed and pushed to `master`, not by being uploaded directly to the running server.
- The automatic push trigger is scoped to `posts/**/*.md`; other Markdown locations such as `pages/` are not covered by that push-path filter and therefore do not use the same automatic publish trigger unless deployment is initiated through another path such as a release.

## Content and Persistence Model

- MDBlog has no database because its source of truth is the repository content itself: posts are Markdown files under `posts/`, standalone pages are Markdown files under `pages/`, and runtime settings live in `config.toml`.
- Listing and search do not query a database; they read the pre-built `posts/posts.index.json` metadata file generated during the image build.
- Single post and page requests still read the corresponding Markdown file from disk inside the container and render it on demand.
- RSS and SEO files are also pre-generated during the image build, then served as static runtime artifacts.
- If the post index is missing or invalid, the blog can fall back to scanning Markdown files directly, which preserves correctness without introducing a database dependency.
- This works well on Lambda because content changes only when a new image is built and deployed, so the application does not need runtime writes or external persistence.

## Search Flow

- Search is handled on the main `/` route using the `q` query parameter or the presence of `search=1`.
- `search=1` opens the standalone search page with an empty query; `q=<term>` executes the search and renders results with `templates/search.html`.
- Runtime search uses `blog.SearchPosts()`, which depends on the prebuilt `posts/posts.index.json` metadata file.
- Search is a simple case-insensitive substring match against the concatenation of each indexed post's `title`, `excerpt`, and `tags` fields.
- Results are sorted newest-first, paginated with `posts_per_page`, and rendered through the same `post_preview` partial used by other listing pages.
- Search does not render full Markdown bodies or scan the filesystem for body text at request time.
- Unlike listing pages, search has no live filesystem fallback when the index is missing; it returns an empty result set with pagination scaffolding instead.
- The excerpts searched at runtime come from the build-index step: the index builder prefers the post `description` field, and otherwise derives an excerpt from the raw Markdown body.

## Post Structure

- A post is a Markdown file stored under `posts/`, usually inside a category subfolder such as `posts/personal/slug.md`.
- The expected filename convention is `YYYY-MM-DD-slug-with-hyphens.md`; `make new-post` scaffolds files in that shape automatically.
- Post files may begin with a simple YAML-style front matter block delimited by `---` lines, followed by the Markdown body.
- The parser recognizes these front matter keys: `title`, `date`, `author`, `tags`, `description`, and `js`. Unknown keys are preserved in an `Extra` map but are not part of the main rendering contract.
- `js` refers to an optional JavaScript file loaded from `assets/js/` for that post.
- The Markdown body is rendered with Goldmark using GFM features, footnotes, auto heading IDs, and safe-mode HTML handling.
- In practice, `title` and `date` are optional because runtime parsing falls back to the filename-derived title and the file modification time when those fields are missing.
- `tags` is stored as a raw comma-separated string in front matter and split later for template rendering.
- `description` is used for metadata and excerpts when present.
- A typical post shape is:

```yaml
---
title: My Post Title
date: 2024-01-15
author: Your Name
tags: tag1, tag2
description: Optional meta description
js: optional-script.js
---

Markdown body here.
```

## Repo-Specific Exclusions

- Ignore `posts/` during routine wiki ingestion and linting.
- Reason: it is a large body of user-authored content, not the primary architecture surface.
- Exception: read `posts/` only when the user explicitly asks about post content, post rendering behavior, or content-driven bugs.

## Wiki-Relevant Facts

- The root `AGENTS.md` is a primary instruction surface for the agent.
- The repo now uses `wiki/` as the persistent knowledge layer for architecture and process notes.
- The wiki should summarize code and workflows, not duplicate or rewrite user-authored posts.
