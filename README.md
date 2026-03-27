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

MDBlog is deployed as a Docker container image on AWS Lambda behind API Gateway, using [Bref](https://bref.sh/).

1. Clone the repo and configure `config.php` (blog name, author, categories)
2. Add `.md` posts to the `posts/<category>/` directories
3. Build and push the Docker image: `make docker-build && make docker-push`
4. Deploy the image to AWS Lambda as a container image function behind API Gateway

For local development, see the [Running Locally](#running-locally) section below.

## Running Locally

A `Makefile` is included for local development. Requires PHP 8.3+, Composer, and `make`.

**Install dependencies first:**

```bash
composer install
```

Then start the dev server:

```bash
make serve           # Start the built-in PHP dev server (default: http://localhost:8080)
make lint            # Check all PHP files for syntax errors
make clear-cache     # Delete all cached .json files from cache/
make utf8-fix        # Re-encode any non-UTF-8 .md files in posts/ to valid UTF-8
```

## Deployment (AWS Lambda)

The production image is built on [`bref/php-83-fpm:2`](https://bref.sh/) and deployed as an **AWS Lambda container image function** behind API Gateway. Bref's runtime translates API Gateway HTTP events into PHP-FPM requests automatically. Posts are bundled inside the Docker image — there is no database and no external content source.

```bash
make docker-build                        # Bake version.php + build image
make docker-push REGISTRY=ghcr.io/...   # Tag and push to a container registry
make docker-pull TAG=1.2.3              # Pull a release image and tag as latest
```

After pushing the image, update the Lambda function to use the new image URI. AWS will cold-start new containers from the updated image.

### Cache on Lambda

The file-based JSON cache (`cache/`) **does not work on Lambda** and is disabled by default (`cache_enabled => false` in `config.php`):

- Lambda's container filesystem is immutable (read-only except `/tmp`)
- Each concurrent Lambda container has its own isolated `/tmp` — containers cannot share cache state
- Cold starts reset any `/tmp` content

**Recommended caching strategy for Lambda:** place **CloudFront** (or enable **API Gateway caching**) in front of the Lambda function. Since posts are baked into the image and change only on redeploy, a CloudFront TTL of several hours or days works well. Invalidate the distribution after each `make docker-push`.

The file-based cache code is retained for anyone self-hosting on a traditional web server — enable it by setting `cache_enabled => true` in `config.php`.

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

**Production (Lambda):**
- Docker (to build the image)
- AWS account with Lambda + API Gateway (+ optionally CloudFront)
- Container registry (e.g. ghcr.io or ECR) to host the image

**Local development:**
- PHP 8.3+
- Composer
- `make`
- `zlib` PHP extension (enabled by default) for gzip compression

No database required.

## License

MIT License
