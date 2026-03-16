<?php

/**
 * MDBlog Configuration
 * 
 * This file contains all configurable settings for the blog.
 * Modify these values to customize your blog's behavior.
 */

// Basic blog settings
return [
    'blog_name' => 'Rodrigo A. ',
    'author_name' => 'Rodrigo Amaya',
    // Header and footer content
    'header_content' => 'Wholesome Software Development.',
    // Navigation menu: custom links always shown in the nav bar.
    // Category links are added automatically when 'menu' => true in a category entry.
    'menu_links' => [
        ['label' => 'Home',  'url' => 'index.php'],
        ['label' => 'About', 'url' => 'post.php?slug=about'],
    ],

    // Category configuration
    'categories' => [
        'personal' => [
            'blog_name' => 'Personal 🏠',
            'header_content' => 'Personal Thoughts and Musings.',
            'folder' => 'personal',
            'index' => true,  // shows it in the main index
            'menu'  => true,  // shows it in the nav menu
        ],
        'srbyte' => [
            'blog_name' => 'Sr. Byte 👨‍💻',
            'header_content' => '(2007-2010) Tecnologia para todos.',
            'folder' => 'srbyte',
            'index' => false, // shows it in the main index
            'menu'  => true,  // shows it in the nav menu
        ],
        'substack' => [
            'blog_name' => 'Code Forward ⏩',
            'header_content' => '(2023) Substack failed newsletter.',
            'folder' => 'substack',
            'index' => false, // shows it in the main index
            'menu'  => true,  // shows it in the nav menu
        ],
        'mdblog' => [
            'blog_name' => 'MDBlog 📝',
            'header_content' => '(2024) Instructions to use MDBlog.',
            'folder' => 'mdblog',
            'index' => true, // shows it in the main index
            'menu'  => true,  // shows it in the nav menu
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
    'cache_dir' => 'cache',

    // Cache settings
    'cache_enabled' => true,
    'cache_ttl' => 604800, // 1 week in seconds (60 * 60 * 24 * 7)
    
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
