# MDBlog

A simple, static blog system powered by Markdown and PHP. Similar to Kirby CMS, but focused on simplicity and ease of use.

## Features

- ✨ **Markdown-based** - Write posts in Markdown format
- 🚀 **Static generation** - No database required
- 📄 **Pagination** - Built-in pagination (25 posts per page by default)
- 🎨 **Code highlighting** - Syntax highlighting for code blocks
- 📱 **Responsive design** - Mobile-friendly layout
- ⚡ **Custom JavaScript** - Add Processing.js sketches or other scripts to specific posts
- 🎯 **Simple deployment** - Just upload files to any PHP-enabled server
- 📝 **Front matter support** - YAML front matter for post metadata

## Installation

1. Clone or download this repository
2. Upload the files to your PHP-enabled web server
3. Ensure the web server has read permissions for all directories
4. That's it! No configuration required.

## Usage

### Creating Posts

1. Create a new `.md` file in the `posts` directory
2. Add front matter at the top of your file (optional):

```yaml
---
title: Your Post Title
date: 2024-01-15
author: Your Name
tags: tag1, tag2, tag3
description: A short description of your post
js: script.js  # Optional: include custom JavaScript
---
```

3. Write your content in Markdown below the front matter
4. Save the file and refresh your blog

### Front Matter Options

- `title` - Post title (defaults to filename if not specified)
- `date` - Publication date (defaults to file modification time)
- `author` - Author name
- `tags` - Comma-separated tags or YAML array
- `description` - Meta description for SEO
- `js` - JavaScript file(s) to include (string or array)

### Custom JavaScript

To include custom JavaScript files (like Processing.js sketches):

1. Upload your JavaScript files to `assets/js/`
2. Reference them in your post's front matter:

```yaml
js: processing.min.js
```

Or multiple files:

```yaml
js:
  - processing.min.js
  - sketch.js
```

### Customizing Header and Footer

Edit the following files to customize your blog's header and footer:

- `includes/header.md` - Site header content
- `includes/footer.md` - Site footer content

Both files support full Markdown syntax.

## File Structure

```
MDBlog/
├── assets/
│   ├── css/
│   │   └── style.css          # Main stylesheet
│   └── js/                    # JavaScript files
├── includes/
│   ├── Blog.php               # Main blog class
│   ├── MarkdownParser.php     # Markdown parser wrapper
│   ├── Parsedown.php          # Parsedown library
│   ├── header.md              # Site header
│   └── footer.md              # Site footer
├── posts/                     # Your blog posts (.md files)
├── index.php                  # Main blog page
├── post.php                   # Single post view
└── README.md                  # This file
```

## Customization

### Styling

Edit `assets/css/style.css` to customize the appearance of your blog.

### Posts per Page

To change the number of posts per page, modify the `Blog` class instantiation in `index.php`:

```php
$blog = new Blog('posts', 10); // 10 posts per page instead of 25
```

### Markdown Parser

The blog uses [Parsedown](https://parsedown.org/), a fast and lightweight PHP Markdown parser that supports:

- Headers (# ## ###)
- Bold and italic text (**bold**, *italic*)
- Links and images
- Code blocks with syntax highlighting
- Lists (ordered and unordered)
- Inline code
- Blockquotes
- Tables
- Strikethrough
- And much more!

Parsedown is included in the project, so no additional setup is required.

## Requirements

- PHP 7.0 or higher
- Web server with PHP support (Apache, Nginx, etc.)
- No database required

## Example Posts

The installation includes several example posts demonstrating various features:

1. Welcome post with basic Markdown
2. Code examples with syntax highlighting  
3. Processing.js integration example

## License

MIT License - feel free to use this for any project!

## Contributing

Feel free to submit issues and pull requests to improve MDBlog.