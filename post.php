<?php

// Enable gzip compression if supported by the client
// Skip on AWS Lambda — Bref JSON-encodes the response body and cannot handle binary gzip output
if (!isset($_ENV['AWS_LAMBDA_FUNCTION_NAME']) &&
    substr_count($_SERVER['HTTP_ACCEPT_ENCODING'] ?? '', 'gzip')) {
    ob_start('ob_gzhandler');
}

// Composer autoloader
require __DIR__ . '/vendor/autoload.php';

// Start timing
$start_time = microtime(true);

// Load configuration
$config = require 'config.php';

require_once 'includes/Blog.php';

$blog = new Blog();

// Get post slug and category from URL parameters
$slug = isset($_GET['slug']) ? $_GET['slug'] : '';
$categorySlug = isset($_GET['category']) ? $_GET['category'] : null;

if (empty($slug)) {
    header('Location: index.php');
    exit;
}

// Validate slug format - only allow alphanumeric, hyphens, underscores
if (!preg_match('/^[a-zA-Z0-9_-]+$/', $slug)) {
    header('Location: index.php');
    exit;
}

// Additional length check
if (strlen($slug) > 200) {
    header('Location: index.php');
    exit;
}

// Get post data
$post = $blog->getPostBySlug($slug, $categorySlug);

if (!$post) {
    http_response_code(404);
    $menu = $blog->getMenu();
    $pageTitle = '404 Not Found — ' . $config['blog_name'];
    include 'includes/head.php';
    echo '<body><div class="container">';
    if ($menu) {
        echo '<nav class="site-menu">' . $menu . '</nav>';
    }
    echo '<main class="main-content"><div class="no-posts">';
    echo '<h2>404 &mdash; Post Not Found</h2>';
    echo '<p>The post you&rsquo;re looking for doesn&rsquo;t exist.</p>';
    echo '<a href="index.php" class="back-to-home">&larr; Back to home</a>';
    echo '</div></main></div></body></html>';
    exit;
}

// Get Menu
$menu = $blog->getMenu();

// Check for custom JavaScript files
$jsFiles = [];
if (isset($post['frontMatter']['js'])) {
    $jsFiles = is_array($post['frontMatter']['js']) ? $post['frontMatter']['js'] : [$post['frontMatter']['js']];
}

// HTTP cache headers for post pages
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
?>
<?php
$pageTitle       = htmlspecialchars($post['title']) . ' - ' . $config['blog_name'];
$pageDescription = isset($post['frontMatter']['description']) ? $post['frontMatter']['description'] : '';
$ogType          = 'article';
include 'includes/head.php';
?>
<body>
    <div class="container">
        <?php if ($menu): ?>
        <nav class="site-menu">
            <?php echo $menu; ?>
        </nav>
        <?php endif; ?>
       
        <main class="main-content">
            <article class="post">
                <header class="post-header">
                    <h1 class="post-title"><?php echo htmlspecialchars($post['title']); ?></h1>
                    <div class="post-meta">
                        <?php if ($post['category']): ?>
                            <span class="post-category">
                                <?php echo date('F j, Y', strtotime($post['date'])); ?> in <a href="index.php?category=<?php echo urlencode($post['category_slug']); ?>"><?php echo htmlspecialchars($post['category']['blog_name']); ?></a>
                            </span>
                        <?php endif; ?>
                        <?php if (isset($post['frontMatter']['author'])): ?>
                            <span class="post-author">by <?php echo htmlspecialchars($post['frontMatter']['author']); ?></span>
                        <?php endif; ?>
                    </div>
                </header>
                
                <div class="post-content">
                    <?php echo $post['content']; ?>
                </div>
                
                <?php if (isset($post['frontMatter']['tags'])): ?>
                    <div class="post-tags">
                        <strong>Tags:</strong>
                        <?php 
                        $tags = is_array($post['frontMatter']['tags']) ? $post['frontMatter']['tags'] : explode(',', $post['frontMatter']['tags']);
                        foreach ($tags as $tag): ?>
                            <span class="tag"><?php echo htmlspecialchars(trim($tag)); ?></span>
                        <?php endforeach; ?>
                    </div>
                <?php endif; ?>
            </article>
            
            <nav class="post-navigation">
                <?php if ($categorySlug): ?>
                    <a href="index.php?category=<?php echo urlencode($categorySlug); ?>" class="back-to-home">← Back to <?php echo htmlspecialchars($post['category']['blog_name']); ?></a>
                <?php else: ?>
                    <a href="index.php" class="back-to-home">← Back to all posts</a>
                <?php endif; ?>
            </nav>
        </main>
        
        <?php
            $footerVersionInfo = $blog->getVersionInfo();
        ?>
        <?php if (!empty($config['footer_content']) || !empty($footerVersionInfo['commit'])): ?>
        <footer class="site-footer">
            <?php if (!empty($config['footer_content'])): ?>
                <?php echo $blog->parseMarkdown($config['footer_content']); ?>
            <?php endif; ?>
            <?php if (!empty($footerVersionInfo['commit'])): ?>
            <small class="site-version">
                <?php if (!empty($footerVersionInfo['version'])): ?><?php echo htmlspecialchars($footerVersionInfo['version']); ?> &middot; <?php endif; ?>
                <?php echo htmlspecialchars($footerVersionInfo['commit']); ?>
                <?php if (!empty($footerVersionInfo['date'])): ?>&middot; <?php echo htmlspecialchars($footerVersionInfo['date']); ?><?php endif; ?>
            </small>
            <?php endif; ?>
        </footer>
        <?php endif; ?>
    </div>
    
    <?php foreach ($jsFiles as $jsFile): ?>
        <?php if (file_exists('assets/js/' . $jsFile)): ?>
            <script src="assets/js/<?php echo htmlspecialchars($jsFile); ?>"></script>
        <?php endif; ?>
        <?php include 'includes/debug.php'; ?>
    <?php endforeach; ?>
    
</body>
</html>