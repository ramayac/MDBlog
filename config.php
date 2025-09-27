<?php

/**
 * MDBlog Configuration
 * 
 * This file contains all configurable settings for the blog.
 * Modify these values to customize your blog's behavior.
 */

// Basic blog settings
return [
    // Basic blog settings
    'blog_name' => 'MDBlog',
    // Header and footer content
    'header_content' => 'A simple PHP + Markdown blog.',
    // Category configuration
    'categories' => [
        'old' => [
            'blog_name' => 'Old musings 🎵',
            'header_content' => 'Old thoughts before 2020.',
            'folder' => 'old',
            'index' => false, //shows it in the main index
        ],
    ],
    
    'footer_content' => '*Built with [MDBlog](https://github.com/ramayac/MDBlog). 💻*',

    // Content settings
    'posts_per_page' => 25,
    'excerpt_length' => 200,
    
    // Enable/disable uncategorized posts (posts in root posts directory)
    'show_uncategorized' => true,
    'uncategorized_label' => 'General',
    
    // Performance and debugging
    'show_render_time' => true, // Set to false to hide render time comments
    
    // Directory settings
    'posts_dir' => 'posts', // Base directory for all posts and category folders
    'includes_dir' => 'includes',
    
    // Date format (used for file modification dates)
    'date_format' => 'Y-m-d',
    
    // SEO and meta settings
    'default_meta_description' => '',
    
    // Content Security Policy
    'csp_enabled' => true,
    'csp_header' => "Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:;",

    // CSS theme
    'css_theme' => 'assets/css/default.style.css',
];
