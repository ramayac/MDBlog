---
title: Interactive Processing.js Sketch
date: 2024-01-25
author: Creative Coder
tags: processing, javascript, interactive, art
js: processing.min.js
description: A blog post with an embedded Processing.js sketch for interactive content
---

# Interactive Processing.js Sketch

This post demonstrates how to include interactive Processing.js sketches in your MDBlog posts.

## The Magic of Processing.js

Processing.js brings the Processing visualization language to the web. You can create interactive animations, generative art, and data visualizations right in your blog posts.

## How It Works

To include a Processing.js sketch in your post:

1. Add the JavaScript file to the front matter using the `js` property
2. Upload your Processing.js library to `assets/js/`
3. Include your sketch code in a code block or separate file

## Sample Sketch Code

Here's the Processing code for a simple interactive sketch:

```javascript
// This would be processed by Processing.js
void setup() {
    size(400, 300);
    background(50);
}

void draw() {
    // Draw colorful circles that follow the mouse
    fill(random(255), random(255), random(255), 100);
    noStroke();
    ellipse(mouseX + random(-20, 20), mouseY + random(-20, 20), 30, 30);
    
    // Fade the background slightly
    fill(50, 10);
    rect(0, 0, width, height);
}
```

## Multiple JavaScript Files

You can also include multiple JavaScript files by using an array in the front matter:

```yaml
js: 
  - processing.min.js
  - sketch.js
  - custom-functions.js
```

## Use Cases

Processing.js sketches are perfect for:

* **Data visualization** - Interactive charts and graphs
* **Generative art** - Algorithmic and procedural art
* **Educational content** - Interactive demonstrations
* **Games** - Simple browser-based games
* **Simulations** - Physics simulations, particle systems

## Getting Started

1. Download Processing.js from processingjs.org
2. Upload it to your `assets/js/` folder
3. Create your sketch and add the `js` front matter
4. Publish your post!

*Note: Make sure to include the Processing.js library file in your assets folder for the interactive features to work.*