<?php

require_once 'includes/MarkdownParser.php';

class Blog {
    private $postsDir;
    private $parser;
    private $postsPerPage;
    
    public function __construct($postsDir = 'posts', $postsPerPage = 25) {
        $this->postsDir = $postsDir;
        $this->parser = new MarkdownParser();
        $this->postsPerPage = $postsPerPage;
    }
    
    public function getPosts($page = 1) {
        $files = $this->getMarkdownFiles();
        $posts = [];
        
        foreach ($files as $file) {
            $content = file_get_contents($this->postsDir . '/' . $file);
            $parsed = $this->parser->parse($content);
            
            $frontMatterDate = isset($parsed['frontMatter']['date']) ? $parsed['frontMatter']['date'] : null;
            $date = $frontMatterDate;
            
            // Handle case where date might be parsed as array
            if (is_array($frontMatterDate)) {
                $date = isset($frontMatterDate[0]) ? $frontMatterDate[0] : date('Y-m-d', filemtime($this->postsDir . '/' . $file));
            } elseif (empty($frontMatterDate)) {
                $date = date('Y-m-d', filemtime($this->postsDir . '/' . $file));
            }
            
            $post = [
                'filename' => $file,
                'slug' => $this->generateSlug($file),
                'title' => isset($parsed['frontMatter']['title']) ? $parsed['frontMatter']['title'] : $this->getTitleFromFilename($file),
                'date' => $date,
                'excerpt' => $this->generateExcerpt($parsed['html']),
                'content' => $parsed['html'],
                'frontMatter' => $parsed['frontMatter']
            ];
            
            $posts[] = $post;
        }
        
        // Sort posts by date (newest first)
        usort($posts, function($a, $b) {
            $dateA = is_array($a['date']) ? (isset($a['date'][0]) ? $a['date'][0] : date('Y-m-d')) : $a['date'];
            $dateB = is_array($b['date']) ? (isset($b['date'][0]) ? $b['date'][0] : date('Y-m-d')) : $b['date'];
            return strtotime($dateB) - strtotime($dateA);
        });
        
        // Calculate pagination
        $totalPosts = count($posts);
        $totalPages = ceil($totalPosts / $this->postsPerPage);
        $offset = ($page - 1) * $this->postsPerPage;
        $posts = array_slice($posts, $offset, $this->postsPerPage);
        
        return [
            'posts' => $posts,
            'pagination' => [
                'current' => $page,
                'total' => $totalPages,
                'hasNext' => $page < $totalPages,
                'hasPrev' => $page > 1,
                'next' => $page + 1,
                'prev' => $page - 1
            ]
        ];
    }
    
    public function getPost($slug) {
        $files = $this->getMarkdownFiles();
        
        foreach ($files as $file) {
            if ($this->generateSlug($file) === $slug) {
                $content = file_get_contents($this->postsDir . '/' . $file);
                $parsed = $this->parser->parse($content);
                
                $frontMatterDate = isset($parsed['frontMatter']['date']) ? $parsed['frontMatter']['date'] : null;
                $date = $frontMatterDate;
                
                // Handle case where date might be parsed as array
                if (is_array($frontMatterDate)) {
                    $date = isset($frontMatterDate[0]) ? $frontMatterDate[0] : date('Y-m-d', filemtime($this->postsDir . '/' . $file));
                } elseif (empty($frontMatterDate)) {
                    $date = date('Y-m-d', filemtime($this->postsDir . '/' . $file));
                }
                
                return [
                    'filename' => $file,
                    'slug' => $slug,
                    'title' => isset($parsed['frontMatter']['title']) ? $parsed['frontMatter']['title'] : $this->getTitleFromFilename($file),
                    'date' => $date,
                    'content' => $parsed['html'],
                    'frontMatter' => $parsed['frontMatter']
                ];
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
        $slug = preg_replace('/[^a-z0-9]+/', '-', $slug);
        $slug = trim($slug, '-');
        return $slug;
    }
    
    private function getTitleFromFilename($filename) {
        $name = pathinfo($filename, PATHINFO_FILENAME);
        // Remove date prefix if present (YYYY-MM-DD-)
        $name = preg_replace('/^\d{4}-\d{2}-\d{2}-/', '', $name);
        // Replace dashes and underscores with spaces and capitalize
        $title = str_replace(['-', '_'], ' ', $name);
        return ucwords($title);
    }
    
    private function generateExcerpt($html, $length = 150) {
        $text = strip_tags($html);
        $text = preg_replace('/\s+/', ' ', $text);
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