# MDBlog — Agent Instructions

## Project Overview

MDBlog is a lightweight, flat-file blog engine written in PHP. Posts are Markdown files with YAML front matter. There is no database, no build step, and no JavaScript framework. Keep it that way.

## Architecture

```
index.php          # Landing page (category cards) OR category post listing
post.php           # Single post view
config.php         # All runtime configuration (single source of truth)
includes/
  Blog.php          # Core logic: post scanning, slug resolution, pagination, menu
  MarkdownParser.php # Front matter parser + Parsedown wrapper
  head.php          # HTML <head> template, CSP header
  debug.php         # Render-time debug output
  templates/        # Reusable HTML partials
vendor/             # Composer-managed dependencies (not committed; auto-generated in Docker build)
assets/
  css/              # CSS themes
  js/               # Optional per-post JavaScript files
posts/              # All content lives here
  index.md          # Optional landing page blurb (above category cards)
  *.md              # Root-level posts (uncategorized, legacy)
  srbyte/           # Category folder
  substack/         # Category folder
cache/              # JSON cache files (auto-generated, do not commit)
Makefile            # Developer targets: help, serve, new-post, docker-*
Dockerfile          # Production image — Bref FPM base (Lambda-compatible via API Gateway)
docker/
  nginx.conf        # nginx config (port 8080, blocks internal paths) — local/dev only
  php.ini           # PHP hardening overrides (mounted into Bref at /opt/bref/etc/php/conf.d/custom.ini)
  entrypoint.sh     # Starts php-fpm then execs nginx under tini — local/dev only
```

## Configuration

All settings live in `config.php`. Key fields:

| Key | Purpose |
|-----|---------|
| `author_name` | Default post author (used by `make new-post`) |
| `menu_links` | Ordered array of static nav links `[['label'=>'', 'url'=>'']]` |
| `categories` | Named category folders with display metadata |
| `posts_per_page` | Pagination size |
| `csp_enabled` / `csp_header` | Content Security Policy |
| `css_theme` | Path to the active CSS theme |
| `show_render_time` | Toggles render-time HTML comment in debug.php |
| `cache_enabled` / `cache_ttl` | JSON cache toggle and TTL in seconds |

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
Category posts live under `posts/<category-folder>/`.

## Developer Workflow

```bash
make serve                                      # PHP dev server at localhost:8080
make new-post TITLE="Title" TAGS="tag1, tag2"   # Scaffold a new post
make version                                    # Bake git commit/tag into version.php
make clear-cache                                # Delete all cache/*.json files
```

`make new-post` reads `author_name` from `config.php` automatically via a PHP one-liner.

## Docker Workflow

```bash
make docker-build                               # Bake version.php + build image
make docker-run                                 # Run via docker compose (localhost:8080)
make docker-run-release                         # Run pulled release image without rebuilding
make docker-stop                                # Stop and remove containers
make docker-push                                # Tag and push to ghcr.io/ramayac/mdblog
make docker-pull [TAG=1.2.3]                    # Pull release image and retag as mdblog:latest
```

The image is built from `Dockerfile` using `bref/php-83-fpm:2` as the base. This image includes the Lambda Runtime Interface Client (RIC); when deployed to AWS Lambda behind API Gateway, Bref's runtime translates API Gateway HTTP events into PHP-FPM requests automatically. `docker/php.ini` is loaded at `/opt/bref/etc/php/conf.d/custom.ini`.

**When modifying the Makefile**, always update `README.md` and `AGENTS.md` to reflect the new or changed targets.

## Navigation Menu

The nav bar is built by `Blog::getMenu()` — **do not edit `includes/menu.md`**.  
Two sources, rendered in order:
1. `menu_links` array in `config.php` — static custom links.
2. Any category with `'menu' => true` in its config — auto-generates a `?category=slug` link.

## Category Config Flags

Each entry in `config.php` `categories` supports:

| Flag | Effect |
|------|--------|
| `index` | `true` = posts appear in legacy aggregated listing |
| `menu` | `true` = category link added to nav bar |

## Landing Page vs Category Page (`index.php`)

`index.php` has **two rendering paths**:

- **No `?category`** → Static landing page. Shows blog header, optional `posts/index.md` blurb, then auto-generated category cards (one per category with posts). **No post scanning, nothing cached.**
- **`?category=slug`** → Category post listing with pagination. Cache applies here.

This split means the homepage is always fast and never stale. The expensive post scan only runs when a user browses a specific category.

To add a landing page blurb: create `posts/index.md` with Markdown content. Delete it to show only the category cards.

## Coding Conventions

- **PHP 8.3+ required.** PHP 8.x syntax (named arguments, `match` expressions, nullsafe operator `?->`, `array_key_last()`, etc.) is fine to use.
- **No frameworks, no Composer dependencies** beyond `bref/bref` and `erusev/parsedown` (managed via `composer.json`).
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
       'index'          => false, // legacy aggregated index
       'menu'           => true,  // show in nav bar
   ],
   ```

## Cache

JSON files in `cache/` are auto-generated at runtime. Do not edit them manually. Do not commit them (add to `.gitignore` if not already present).

Cache is **scoped per category folder** — adding a post only invalidates that category's cache, not the whole site. The landing page does no caching because it does no post scanning.

## What NOT to Do

- Do not add a database layer.
- Do not introduce a front-end build tool (webpack, vite, etc.).
- Do not modify `Parsedown.php` directly — it is managed by Composer (`erusev/parsedown`). Update the version constraint in `composer.json` instead.
- Do not disable Parsedown's safe mode or markup escaping.
- Do not add `unsafe-eval` to the CSP without a documented security justification.
- Do not put new content in the root `posts/` dir — use a category folder.
- Do not run the container as root or re-add dropped capabilities without justification.
- Do not modify `docker/nginx.conf` to expose `includes/`, `posts/`, or `cache/` directly.
- **When modifying the Makefile**, always update `README.md` and `AGENTS.md` to match.
