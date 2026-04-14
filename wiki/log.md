# Wiki Log

## [2026-04-13] bootstrap | established repo-local wiki scaffold

- Added a dedicated `wiki/` directory with a stable index, log, schema, operations, and rollout plan.
- Defined MDBlog-specific exclusions so routine wiki maintenance ignores `posts/` unless explicitly requested.
- Added repo instructions so the agent reads the wiki before broad analysis and files durable findings back into it.

## [2026-04-13] ingest | added shell-first wiki helper targets

- Added `wiki-*` Make targets for listing, searching, diff-driven ingest, linting, and a combined refresh snapshot.
- Added plain `sh` helper scripts under `scripts/` so the workflow remains unix-friendly and portable.
- Documented the new manual entrypoints in the repo docs and wiki operations pages.
