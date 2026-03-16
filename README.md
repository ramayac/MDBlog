# MDBlog

A simple Markdown-based blog system powered by PHP. 
No database required, Javascript or any BS.

Made by [@ramayac](https://x.com/ramayac).


## Features

- Write posts in Markdown
- Built-in pagination
- Code syntax highlighting
- Custom JavaScript support
- Responsive design
- Support for Dark/Light OS selection
- Gzip compression (when supported by the client)
- Dynamic navigation menu (driven by `config.php`)
- Landing page with category cards (no full post scan on homepage)

## Quick Start

1. Upload files to a PHP-enabled web server
2. Update `config.php` for blog name, author, categories, and menu links
3. Create `.md` files in the `posts/` directory
4. That's it!

## Running Locally

A `Makefile` is included for local development. Requires PHP and `make`.

```bash
# Start the built-in PHP dev server (default: http://localhost:8080)
make serve

# Use a custom host/port
make serve HOST=0.0.0.0 PORT=9000
```

## Docker

Requires Docker with the Compose plugin.

```bash
# Build the image (bakes version.php automatically)
make docker-build

# Run at http://localhost:8080
make docker-run

# Stop and remove containers
make docker-stop

# Push to a registry
make docker-push REGISTRY=ghcr.io/your-user/mdblog
```

The container runs as a non-root user with a read-only filesystem. Only `cache/`, `/tmp`, and `/run` are writable (tmpfs).

## Creating a New Post

```bash
make new-post TITLE="My Post Title" TAGS="tag1, tag2"
```

Creates a pre-filled `.md` file in `posts/` named `YYYY-MM-DD-my-post-title.md`.  
Author is read automatically from `config.php`. `TAGS` is optional.

## Writing Posts

Create a `.md` file in `posts/` (or a category subfolder) with front matter:

```yaml
---
title: My Post
date: 2024-01-15
author: Your Name
tags: tag1, tag2
description: Optional meta description
js: optional-script.js
---

Your markdown content here...
```

## Navigation Menu

The nav bar is generated automatically from `config.php` — no more editing `menu.md`.

```php
// Static custom links (always shown, in order)
'menu_links' => [
    ['label' => 'Home',  'url' => 'index.php'],
    ['label' => 'About', 'url' => 'post.php?slug=about'],
],

// Per-category: set 'menu' => true to include it in the nav
'categories' => [
    'srbyte' => [
        ...
        'menu' => true,
    ],
],
```

## Landing Page

`index.php` (no `?category`) shows a static landing page with category cards.  
To add an optional intro blurb above the cards, create `posts/index.md` with any Markdown content.

Browsing posts works via `index.php?category=slug`.

## Categories

Register categories in `config.php`:

```php
'categories' => [
    'my-category' => [
        'blog_name'      => 'Display Name',
        'header_content' => 'Subtitle.',
        'folder'         => 'my-category',   // subfolder under posts/
        'index'          => true,            // include in aggregated index (legacy)
        'menu'           => true,            // show in nav bar
    ],
],
```

Then add `.md` files to `posts/my-category/`.

## Configuration

Edit `config.php` to customize all settings. Key fields:

| Key | Purpose |
|-----|---------|
| `blog_name` | Site title |
| `author_name` | Default author for new posts |
| `header_content` | Landing page subtitle |
| `menu_links` | Static nav links |
| `categories` | Category definitions |
| `posts_per_page` | Pagination size |
| `cache_enabled` / `cache_ttl` | JSON post cache |
| `csp_enabled` / `csp_header` | Content Security Policy |
| `css_theme` | Active CSS theme path |

## Requirements

- PHP 7.0+
- Web server with PHP support (or Docker)
- No database required
- `zlib` PHP extension (enabled by default) for gzip compression

## License

MIT License
