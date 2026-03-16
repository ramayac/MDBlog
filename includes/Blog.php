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
    private $cacheEnabled;
    private $cacheTtl;
    private $cacheDir;
    
    public function __construct($configFile = 'config.php') {
        // Load configuration
        $this->config = require $configFile;
        
        // Set properties from config
        $this->postsDir = $this->config['posts_dir'];
        $this->parser = new MarkdownParser();
        $this->postsPerPage = $this->config['posts_per_page'];
        $this->excerptLength = $this->config['excerpt_length'];
        $this->dateFormat = $this->config['date_format'];
        
        // Cache settings
        $this->cacheEnabled = $this->config['cache_enabled'] ?? false;
        $this->cacheTtl = $this->config['cache_ttl'] ?? 604800;
        $this->cacheDir = $this->config['cache_dir'] ?? 'cache';
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

    /**
     * Build the navigation menu HTML from config.
     * Renders custom menu_links first, then any category with 'menu' => true.
     */
    public function getMenu() {
        $links = [];

        // Custom static links
        foreach ($this->config['menu_links'] ?? [] as $link) {
            $label = htmlspecialchars($link['label'] ?? '');
            $url   = htmlspecialchars($link['url']   ?? '#');
            $links[] = '<a href="' . $url . '">' . $label . '</a>';
        }

        // Auto-generated category links
        foreach ($this->config['categories'] ?? [] as $slug => $category) {
            if (!empty($category['menu'])) {
                $label   = htmlspecialchars($category['blog_name'] ?? ucfirst($slug));
                $catSlug = htmlspecialchars(urlencode($slug));
                $links[] = '<a href="index.php?category=' . $catSlug . '">' . $label . '</a>';
            }
        }

        return implode(' | ', $links);
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
        // Try to get from cache first
        $cacheKey = $this->getCacheKey($filepath);
        $cachedPost = $this->getFromCache($cacheKey, $filepath);
        
        if ($cachedPost) {
            return $cachedPost;
        }

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
        
        // Save to cache
        $this->saveToCache($cacheKey, $post);
        
        return $post;
    }

    private function getCacheKey($filepath) {
        return md5($filepath) . '.json';
    }

    private function getFromCache($key, $originalFilepath) {
        if (!$this->cacheEnabled) {
            return null;
        }

        $cacheFile = $this->cacheDir . '/' . $key;
        if (!file_exists($cacheFile)) {
            return null;
        }

        // Check TTL
        if (time() - filemtime($cacheFile) > $this->cacheTtl) {
            return null;
        }

        // Check if original file is newer than cache
        if (file_exists($originalFilepath) && filemtime($originalFilepath) > filemtime($cacheFile)) {
            return null;
        }

        $data = json_decode(file_get_contents($cacheFile), true);
        return $data;
    }

    private function saveToCache($key, $data) {
        if (!$this->cacheEnabled) {
            return;
        }

        if (!is_dir($this->cacheDir)) {
            mkdir($this->cacheDir, 0777, true);
        }

        file_put_contents($this->cacheDir . '/' . $key, json_encode($data));
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

    /**
     * Returns the latest git commit hash, date, and version tag.
     *
     * Priority:
     *   1. version.php in the project root — generated by `make version` before
     *      FTP upload or Docker build. No git required at runtime.
     *   2. Live `git` shell commands — works on a local dev machine that has git.
     *   3. All nulls — version row silently hidden in the footer.
     *
     * Result is cached in a static variable to avoid repeated calls per request.
     */
    public function getVersionInfo() {
        static $info = null;
        if ($info !== null) {
            return $info;
        }

        // 1. Baked-in file takes priority (FTP / Docker deployments)
        $versionFile = __DIR__ . '/../version.php';
        if (file_exists($versionFile)) {
            $baked = include $versionFile;
            if (is_array($baked)) {
                $info = $baked;
                return $info;
            }
        }

        // 2. Fall back to live git (local dev)
        $commit  = null;
        $date    = null;
        $version = null;

        if (function_exists('shell_exec')) {
            $rawCommit = shell_exec('git log -1 --format="%h" 2>/dev/null');
            $rawDate   = shell_exec('git log -1 --format="%ad" --date=short 2>/dev/null');
            $rawTag    = shell_exec('git describe --tags --abbrev=0 2>/dev/null');

            $commit  = ($rawCommit !== null) ? trim($rawCommit) : null;
            $date    = ($rawDate   !== null) ? trim($rawDate)   : null;
            $tag     = ($rawTag    !== null) ? trim($rawTag)    : null;

            // Fall back to commit hash when no annotated tag exists
            $version = ($tag !== '') ? $tag : $commit;

            // Treat empty strings as null
            $commit  = ($commit  !== '') ? $commit  : null;
            $date    = ($date    !== '') ? $date    : null;
            $version = ($version !== '') ? $version : null;
        }

        $info = [
            'commit'  => $commit,
            'date'    => $date,
            'version' => $version,
        ];

        return $info;
    }
}

?>