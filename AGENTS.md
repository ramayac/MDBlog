# MDBlog — Agent Instructions

## Project Overview

MDBlog is a lightweight, flat-file blog engine written in PHP. Posts are Markdown files with YAML front matter. There is no database, no build step, and no JavaScript framework. Keep it that way.

## Architecture

```
index.php          # Post listing (pagination, categories)
post.php           # Single post view
config.php         # All runtime configuration (single source of truth)
includes/
  Blog.php          # Core logic: post scanning, slug resolution, pagination
  MarkdownParser.php # Front matter parser + Parsedown wrapper
  Parsedown.php     # Third-party Markdown renderer (do not modify)
  head.php          # HTML <head> template, CSP header
  debug.php         # Render-time debug output
  templates/        # Reusable HTML partials
assets/
  css/              # CSS themes
  js/               # Optional per-post JavaScript files
posts/              # All content lives here
  *.md              # Root-level posts (uncategorized)
  srbyte/           # Category folder
  substack/         # Category folder
cache/              # JSON cache files (auto-generated, do not commit)
Makefile            # Developer targets: help, serve, new-post
```

## Configuration

All settings live in `config.php`. Key fields:

| Key | Purpose |
|-----|---------|
| `author_name` | Default post author (used by `make new-post`) |
| `categories` | Named category folders with display metadata |
| `posts_per_page` | Pagination size |
| `csp_enabled` / `csp_header` | Content Security Policy |
| `css_theme` | Path to the active CSS theme |
| `show_render_time` | Toggles render-time HTML comment in debug.php |

When adding a new config key, add it to `config.php` with a comment and consume it through `$blog->getConfig('key')` or the `$config` array — never hardcode values in PHP templates.

## Post Format

```yaml
---
title: My Post Title
date: YYYY-MM-DD
author: Name
tags: tag1, tag2
description: Optional meta description
js: optional-script.js   # loaded from assets/js/
---

Markdown body here.
```

File naming convention: `YYYY-MM-DD-slug-with-hyphens.md`  
Category posts live under `posts/<category-folder>/`. The filename does **not** need the date prefix for category posts (see existing srbyte/substack posts), but the date prefix is preferred for new content.

## Developer Workflow

```bash
make serve                                      # PHP dev server at localhost:8080
make new-post TITLE="Title" TAGS="tag1, tag2"   # Scaffold a new post
```

## Coding Conventions

- **PHP 7.0+ compatible.** Do not use PHP 8-only syntax (named arguments, `match` expressions, nullsafe operator `?->`) without verifying the minimum version in `README.md`.
- **No frameworks, no Composer dependencies** beyond the bundled `Parsedown.php`.
- **XSS prevention:** all user-visible output must pass through `htmlspecialchars()`. Posts are rendered via Parsedown with `setSafeMode(true)` — do not disable this.
- **Path traversal prevention:** `getPostBySlug()` already validates slugs. Any new file-reading code must validate user input before constructing filesystem paths.
- **CSP:** the `csp_header` in `config.php` controls the policy. If adding new external resource types (fonts, embeds, etc.), update the policy there, not inline.
- **Output buffering / gzip:** `index.php` and `post.php` start `ob_gzhandler` when the client supports gzip. Any new entry-point file should do the same.
- **No globals:** pass `$config` explicitly or use `$blog->getConfig()`. Do not introduce global state.

## Adding a New Category

1. Create `posts/<folder-name>/` and add `.md` files.
2. Register the category in `config.php` under `categories`:
   ```php
   'my-category' => [
       'blog_name'      => 'Display Name',
       'header_content' => 'Subtitle shown on category index.',
       'folder'         => 'my-category',
       'index'          => true, // false = hidden from main index
   ],
   ```

## Cache

JSON files in `cache/` are auto-generated at runtime. Do not edit them manually. Do not commit them (add to `.gitignore` if not already present).

## What NOT to Do

- Do not add a database layer.
- Do not introduce a front-end build tool (webpack, vite, etc.).
- Do not modify `Parsedown.php` — update it by replacing the file wholesale.
- Do not disable Parsedown's safe mode or markup escaping.
- Do not add `unsafe-eval` to the CSP without a documented security justification.
