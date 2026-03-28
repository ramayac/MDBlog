<?php

class MarkdownParser {
    private $parsedown;
    
    public function __construct() {
        $this->parsedown = new Parsedown();
        $this->parsedown->setBreaksEnabled(true); // Enable line breaks
        $this->parsedown->setSafeMode(true); // Enable safe mode
        $this->parsedown->setMarkupEscaped(true); // Escape HTML for security
    }
    
    /**
     * Extract front matter and return the raw Markdown body without rendering HTML.
     * Faster than parse() — use this during build-time index generation when only
     * metadata is needed, so Parsedown is never invoked.
     */
    public function parseMetaOnly(string $text): array {
        $text = mb_convert_encoding($text, 'UTF-8', 'UTF-8');
        $text = str_replace("\r\n", "\n", $text);
        $text = str_replace("\r", "\n", $text);

        $frontMatter = [];
        $body        = $text;
        if (preg_match('/^---\n(.*?)\n---\n(.*)$/s', $text, $matches)) {
            $frontMatter = $this->parseFrontMatter($matches[1]);
            $body        = $matches[2];
        }

        return [
            'frontMatter' => $frontMatter,
            'body'        => $body,
        ];
    }

    public function parse($text) {
        // Ensure input is valid UTF-8; strip or replace any invalid byte sequences
        // so that Parsedown and json_encode() downstream never see malformed bytes.
        $text = mb_convert_encoding($text, 'UTF-8', 'UTF-8');

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