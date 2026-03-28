<?php

// ── Request router ──────────────────────────────────────────────────────────
$requestPath = parse_url($_SERVER['REQUEST_URI'] ?? '', PHP_URL_PATH);

if ($requestPath && $requestPath === '/post.php') {
    require __DIR__ . '/post.php';
    exit;
}

// ── Static-asset handler ────────────────────────────────────────────────────
if ($requestPath && preg_match('#^/assets/(.+)$#', $requestPath, $m)) {
    $assetFile = __DIR__ . '/assets/' . $m[1];
    $realBase  = realpath(__DIR__ . '/assets');
    $realFile  = realpath($assetFile);

    if ($realBase && $realFile && str_starts_with($realFile, $realBase . DIRECTORY_SEPARATOR) && is_file($realFile)) {
        $ext = strtolower(pathinfo($realFile, PATHINFO_EXTENSION));
        $mimeMap = [
            'css'   => 'text/css',
            'js'    => 'application/javascript',
            'png'   => 'image/png',
            'jpg'   => 'image/jpeg',
            'jpeg'  => 'image/jpeg',
            'gif'   => 'image/gif',
            'svg'   => 'image/svg+xml',
            'ico'   => 'image/x-icon',
            'woff'  => 'font/woff',
            'woff2' => 'font/woff2',
            'ttf'   => 'font/ttf',
        ];
        $contentType = $mimeMap[$ext] ?? 'application/octet-stream';

        header('Content-Type: ' . $contentType);
        header('Cache-Control: public, max-age=86400');
        readfile($realFile);
        exit;
    }
}

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

$categorySlug = isset($_GET['category']) ? $_GET['category'] : null;
$page = isset($_GET['page']) ? (int)$_GET['page'] : 1;
$page = max(1, $page);

$currentCategory = null;
if ($categorySlug) {
    $currentCategory = $blog->getCategoryBySlug($categorySlug);
    if (!$currentCategory) {
        $categorySlug = null;
    }
}

$menu = $blog->getMenu();
$pageTitle = $currentCategory
    ? $config['blog_name'] . ' - ' . $currentCategory['blog_name']
    : $config['blog_name'];

if ($currentCategory) {
    $data = $blog->getPosts($page, $categorySlug);
    View::render('category.php', [
        'config' => $config,
        'blog' => $blog,
        'menu' => $menu,
        'pageTitle' => $pageTitle,
        'currentCategory' => $currentCategory,
        'categorySlug' => $categorySlug,
        'posts' => $data['posts'],
        'pagination' => $data['pagination'],
        'start_time' => $start_time
    ]);
} else {
    $allCategories = $blog->getCategories();
    View::render('home.php', [
        'config' => $config,
        'blog' => $blog,
        'menu' => $menu,
        'pageTitle' => $pageTitle,
        'allCategories' => $allCategories,
        'start_time' => $start_time
    ]);
}
