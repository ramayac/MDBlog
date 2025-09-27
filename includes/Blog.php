<?php

require_once 'includes/MarkdownParser.php';

class Blog {
    // Regex patterns
    private const DATE_PREFIX_REGEX = '/^\d{4}-\d{2}-\d{2}-/';
    private const SLUG_CLEANUP_REGEX = '/[^a-z0-9]+/';
    private const WHITESPACE_CLEANUP_REGEX = '/\s+/';
    
    private $config;
    private $postsDir;
    private $parser;
    private $postsPerPage;
    private $excerptLength;
    private $dateFormat;
    
    public function __construct($configFile = 'config.php') {
        // Load configuration
        $this->config = require $configFile;
        
        // Set properties from config
        $this->postsDir = $this->config['posts_dir'];
        $this->parser = new MarkdownParser();
        $this->postsPerPage = $this->config['posts_per_page'];
        $this->excerptLength = $this->config['excerpt_length'];
        $this->dateFormat = $this->config['date_format'];
    }
    
    public function getConfig($key = null) {
        if ($key === null) {
            return $this->config;
        }
        return isset($this->config[$key]) ? $this->config[$key] : null;
    }
    
    public function getPosts($page = 1, $categorySlug = null) {
        if ($categorySlug) {
            $category = $this->getCategoryBySlug($categorySlug);
            if (!$category) {
                return ['posts' => [], 'pagination' => ['current' => 1, 'total' => 1, 'hasPrev' => false, 'hasNext' => false]];
            }
            $posts = $this->scanPostsInFolder($this->postsDir . '/' . $category['folder'], $categorySlug);
        } else {
            $posts = $this->scanAllPosts($this->postsDir);
        }
        
        // Sort posts by date (newest first)
        usort($posts, [$this, 'sortPostsByDate']);
        
        // Calculate pagination
        $totalPosts = count($posts);
        $totalPages = ceil($totalPosts / $this->postsPerPage);
        $offset = ($page - 1) * $this->postsPerPage;
        $posts = array_slice($posts, $offset, $this->postsPerPage);
        
        return [
            'posts' => $posts,
            'pagination' => $this->buildPaginationData($page, $totalPages)
        ];
    }
    
    public function getPost($slug) {
        return $this->getPostBySlug($slug);
    }
    
    public function getPostBySlug($slug, $categorySlug = null) {
        // Validate slug to prevent path traversal
        if (strpos($slug, '..') !== false || strpos($slug, '/') !== false || strpos($slug, '\\') !== false) {
            return null;
        }
        
        if ($categorySlug) {
            $category = $this->getCategoryBySlug($categorySlug);
            if (!$category) {
                return null;
            }
            $posts_dir = $this->postsDir . '/' . $category['folder'];
        } else {
            $posts_dir = $this->postsDir;
        }
        
        if (!is_dir($posts_dir)) {
            return null;
        }
        
        $filepath = $posts_dir . '/' . $slug . '.md';
        
        if (file_exists($filepath)) {
            $post = $this->parsePost($filepath);
            if ($post) {
                $post['category_slug'] = $categorySlug;
                $post['category'] = $categorySlug ? $this->getCategoryBySlug($categorySlug) : null;
                return $post;
            }
        }
        
        return null;
    }
    
    public function getInclude($filename) {
        $path = 'includes/' . $filename;
        if (file_exists($path)) {
            $content = file_get_contents($path);
            $parsed = $this->parser->parse($content);
            return $parsed['html'];
        }
        return '';
    }
    
    public function parseMarkdown($content) {
        $parsed = $this->parser->parse($content);
        return $parsed['html'];
    }
    
    public function getCategories() {
        $categories = [];
        
        if (!isset($this->config['categories']) || !is_array($this->config['categories'])) {
            return $categories;
        }
        
        foreach ($this->config['categories'] as $key => $category) {
            $folder = $category['folder'] ?? $key;
            $path = $this->postsDir . '/' . $folder;
            
            if (is_dir($path)) {
                $post_count = count(glob($path . '/*.md'));
                if ($post_count > 0) {
                    $categories[$key] = [
                        'blog_name' => $category['blog_name'] ?? ucfirst($key),
                        'header_content' => $category['header_content'] ?? '',
                        'folder' => $folder,
                        'slug' => $key,
                        'count' => $post_count
                    ];
                }
            }
        }
        
        return $categories;
    }
    
    public function getCategoryBySlug($slug) {
        $categories = $this->getCategories();
        return $categories[$slug] ?? null;
    }
    
    /**
     * Create a post array from a markdown file
     */
    private function createPostFromFile($file, $includeExcerpt = false) {
        $content = file_get_contents($this->postsDir . '/' . $file);
        $parsed = $this->parser->parse($content);
        
        $date = $this->extractPostDate($parsed['frontMatter'], $file);
        $slug = $this->generateSlug($file);
        $title = $this->extractPostTitle($parsed['frontMatter'], $file);
        
        $post = [
            'filename' => $file,
            'slug' => $slug,
            'title' => $title,
            'date' => $date,
            'content' => $parsed['html'],
            'frontMatter' => $parsed['frontMatter']
        ];
        
        if ($includeExcerpt) {
            $post['excerpt'] = $this->generateExcerpt($parsed['html']);
        }
        
        return $post;
    }
    
