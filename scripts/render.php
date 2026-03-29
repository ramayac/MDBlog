<?php
// scripts/render.php

chdir(dirname(__DIR__));

if (!file_exists('vendor/autoload.php')) {
    echo "Run 'composer install' first.\n";
    exit(1);
}

require 'vendor/autoload.php';
$config = require 'config.php';
require_once 'includes/Blog.php';
require_once 'includes/View.php';

$indexFile = $config['post_index_file'] ?? 'posts/posts.index.json';
if (!file_exists($indexFile)) {
    echo "Post index not found. Building it temporarily...\n";
    exec('php scripts/build-index.php');
}

$allEntries = json_decode(file_get_contents($indexFile), true);
if (empty($allEntries)) {
    echo "No posts found in index.\n";
    exit(1);
}

$args = array_slice($argv, 1);
if (empty($args)) {
    echo "Usage:\n";
    echo "  make render random\n";
    echo "  make render [category] random\n";
    echo "  make render [filename.md]\n";
    exit(1);
}

$targetEntry = null;

if (count($args) === 1 && $args[0] === 'random') {
    $targetEntry = $allEntries[array_rand($allEntries)];
} elseif (count($args) === 2 && $args[1] === 'random') {
    $catSlug = $args[0];
    $catEntries = array_filter($allEntries, function ($e) use ($catSlug) {
        return ($e['category_slug'] ?? '') === $catSlug;
    });
    if (empty($catEntries)) {
        echo "No posts found for category: $catSlug\n";
        exit(1);
    }
    $targetEntry = $catEntries[array_rand($catEntries)];
} elseif (count($args) >= 1) {
    $search = $args[0];
    $searchSlug = preg_replace('/\.md$/', '', $search);
    $searchSlug = preg_replace('/^\d{4}-\d{2}-\d{2}-/', '', $searchSlug); // trim date if present
    
    foreach ($allEntries as $entry) {
        if ($entry['filename'] === $search || $entry['slug'] === $searchSlug || strpos($entry['filename'], $search) !== false) {
            $targetEntry = $entry;
            break;
        }
    }
}

if (!$targetEntry) {
    echo "Could not find a matching post.\n";
    exit(1);
}

$slug = $targetEntry['slug'];
$catSlug = $targetEntry['category_slug'] ?? null;

$details = "Target Post details:\n" .
"  Title:     " . $targetEntry['title'] . "\n" .
"  Date:      " . $targetEntry['date'] . "\n" .
"  Slug:      " . $targetEntry['slug'] . "\n" .
"  Category:  " . ($catSlug ?: '(uncategorized)') . "\n" .
"  File:      " . $targetEntry['source_path'] . "\n\n";

$blog = new Blog();
$post = $blog->getPostBySlug($slug, $catSlug);
$menu = $blog->getMenu();

if (!$post) {
    echo "ERROR: Blog::getPostBySlug() returned null. The file might not exist or the config is misconfigured.\n";
    exit(1);
}

$jsFiles = [];
if (isset($post['frontMatter']['js'])) {
    $jsFiles = is_array($post['frontMatter']['js']) ? $post['frontMatter']['js'] : [$post['frontMatter']['js']];
}

ob_start();
View::render('post.php', [
    'config' => $config,
    'blog' => $blog,
    'menu' => $menu,
    'post' => $post,
    'categorySlug' => $catSlug,
    'jsFiles' => $jsFiles,
    'pageTitle' => htmlspecialchars($post['title']) . ' - ' . $config['blog_name'],
    'pageDescription' => $post['frontMatter']['description'] ?? '',
    'ogType' => 'article',
    'start_time' => microtime(true)
]);
$html = ob_get_clean();

// Add base tag to fix local paths (e.g. assets, images) in the rendered file
$html = str_replace('<head>', "<head>\n    <base href=\"../\">", $html);

$outDir = __DIR__ . '/../render';
if (!is_dir($outDir)) {
    if (!mkdir($outDir, 0755, true)) {
        echo "ERROR: Failed to create output directory.\n";
        exit(1);
    }
}

$outFile = $outDir . '/' . $slug . '.html';
if (file_put_contents($outFile, $html) === false) {
    echo "ERROR: Failed to write output file: $outFile\n";
    exit(1);
}

echo $details;
echo "SUCCESS! Wrote rendered HTML file to: render/" . $slug . ".html\n";
