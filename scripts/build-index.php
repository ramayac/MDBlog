<?php

/**
 * Build-time post metadata index generator.
 *
 * Scans all Markdown post files and writes a lightweight JSON index to
 * cache/posts.index.json containing only the metadata needed for listing
 * and pagination — no full Markdown bodies are rendered.
 *
 * Usage:   php scripts/build-index.php
 * Make:    make build-index
 */

// Always run from the project root so relative paths in config.php resolve correctly.
chdir(dirname(__DIR__));

require_once 'vendor/autoload.php';
require_once 'includes/MarkdownParser.php';

$config     = require 'config.php';
$parser     = new MarkdownParser();
$postsDir   = $config['posts_dir']        ?? 'posts';
$indexFile  = $config['post_index_file']  ?? 'cache/posts.index.json';
$dateFormat = $config['date_format']      ?? 'Y-m-d';
$categories = $config['categories']       ?? [];

// ---------------------------------------------------------------------------
// Helper: plain-text excerpt from raw Markdown (no Parsedown needed)
// ---------------------------------------------------------------------------
function rawMarkdownExcerpt(string $raw, int $length = 200): string
{
    // Strip fenced code blocks
    $text = preg_replace('/```[\s\S]*?```/', '', $raw);
    // Strip inline code
    $text = preg_replace('/`[^`]+`/', '', $text);
    // Strip headings
    $text = preg_replace('/^#{1,6}\s+/m', '', $text);
    // Strip images (keep nothing)
    $text = preg_replace('/!\[([^\]]*)\]\([^\)]+\)/', '', $text);
    // Strip links — keep label text
    $text = preg_replace('/\[([^\]]+)\]\([^\)]+\)/', '$1', $text);
    // Strip bold / italic markers
    $text = preg_replace('/\*{1,2}([^*]+)\*{1,2}/', '$1', $text);
    $text = preg_replace('/_{1,2}([^_]+)_{1,2}/', '$1', $text);
    // Collapse whitespace
    $text = preg_replace('/\s+/', ' ', $text);
    $text = trim($text);

    if (mb_strlen($text) > $length) {
        $text = mb_substr($text, 0, $length);
        $pos  = mb_strrpos($text, ' ');
        if ($pos !== false) {
            $text = mb_substr($text, 0, $pos);
        }
        $text .= '...';
    }

    return $text;
}

// ---------------------------------------------------------------------------
// Helper: generate a URL slug from a filename (mirrors Blog::generateSlug)
// ---------------------------------------------------------------------------
function slugFromFilename(string $filename): string
{
    $name = pathinfo($filename, PATHINFO_FILENAME);
    $slug = strtolower($name);
    $slug = preg_replace('/[^a-z0-9]+/', '-', $slug);
    return trim($slug, '-');
}

// ---------------------------------------------------------------------------
// Helper: human-readable title from filename (mirrors Blog::getTitleFromFilename)
// ---------------------------------------------------------------------------
function titleFromFilename(string $filename): string
{
    $name = pathinfo($filename, PATHINFO_FILENAME);
    $name = preg_replace('/^\d{4}-\d{2}-\d{2}-/', '', $name);
    return ucwords(str_replace(['-', '_'], ' ', $name));
}

// ---------------------------------------------------------------------------
// Helper: resolve the post date (mirrors Blog::extractPostDate)
// ---------------------------------------------------------------------------
function resolveDate(array $frontMatter, string $filepath, string $dateFormat): string
{
    $d = $frontMatter['date'] ?? null;
    if (is_array($d)) {
        $d = $d[0] ?? null;
    }
    if (empty($d)) {
        return date($dateFormat, filemtime($filepath));
    }
    return (string) $d;
}

