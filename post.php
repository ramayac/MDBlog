<?php

if (!isset($_ENV['AWS_LAMBDA_FUNCTION_NAME']) &&
    substr_count($_SERVER['HTTP_ACCEPT_ENCODING'] ?? '', 'gzip')) {
    ob_start('ob_gzhandler');
}

require __DIR__ . '/vendor/autoload.php';

$start_time = microtime(true);
$config = require 'config.php';
require_once 'includes/Blog.php';
require_once 'includes/View.php';

$blog = new Blog();

$slug = isset($_GET['slug']) ? $_GET['slug'] : '';
$categorySlug = isset($_GET['category']) ? $_GET['category'] : null;

if (empty($slug) || !preg_match('/^[a-zA-Z0-9_-]+$/', $slug) || strlen($slug) > 200) {
    header('Location: index.php');
    exit;
}

$post = $blog->getPostBySlug($slug, $categorySlug);
$menu = $blog->getMenu();

if (!$post) {
    http_response_code(404);
    View::render('404.php', [
        'config' => $config,
        'blog' => $blog,
        'menu' => $menu,
        'pageTitle' => isset($config['labels']['not_found_title']) ? $config['labels']['not_found_title'] . ' â€d ' . $config['blog_name'] : '404 Not Found â€” ' . $config['blog_name'],
        'start_time' => $start_time
    ]);
    exit;
}

$jsFiles = [];
if (isset($post['frontMatter']['js'])) {
    $jsFiles = is_array($post['frontMatter']['js']) ? $post['frontMatter']['js'] : [$post['frontMatter']['js']];
}

$postMtime    = strtotime($post['date']) ?: time();
$lastModified = gmdate('D, d M Y H:i:s', $postMtime) . ' GMT';
$etag         = '"' . md5($post['slug'] . $post['date']) . '"';
header('Last-Modified: ' . $lastModified);
header('ETag: ' . $etag);
header('Cache-Control: public, max-age=3600');

if (
    (isset($_SERVER['HTTP_IF_NONE_MATCH']) && trim($_SERVER['HTTP_IF_NONE_MATCH']) === $etag) ||
    (isset($_SERVER['HTTP_IF_MODIFIED_SINCE']) && strtotime($_SERVER['HTTP_IF_MODIFIED_SINCE']) >= $postMtime)
) {
    http_response_code(304);
    exit;
}

View::render('post.php', [
    'config' => $config,
    'blog' => $blog,
    'menu' => $menu,
    'post' => $post,
    'categorySlug' => $categorySlug,
    'jsFiles' => $jsFiles,
    'pageTitle' => htmlspecialchars($post['title']) . ' - ' . $config['blog_name'],
    'pageDescription' => $post['frontMatter']['description'] ?? '',
    'ogType' => 'article',
    'start_time' => $start_time
]);
