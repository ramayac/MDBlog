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
    
    public function getPosts($page = 1) {
        $files = $this->getMarkdownFiles();
        $posts = [];
        
        foreach ($files as $file) {
            $post = $this->createPostFromFile($file, true);
            if ($post) {
                $posts[] = $post;
            }
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
        // Validate slug to prevent path traversal
        if (strpos($slug, '..') !== false || strpos($slug, '/') !== false || strpos($slug, '\\') !== false) {
            return null;
        }
        
        $files = $this->getMarkdownFiles();
        
        foreach ($files as $file) {
            if ($this->generateSlug($file) === $slug) {
                return $this->createPostFromFile($file, false);
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