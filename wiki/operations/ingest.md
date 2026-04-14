# Ingest Workflow

## Goal

Absorb a repo change into the wiki without rediscovering the entire codebase.

## Procedure

1. Read `wiki/index.md` and the latest entries in `wiki/log.md`.
2. Inspect the changed files first.
3. Ignore repo-specific excluded paths from [../repo-map.md](../repo-map.md).
4. Decide whether the change updates an existing page or needs a new page.
5. Update the relevant wiki page with the durable facts only.
6. Update `wiki/index.md` if page coverage changed.
7. Append an entry to `wiki/log.md`.

## Shell-First Inputs

```bash
git diff --name-only
git diff --stat
find wiki -maxdepth 2 -type f | sort
grep -R "^## " wiki
make wiki-changed WIKI_DIFF=master...HEAD
make wiki-ingest-candidates WIKI_DIFF=master...HEAD
```

## Ignore Rules

Routine wiki ingest ignores these paths unless the user explicitly asks otherwise:

- `posts/`
- `wiki/`
- `bin/`
- `render/`
- `feed.xml`
- `robots.txt`
- `sitemap.xml`
- `*.log`
- `*.tmp`

## Page Decision Rule

- Update an existing page when the change fits an existing concern.
- Create a new page when the change introduces a new subsystem, workflow, or recurring question.

## Log Format

Use this exact heading pattern:

```md
## [YYYY-MM-DD] ingest | short summary
```