// ---------------------------------------------------------------------------
// Scan a folder, extract metadata from each .md file
// ---------------------------------------------------------------------------
function scanFolder(
    string $dir,
    ?string $categorySlug,
    MarkdownParser $parser,
    string $dateFormat
): array {
    $entries = [];

    if (!is_dir($dir)) {
        return $entries;
    }

    $files = scandir($dir);
    if ($files === false) {
        return $entries;
    }

    foreach ($files as $file) {
        if ($file === '.' || $file === '..' || is_dir($dir . '/' . $file)) {
            continue;
        }
        if (pathinfo($file, PATHINFO_EXTENSION) !== 'md') {
            continue;
        }

        $filepath = $dir . '/' . $file;
        $content  = file_get_contents($filepath);
        if ($content === false) {
            fwrite(STDERR, "WARNING: Cannot read {$filepath}, skipping." . PHP_EOL);
            continue;
        }

        $meta        = $parser->parseMetaOnly($content);
        $fm          = $meta['frontMatter'];
        $body        = $meta['body'];

        $slug        = slugFromFilename($file);
        $title       = !empty($fm['title']) ? $fm['title'] : titleFromFilename($file);
        $date        = resolveDate($fm, $filepath, $dateFormat);
        $author      = $fm['author']      ?? '';
        $tags        = $fm['tags']        ?? '';
        $description = $fm['description'] ?? '';
        $excerpt     = !empty($description) ? $description : rawMarkdownExcerpt($body);

        // Source path relative to project root
        $sourcePath = ltrim(
            str_replace(getcwd() . DIRECTORY_SEPARATOR, '', realpath($filepath)),
            DIRECTORY_SEPARATOR
        );

        $entries[] = [
            'slug'          => $slug,
            'title'         => $title,
            'date'          => $date,
            'author'        => $author,
            'tags'          => $tags,
            'description'   => $description,
            'excerpt'       => $excerpt,
            'category_slug' => $categorySlug,
            'source_path'   => $sourcePath,
            'filename'      => $file,
        ];
    }

    return $entries;
}

// ---------------------------------------------------------------------------
// Collect all post entries
// ---------------------------------------------------------------------------
$allEntries = [];

// Uncategorized posts (root posts dir, direct .md files only — not index.md blurb)
if ($config['show_uncategorized'] ?? true) {
    $allEntries = array_merge(
        $allEntries,
        scanFolder($postsDir, null, $parser, $dateFormat)
    );
}

// Categorized posts
foreach ($categories as $slug => $category) {
    $folder = $category['folder'] ?? $slug;
    $allEntries = array_merge(
        $allEntries,
        scanFolder($postsDir . '/' . $folder, $slug, $parser, $dateFormat)
    );
}

// Sort by date descending (newest first) — pre-sorted so runtime needs no sort
usort($allEntries, function (array $a, array $b): int {
    return strtotime($b['date']) <=> strtotime($a['date']);
});

// ---------------------------------------------------------------------------
// Write the index (atomic write to avoid partial reads)
// ---------------------------------------------------------------------------
$indexDir = dirname($indexFile);
if (!is_dir($indexDir)) {
    if (!mkdir($indexDir, 0755, true)) {
        fwrite(STDERR, 'ERROR: Cannot create directory ' . $indexDir . PHP_EOL);
        exit(1);
    }
}

$encoded = json_encode($allEntries, JSON_UNESCAPED_UNICODE | JSON_PRETTY_PRINT);
if ($encoded === false) {
    fwrite(STDERR, 'ERROR: json_encode failed: ' . json_last_error_msg() . PHP_EOL);
    exit(1);
}

$tmpFile = $indexFile . '.tmp.' . getmypid();
if (file_put_contents($tmpFile, $encoded, LOCK_EX) === false) {
    fwrite(STDERR, 'ERROR: Cannot write to ' . $tmpFile . PHP_EOL);
    exit(1);
}
rename($tmpFile, $indexFile);

$count = count($allEntries);
echo "Built post index: {$count} post(s) → {$indexFile}" . PHP_EOL;
