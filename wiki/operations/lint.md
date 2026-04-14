# Lint Workflow

## Goal

Keep the wiki coherent, linked, and current.

## Checks

- Pages mentioned in the index still exist.
- Important repo areas have coverage.
- Stale claims are updated when source files changed.
- Exclusions still match repo reality.
- New recurring topics have a page instead of being trapped in chat history.

## Shell-First Checks

```bash
find wiki -type f | sort
grep -R "TODO\|TBD\|UNKNOWN" wiki
grep -R "posts/" wiki
grep "^## \[" wiki/log.md | tail -10
```

## Repair Order

1. Fix stale or incorrect topic pages.
2. Fix `wiki/index.md` links or summaries.
3. Append a log entry if the lint changed durable content.

## Log Format

Use this exact heading pattern:

```md
## [YYYY-MM-DD] lint | short summary
```
