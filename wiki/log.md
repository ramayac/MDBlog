# Wiki Log

## [2026-04-13] bootstrap | established repo-local wiki scaffold

- Added a dedicated `wiki/` directory with a stable index, log, schema, operations, and rollout plan.
- Defined MDBlog-specific exclusions so routine wiki maintenance ignores `posts/` unless explicitly requested.
- Added repo instructions so the agent reads the wiki before broad analysis and files durable findings back into it.

## [2026-04-13] ingest | added shell-first wiki helper targets

- Added `wiki-*` Make targets for listing, searching, diff-driven ingest, linting, and a combined refresh snapshot.
- Added plain `sh` helper scripts under `scripts/` so the workflow remains unix-friendly and portable.
- Documented the new manual entrypoints in the repo docs and wiki operations pages.

## [2026-04-13] ingest | added wiki slash prompts

- Added workspace prompt files for `wiki-refresh`, `wiki-ingest`, and `wiki-query` under `.github/prompts/`.
- Matched the prompt workflows to the shell-first wiki commands instead of introducing a second maintenance path.
- Documented the new on-demand prompt entrypoints in the repo docs and wiki overview.

## [2026-04-13] ingest | gated wiki refresh on actual branch changes

- Updated `make wiki-refresh` to exit early when `make wiki-ingest-candidates` finds no ingestable changes for the current diff range.
- Added the same short-circuit rule to the `wiki-refresh` prompt so chat-driven refreshes do not run a no-op maintenance cycle.

## [2026-04-13] query | documented template rendering architecture

- Expanded `wiki/repo-map.md` with the server-side template lifecycle used by MDBlog.
- Captured the split between route-specific inner templates, the shared `layout.html` wrapper, and the reusable `_post_preview.html` partial.
- Recorded that embed builds can swap template and asset loading from disk to embedded filesystems.

## [2026-04-13] query | documented embed.go purpose and wiring

- Expanded `wiki/repo-map.md` with the `embed.go` mechanism used by the Lambda embed variant.
- Recorded why `embed.go` lives at the repo root, how it exposes subtree-scoped filesystems, and where those filesystems are injected into the server.
- Noted that the embed build removes runtime `templates/` and `assets/` dependencies, but still keeps posts, pages, config, and generated SEO/feed files on disk.

## [2026-04-13] query | documented AWS runtime and no-database model

- Expanded `wiki/repo-map.md` with the AWS Lambda container runtime path and the role of API Gateway plus `algnhsa`.
- Recorded which content and generated artifacts are copied into the production image and how the embed variant differs.
- Documented why MDBlog does not need a database: Markdown files and the pre-built JSON index are the runtime data layer.

## [2026-04-13] query | documented testing model

- Expanded `wiki/repo-map.md` with the repo's Go testing approach and the `make test` workflow.
- Recorded that tests live under `internal/` and cover config, parsing, domain logic, build artifacts, handler behavior, and SEO behavior.
- Noted that the test flow regenerates derived artifacts before running `go test ./...`.

## [2026-04-13] query | documented Makefile target inventory

- Expanded `wiki/repo-map.md` with the repo's Makefile target groups for development, wiki maintenance, and Docker workflows.
- Recorded that `help` is the default goal and that `render` plus `new-post` are variable-driven command entrypoints.

## [2026-04-13] query | documented CSS theme system

- Expanded `wiki/repo-map.md` with how CSS theme selection works through `config.toml`, the layout template, and asset serving.
- Recorded the distinction between the shared `base.style.css` layer, the standalone default theme, and the anthropic theme that imports the base layer.
- Noted that light and dark mode switching is handled inside the selected stylesheet via CSS variables plus the `data-theme` attribute set by `layout.html`.

## [2026-04-13] query | documented compression and security model

- Expanded `wiki/repo-map.md` with the runtime compression path and the Lambda-specific handoff to API Gateway or CloudFront.
- Recorded the main in-repo security controls: CSP headers, safe Markdown rendering, `html/template` escaping, and path-traversal guards for posts, pages, and assets.
- Noted the deployment hardening angle of the minimal container image and baked-in content artifacts.

## [2026-04-13] query | documented feed, sitemap, and robots lifecycle

- Expanded `wiki/repo-map.md` with how `feed.xml`, `sitemap.xml`, and `robots.txt` are generated from the prebuilt post index.
- Recorded the difference between the machine-readable feed endpoint and the human-readable feed page, plus the lack of a dynamic fallback for `feed.xml`.
- Noted that sitemap and robots have dynamic dev fallbacks while production prefers the prebuilt files on disk.

## [2026-04-13] query | documented post structure

- Expanded `wiki/repo-map.md` with the Markdown post file format, filename convention, and recognized front matter keys.
- Recorded which fields are optional in practice because the runtime has title and date fallback behavior.
- Noted that per-post JavaScript comes from the `js` front matter key and `assets/js/`.

## [2026-04-13] query | documented exact Docker image contents

- Expanded `wiki/repo-map.md` with the exact files produced during the Docker build stage and copied into the standard and embed final images.
- Recorded the distinction between generated artifacts such as `feed.xml` and `sitemap.xml` versus copied runtime content such as `posts/` and `pages/`.

## [2026-04-13] query | documented search flow

- Expanded `wiki/repo-map.md` with the request flow for standalone search, including the `q` and `search=1` query parameters.
- Recorded that search is a case-insensitive substring match over indexed title, excerpt, and tags fields from `posts/posts.index.json`.
- Noted that search depends on the built index and does not fall back to a live filesystem scan when the index is missing.

## [2026-04-13] query | documented footer flow

- Expanded `wiki/repo-map.md` with how the shared footer is assembled from `footer_content`, build version metadata, and optional render timing.
- Recorded that footer content is Markdown-rendered from config and only appears on HTML pages that pass through `layout.html`.

## [2026-04-13] query | documented Markdown publication flow

- Expanded `wiki/repo-map.md` with the GitHub Actions publication path from `posts/**/*.md` changes to GHCR, ECR, and the Lambda function update.
- Recorded that the automatic push trigger is specific to `posts/**/*.md`, which is narrower than all Markdown files in the repository.
