<?php

// Download Parsedown from https://parsedown.org or via Composer
// For this example, we'll include it directly
require_once 'Parsedown.php';

class MarkdownParser {
    private $parsedown;
    
    public function __construct() {
        $this->parsedown = new Parsedown();
        $this->parsedown->setBreaksEnabled(true); // Enable line breaks
        $this->parsedown->setMarkupEscaped(false); // Allow HTML in markdown
    }
    
    public function parse($text) {
        // Remove Windows line endings
        $text = str_replace("\r\n", "\n", $text);
        $text = str_replace("\r", "\n", $text);
        
        // Extract and parse front matter
        $frontMatter = [];
        if (preg_match('/^---\n(.*?)\n---\n(.*)$/s', $text, $matches)) {
            $frontMatter = $this->parseFrontMatter($matches[1]);
            $text = $matches[2];
        }
        
        // Parse markdown content using Parsedown
        $html = $this->parsedown->text($text);
        
        return [
            'frontMatter' => $frontMatter,
            'html' => $html
        ];
    }
    
    private function parseFrontMatter($frontMatter) {
        $data = [];
        $lines = explode("\n", $frontMatter);
        
        foreach ($lines as $line) {
            $line = trim($line);
            if (empty($line)) continue;
            
            if (strpos($line, ':') !== false) {
                list($key, $value) = array_map('trim', explode(':', $line, 2));
                
                // Handle arrays (values starting with - or empty for array continuation)
                if (empty($value) || strpos($value, '- ') === 0 || (isset($data[$key]) && is_array($data[$key]))) {
                    if (!isset($data[$key])) {
                        $data[$key] = [];
                    }
                    if (strpos($value, '- ') === 0) {
                        $data[$key][] = trim(substr($value, 2));
                    }
                } else {
                    $data[$key] = $value;
                }
            } elseif (strpos($line, '- ') === 0 && !empty($data)) {
                // Handle array continuation
                $lastKey = array_key_last($data);
                if (!is_array($data[$lastKey])) {
                    $data[$lastKey] = [$data[$lastKey]];
                }
                $data[$lastKey][] = trim(substr($line, 2));
            }
        }
        
        return $data;
    }

}
?>