    /**
     * Extract and normalize date from front matter or file modification time
     */
    private function extractPostDate($frontMatter, $filename) {
        $frontMatterDate = isset($frontMatter['date']) ? $frontMatter['date'] : null;
        
        // Handle case where date might be parsed as array
        if (is_array($frontMatterDate)) {
            return isset($frontMatterDate[0]) ? $frontMatterDate[0] : $this->getFileModificationDate($filename);
        }
        
        if (empty($frontMatterDate)) {
            return $this->getFileModificationDate($filename);
        }
        
        return $frontMatterDate;
    }
    
    /**
     * Extract title from front matter or generate from filename
     */
    private function extractPostTitle($frontMatter, $filename) {
        return isset($frontMatter['title']) ? $frontMatter['title'] : $this->getTitleFromFilename($filename);
    }
    
    /**
     * Get file modification date formatted
     */
    private function getFileModificationDate($filename) {
        return date($this->dateFormat, filemtime($this->postsDir . '/' . $filename));
    }
    
    /**
     * Sort posts by date comparison function
     */
    private function sortPostsByDate($a, $b) {
        $dateA = is_array($a['date']) ? (isset($a['date'][0]) ? $a['date'][0] : date($this->dateFormat)) : $a['date'];
        $dateB = is_array($b['date']) ? (isset($b['date'][0]) ? $b['date'][0] : date($this->dateFormat)) : $b['date'];
        return strtotime($dateB) - strtotime($dateA);
    }
    
    /**
     * Build pagination data structure
     */
    private function buildPaginationData($currentPage, $totalPages) {
        return [
            'current' => $currentPage,
            'total' => $totalPages,
            'hasNext' => $currentPage < $totalPages,
            'hasPrev' => $currentPage > 1,
            'next' => $currentPage + 1,
            'prev' => $currentPage - 1
        ];
    }
    
    private function scanAllPosts($posts_dir) {
        $posts = [];
        
        // Get uncategorized posts (files directly in posts_dir)
        if ($this->config['show_uncategorized'] ?? true) {
            $posts = array_merge($posts, $this->scanPostsInFolder($posts_dir, null));
        }
        
        // Get categorized posts that should show in index
        $categories = $this->getCategories();
        foreach ($categories as $slug => $category) {
            // Check if category should be shown in main index
            $categoryConfig = $this->config['categories'][$slug] ?? [];
            $showInIndex = $categoryConfig['index'] ?? true;
            
            if ($showInIndex) {
                $category_dir = $posts_dir . '/' . $category['folder'];
                if (is_dir($category_dir)) {
                    $posts = array_merge($posts, $this->scanPostsInFolder($category_dir, $slug));
                }
            }
        }
        
        return $posts;
    }
    
    private function scanPostsInFolder($dir, $categorySlug = null) {
        $posts = [];
        
        if (!is_dir($dir)) {
            return $posts;
        }
        
        $files = scandir($dir);
        
        foreach ($files as $file) {
            if ($file === '.' || $file === '..' || is_dir($dir . '/' . $file)) continue;
            
            if (pathinfo($file, PATHINFO_EXTENSION) === 'md') {
                $post = $this->parsePost($dir . '/' . $file);
                if ($post) {
                    $post['category_slug'] = $categorySlug;
                    $post['category'] = $categorySlug ? $this->getCategoryBySlug($categorySlug) : null;
                    $posts[] = $post;
                }
            }
        }
        
        return $posts;
    }
    
    private function parsePost($filepath) {
        $content = file_get_contents($filepath);
        $parsed = $this->parser->parse($content);
        
        $filename = basename($filepath);
        $date = $this->extractPostDate($parsed['frontMatter'], $filename);
        $slug = $this->generateSlug($filename);
        $title = $this->extractPostTitle($parsed['frontMatter'], $filename);
        
        $post = [
            'filename' => $filename,
            'slug' => $slug,
            'title' => $title,
            'date' => $date,
            'content' => $parsed['html'],
            'frontMatter' => $parsed['frontMatter'],
            'excerpt' => $this->generateExcerpt($parsed['html'])
        ];
        
        return $post;
    }
    
    private function getMarkdownFiles() {
        if (!is_dir($this->postsDir)) {
            return [];
        }
        
        $files = scandir($this->postsDir);
        $mdFiles = array_filter($files, function($file) {
            return pathinfo($file, PATHINFO_EXTENSION) === 'md';
        });
        
        return array_values($mdFiles);
    }
    
    private function generateSlug($filename) {
        $name = pathinfo($filename, PATHINFO_FILENAME);
        $slug = strtolower($name);
        $slug = preg_replace(self::SLUG_CLEANUP_REGEX, '-', $slug);
        $slug = trim($slug, '-');
        return $slug;
    }
    
    private function getTitleFromFilename($filename) {
        $name = pathinfo($filename, PATHINFO_FILENAME);
        // Remove date prefix if present (YYYY-MM-DD-)
        $name = preg_replace(self::DATE_PREFIX_REGEX, '', $name);
        // Replace dashes and underscores with spaces and capitalize
        $title = str_replace(['-', '_'], ' ', $name);
        return ucwords($title);
    }
    
    private function generateExcerpt($html, $length = null) {
        $length = $length ?? $this->excerptLength;
        $text = strip_tags($html);
        $text = preg_replace(self::WHITESPACE_CLEANUP_REGEX, ' ', $text);
        $text = trim($text);
        
        if (strlen($text) > $length) {
            $text = substr($text, 0, $length);
            $lastSpace = strrpos($text, ' ');
            if ($lastSpace !== false) {
                $text = substr($text, 0, $lastSpace);
            }
            $text .= '...';
        }
        
        return $text;
    }
}

?>