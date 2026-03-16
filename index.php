<?php

// Enable gzip compression if supported by the client
if (substr_count($_SERVER['HTTP_ACCEPT_ENCODING'] ?? '', 'gzip')) {
    ob_start('ob_gzhandler');
}

// Start timing
$start_time = microtime(true);

// Load configuration
$config = require 'config.php';

require_once 'includes/Blog.php';

$blog = new Blog();

// Get category and page from URL parameters
$categorySlug = isset($_GET['category']) ? $_GET['category'] : null;
$page = isset($_GET['page']) ? (int)$_GET['page'] : 1;
$page = max(1, $page);

// Validate category exists
$currentCategory = null;
if ($categorySlug) {
    $currentCategory = $blog->getCategoryBySlug($categorySlug);
    if (!$currentCategory) {
        $categorySlug = null;
    }
}

// Get menu
$menu = $blog->getMenu();

?>
<?php
$pageTitle = $currentCategory
    ? $config['blog_name'] . ' - ' . $currentCategory['blog_name']
    : $config['blog_name'];
include 'includes/head.php';
?>
<body>
    <div class="container">
        <?php if ($menu): ?>
        <nav class="site-menu">
            <?php echo $menu; ?>
        </nav>
        <?php endif; ?>

        <?php if ($currentCategory): ?>
            <?php
                // ── CATEGORY PAGE: post listing + pagination ──────────────────
                $data       = $blog->getPosts($page, $categorySlug);
                $posts      = $data['posts'];
                $pagination = $data['pagination'];
            ?>
            <h2 class="category-title">
                <?php echo htmlspecialchars($currentCategory['blog_name']); ?>
                <?php if (!empty($currentCategory['header_content'])): ?>
                    <br><small><?php echo htmlspecialchars($currentCategory['header_content']); ?></small>
                <?php endif; ?>
            </h2>

            <main class="main-content">
                <?php if (empty($posts)): ?>
                    <div class="no-posts">
                        <p>No posts found in this category yet.</p>
                    </div>
                <?php else: ?>
                    <div class="posts-grid">
                        <?php foreach ($posts as $post): ?>
                            <article class="post-preview">
                                <h2 class="post-title">
                                    <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>&category=<?php echo urlencode($categorySlug); ?>">
                                        <?php echo htmlspecialchars($post['title']); ?>
                                    </a>
                                </h2>
                                <div class="post-meta">
                                    <span class="post-date"><?php echo date('F j, Y', strtotime($post['date'])); ?></span>
                                </div>
                                <div class="post-excerpt">
                                    <?php echo $post['excerpt']; ?>
                                </div>
                                <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>&category=<?php echo urlencode($categorySlug); ?>" class="read-more">
                                    Read more &rarr;
                                </a>
                            </article>
                        <?php endforeach; ?>
                    </div>

                    <?php if ($pagination['total'] > 1): ?>
                        <nav class="pagination">
                            <?php if ($pagination['hasPrev']): ?>
                                <a href="?category=<?php echo urlencode($categorySlug); ?>&page=<?php echo $pagination['prev']; ?>" class="pagination-link prev">
                                    &larr; Previous
                                </a>
                            <?php endif; ?>
                            <span class="pagination-info">
                                Page <?php echo $pagination['current']; ?> of <?php echo $pagination['total']; ?>
                            </span>
                            <?php if ($pagination['hasNext']): ?>
                                <a href="?category=<?php echo urlencode($categorySlug); ?>&page=<?php echo $pagination['next']; ?>" class="pagination-link next">
                                    Next &rarr;
                                </a>
                            <?php endif; ?>
                        </nav>
                    <?php endif; ?>
                <?php endif; ?>
            </main>

        <?php else: ?>
            <?php // ── LANDING PAGE ──────────────────────────────────────────── ?>
            <header class="blog-title">
                <h2>
                    <?php echo htmlspecialchars($config['blog_name']); ?>
                    <?php if (!empty($config['header_content'])): ?>
                        <br><small><?php echo htmlspecialchars($config['header_content']); ?></small>
                    <?php endif; ?>
                </h2>
            </header>

            <main class="main-content">
                <?php
                    // Optional blurb: create posts/index.md to add intro text
                    $indexBlurbFile = $config['posts_dir'] . '/index.md';
                    $indexBlurb = file_exists($indexBlurbFile)
                        ? $blog->parseMarkdown(file_get_contents($indexBlurbFile))
                        : null;
                    if ($indexBlurb): ?>
                    <div class="index-content">
                        <?php echo $indexBlurb; ?>
                    </div>
                <?php endif; ?>

                <?php
                    $allCategories = $blog->getCategories();
                    if (!empty($allCategories)):
                ?>
                <div class="category-cards">
                    <?php foreach ($allCategories as $slug => $cat): ?>
                        <a href="?category=<?php echo urlencode($slug); ?>" class="category-card">
                            <h3><?php echo htmlspecialchars($cat['blog_name']); ?></h3>
                            <?php if (!empty($cat['header_content'])): ?>
                                <p><?php echo htmlspecialchars($cat['header_content']); ?></p>
                            <?php endif; ?>
                            <span class="post-count"><?php echo $cat['count']; ?> posts</span>
                        </a>
                    <?php endforeach; ?>
                </div>
                <?php endif; ?>
            </main>
        <?php endif; ?>

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
            <?php include 'includes/debug.php'; ?>
        </footer>
        <?php endif; ?>
    </div>
</body>
</html>
