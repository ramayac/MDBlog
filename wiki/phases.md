# Wiki Engine Phases

## Status Board

| Phase | Name | Status | Exit Signal |
|---|---|---|---|
| 0 | Bootstrap the wiki | completed | Required files exist and Docker ignores `wiki/` |
| 1 | Teach the agent the workflow | completed | Repo instructions say to read and update the wiki |
| 2 | Establish repo map baseline | completed | Repo architecture and exclusions are recorded |
| 3 | Standardize ingest | completed | New repo changes can be filed with a repeatable checklist |
| 4 | Standardize query | completed | Answers start from the wiki before raw source scans |
| 5 | Standardize lint | completed | Drift checks and repair steps are documented |
| 6 | Add unix-native helpers | completed | Common wiki inspection commands are stable across repos |
| 7 | Add change-detection workflow | completed | File-change driven wiki updates are documented |
| 8 | Define reusable repo template | planned | New repos can copy the same wiki skeleton |
| 9 | Add optional on-demand automation | completed | Manual triggers can refresh the wiki without CI lock-in |
| 10 | Prove multi-repo adoption | planned | At least one more repo uses the same structure |
| 11 | Extract wiki-engine repo | future | Shared engine lives outside any single product repo |

## Phase 0 | Bootstrap the wiki

Goal: create the minimal scaffold that every repo can share.

Deliverables:

- `wiki/` exists.
- `wiki/index.md`, `wiki/log.md`, `wiki/schema.md`, `wiki/phases.md`, and `wiki/repo-map.md` exist.
- `wiki/operations/` contains ingest, query, and lint procedures.
- Docker build context excludes `wiki/`.

## Phase 1 | Teach the agent the workflow

Goal: make the wiki part of normal repo analysis instead of a side document.

Deliverables:

- Root instructions mention the wiki-first read order.
- An on-demand instruction exists for wiki maintenance tasks.
- The repo documents when durable answers should be written back into the wiki.

## Phase 2 | Establish repo map baseline

Goal: capture the current repo model once so future sessions can update it incrementally.

Deliverables:

- Architecture summary.
- Build and deploy path summary.
- Important directories and generated artifacts.
- Repo-specific exclusions.

## Phase 3 | Standardize ingest

Goal: give the agent a repeatable way to absorb new code or docs into the wiki.

Deliverables:

- File-change checklist.
- Rules for when to update existing pages versus create new ones.
- Log entry format for ingests.

## Phase 4 | Standardize query

Goal: answer repo questions from the wiki first and only widen to source files when needed.

Deliverables:

- Search order.
- Escalation rules from wiki pages to source files.
- Rules for filing durable answers back into the wiki.

## Phase 5 | Standardize lint

Goal: keep the wiki healthy as it grows.

Deliverables:

- Orphan-page checks.
- Stale-claim checks.
- Missing-link and missing-page checks.
- Data-gap capture for future work.

## Phase 6 | Add unix-native helpers

Goal: keep the workflow compatible with `find`, `grep`, `sed`, `awk`, `ls`, and `tree`.

Deliverables:

- Stable filename patterns.
- Parseable log headings.
- A small set of documented shell commands that work across repos.

Implemented in MDBlog:

- `make wiki-list`
- `make wiki-headings`
- `make wiki-log-tail`
- `make wiki-search`

## Phase 7 | Add change-detection workflow

Goal: make wiki maintenance faster for ongoing projects.

Deliverables:

- Diff-driven ingest instructions.
- Rules for ignoring generated or user-authored noise.
- Optional lightweight scripts if manual grep is no longer enough.

Implemented in MDBlog:

- `make wiki-changed`
- `make wiki-candidates`
- Ingest filtering excludes `posts/`, generated artifacts, and temporary noise.

## Phase 8 | Define reusable repo template

Goal: make new repos adopt the same wiki layout with minimal changes.

Deliverables:

- Canonical folder structure.
- Required pages.
- Repo-map template with an exclusions section.
- Minimal adoption checklist.

## Phase 9 | Add optional on-demand automation

Goal: support manual refresh workflows without requiring a heavy CI pipeline.

Deliverables:

- Clear manual entrypoints.
- Optional local script or container strategy.
- No mandatory hosted service dependency.

Implemented in MDBlog:

- `make wiki-lint`
- `make wiki-refresh`
- All helper entrypoints run locally through plain `sh` scripts.

## Phase 10 | Prove multi-repo adoption

Goal: verify that the pattern survives outside MDBlog.

Deliverables:

- At least one second repo using the same `wiki/` contract.
- Notes on what remained stable versus what had to vary per repo.

## Phase 11 | Extract wiki-engine repo

Goal: move the reusable parts into a separate repository once the contract is stable.

Deliverables:

- Shared scaffold.
- Shared instructions.
- Optional reusable scripts or container entrypoints.
- Migration notes for existing repos.
