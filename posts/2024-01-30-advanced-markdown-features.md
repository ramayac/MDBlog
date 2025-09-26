---
title: Advanced Markdown Features with Parsedown
date: 2024-01-30
author: Markdown Master
tags: markdown, parsedown, formatting
description: Showcasing the advanced Markdown features available with Parsedown parser
---

# Advanced Markdown Features with Parsedown

Now that we're using **Parsedown**, we have access to much more advanced Markdown features!

## Text Formatting

You can use all the standard formatting options:

- **Bold text** with `**bold**` or `__bold__`
- *Italic text* with `*italic*` or `_italic_`
- ***Bold and italic*** with `***text***`
- ~~Strikethrough~~ with `~~text~~`
- `Inline code` with backticks

## Lists

### Ordered Lists

1. First item
2. Second item
   1. Nested item
   2. Another nested item
3. Third item

### Unordered Lists

- Item one
- Item two
  - Nested item
  - Another nested item
- Item three

### Task Lists

- [x] Completed task
- [ ] Incomplete task
- [x] Another completed task

## Tables

| Feature | Status | Notes |
|---------|--------|-------|
| Headers | ✅ | All levels supported |
| Lists | ✅ | Ordered and unordered |
| Tables | ✅ | With alignment support |
| Code | ✅ | Inline and blocks |
| Links | ✅ | Internal and external |

## Blockquotes

> This is a blockquote. It can contain multiple paragraphs.
> 
> Like this second paragraph.

> Nested blockquotes are also possible:
> 
> > This is nested inside the first blockquote.

## Code Examples

Here's some syntax-highlighted code:

```php
<?php
class ParsedownBlog {
    private $parsedown;
    
    public function __construct() {
        $this->parsedown = new Parsedown();
        $this->parsedown->setBreaksEnabled(true);
    }
    
    public function render($markdown) {
        return $this->parsedown->text($markdown);
    }
}
?>
```

```javascript
// JavaScript with proper highlighting
function processMarkdown(text) {
    const parser = new MarkdownParser();
    return parser.parse(text);
}

const result = processMarkdown('# Hello World');
console.log(result);
```

## Links and Images

Here's a [link to Parsedown](https://parsedown.org/) documentation.

You can also use reference-style links: [Parsedown GitHub][parsedown-repo].

[parsedown-repo]: https://github.com/erusev/parsedown

## Horizontal Rules

You can create horizontal rules with three or more dashes:

---

Or with asterisks:

***

## HTML Support

Parsedown allows <strong>HTML tags</strong> to be mixed with Markdown, giving you even more <em>formatting flexibility</em>.

<div style="background: #f0f8ff; padding: 1rem; border-radius: 4px;">
This is HTML mixed with **Markdown** formatting!
</div>

## Mathematical Expressions

While Parsedown doesn't natively support LaTeX math, you can always add KaTeX or MathJax via the JavaScript front matter feature:

```
E = mc²
```

## Conclusion

With Parsedown, MDBlog now supports the full CommonMark specification plus GitHub Flavored Markdown extensions. This makes it much more powerful while keeping the same simple workflow!

*Much better than our custom parser, right?* 😉