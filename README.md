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

## Quick Start

1. Upload files to a PHP-enabled web server
2. Update config.php for blog name, etc
2. Create `.md` files in the `posts` directory
3. That's it!

## Running Locally

A `Makefile` is included for local development. Requires PHP and `make`.

```bash
# Start the built-in PHP dev server (default: http://localhost:8080)
make serve

# Use a custom host/port
make serve HOST=0.0.0.0 PORT=9000
```

## Creating a New Post

```bash
make new-post TITLE="My Post Title" TAGS="tag1, tag2"
```

This creates a pre-filled `.md` file in `posts/` named `YYYY-MM-DD-my-post-title.md` with front matter populated from `config.php`.

## Writing Posts

Create a `.md` file in `posts/` with optional front matter:

```yaml
---
title: My Post
date: 2024-01-15
author: Your Name
tags: tag1, tag2
js: script.js
---

Your markdown content here...
```

## Configuration

Edit `config.php` to customize blog settings like posts per page and theme.

## Requirements

- PHP 7.0+
- Web server with PHP support
- No database required
- `zlib` PHP extension (enabled by default) for gzip compression

## License

MIT License
