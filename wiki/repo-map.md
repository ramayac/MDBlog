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

## Build and Run Path

- Local dev uses `make serve`.
- Tests use `make test` and depend on building index, feed, and sitemap first.
- Docker builds compile Go binaries and regenerate the index inside the image build.
- Production runs as an AWS Lambda container image.

## Repo-Specific Exclusions

- Ignore `posts/` during routine wiki ingestion and linting.
- Reason: it is a large body of user-authored content, not the primary architecture surface.
- Exception: read `posts/` only when the user explicitly asks about post content, post rendering behavior, or content-driven bugs.

## Wiki-Relevant Facts

- The root `AGENTS.md` is a primary instruction surface for the agent.
- The repo now uses `wiki/` as the persistent knowledge layer for architecture and process notes.
- The wiki should summarize code and workflows, not duplicate or rewrite user-authored posts.
