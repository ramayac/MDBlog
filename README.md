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

## Quick Start

1. Upload files to a PHP-enabled web server
2. Update config.php for blog name, etc
2. Create `.md` files in the `posts` directory
3. That's it!

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

## License

MIT License
