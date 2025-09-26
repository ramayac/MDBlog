<?php

/**
 * MDBlog Configuration
 * 
 * This file contains all configurable settings for the blog.
 * Modify these values to customize your blog's behavior.
 */

return [
    // Basic blog settings
    'blog_name' => 'MDBlog',
    'blog_description' => 'A simple markdown-based blog system',
    
    // Content settings
    'posts_per_page' => 25,
    'excerpt_length' => 150,
    
    // Performance and debugging
    'show_render_time' => true, // Set to false to hide render time comments
    
    // Directory settings
    'posts_dir' => 'posts',
    'includes_dir' => 'includes',
    
    // Date format (used for file modification dates)
    'date_format' => 'Y-m-d',
    
    // SEO and meta settings
    'default_meta_description' => '',

    // CSS theme
    'css_theme' => 'assets/css/default.style.css',
];