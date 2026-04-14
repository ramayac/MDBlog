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
