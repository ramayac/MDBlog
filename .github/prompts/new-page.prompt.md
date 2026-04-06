---
description: "Create a static standalone page like About, Now, Uses, or Colophon. Use when: adding an about page, creating a static page, making a standalone page."
agent: "agent"
argument-hint: "What kind of page? (e.g. About, Now, Uses, Projects...)"
---

You are creating a static/standalone page for an MDBlog blog. MDBlog doesn't have a built-in "pages" system separate from posts, but standalone content pages can be implemented as regular posts filed under a dedicated category.

## Project context

- Read [config.toml](../../config.toml) to get `author_name`, `blog_name`, and existing categories.
- Posts live under `posts/<category>/` as Markdown files with YAML front matter.
- The blog supports an optional `posts/index.md` file that renders as a blurb on the landing page (above category cards). This is the simplest way to add a homepage introduction.

## Strategy for static pages

There are two approaches depending on what the user wants:

### Option A: Landing page blurb (`posts/index.md`)
If the user wants an "about" section on the homepage, create or update `posts/index.md`. This file has **no front matter** — it's pure Markdown rendered above the category cards.

### Option B: Dedicated page as a post
For standalone pages (About, Now, Uses, Colophon, etc.), create a regular post. Consider:
- Filing it under an existing category, or suggesting a new `pages` or `guides` category.
- Using a date-less or backdated filename if the page is evergreen.
- Adding a `[[menu_links]]` entry in `config.toml` so it appears in the nav bar.

## Front matter for standalone pages

```yaml
---
title: About
date: YYYY-MM-DD
author: <author_name from config.toml>
tags:
description: A short description of this page.
---
```

## Steps

1. Ask the user what kind of page they want (About, Now, Uses, etc.) if not already specified.
2. Ask where it should live — homepage blurb (`posts/index.md`) or standalone post.
3. If standalone: pick or create a category, write the Markdown file, and offer to add a `[[menu_links]]` entry to [config.toml](../../config.toml) for nav bar visibility.
4. Draft the page content based on the user's input, using a clean structure appropriate for the page type.
5. Remind the user to run `make build-index` after creating the file.

## Page type templates

**About page:** Brief intro → what you do → what this blog covers → contact/links.
**Now page** (nownownow.com style): What you're focused on right now — projects, reading, interests.
**Uses page:** Hardware, software, tools, and setup you use daily.
**Colophon:** How the blog is built, what tech stack powers it, credits.
