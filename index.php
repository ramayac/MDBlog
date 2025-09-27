<?php

// Start timing
$start_time = microtime(true);

// Load configuration
$config = require 'config.php';

require_once 'includes/Blog.php';

$blog = new Blog();

// Get category and page from URL parameters
$categorySlug = isset($_GET['category']) ? $_GET['category'] : null;
$page = isset($_GET['page']) ? (int)$_GET['page'] : 1;
$page = max(1, $page); // Ensure page is at least 1

// Validate category exists
$currentCategory = null;
if ($categorySlug) {
    $currentCategory = $blog->getCategoryBySlug($categorySlug);
    if (!$currentCategory) {
        $categorySlug = null;
    }
}

// Get posts for current page and category
$data = $blog->getPosts($page, $categorySlug);
$posts = $data['posts'];
$pagination = $data['pagination'];

// Get menu
$menu = $blog->getInclude('menu.md');

?>
<?php
$pageTitle = $currentCategory ? $config['blog_name'] . ' - ' . $currentCategory['blog_name'] : $config['blog_name'];
include 'includes/head.php';
?>
<body>
    <div class="container">
        <?php if ($menu): ?>
        <nav class="site-menu">
            <?php echo $menu; ?>
        </nav>
        <?php endif; ?>
        
        <div class="blog-title">
            <?php if ($currentCategory): ?>
                <h2 class="category-title">
                    <?php echo htmlspecialchars($currentCategory['blog_name']); ?>
                    <?php if (!empty($currentCategory['header_content'])): ?>
                        <br><small><?php echo htmlspecialchars($currentCategory['header_content']); ?></small>
                    <?php endif; ?>
                </h2>
            <?php endif; ?>
        </div>
        
        <?php if ($currentCategory): ?>
            <!-- Category header content is already displayed in the category title above -->
        <?php elseif (!empty($config['header_content'])): ?>
            <header class="site-header">
                <?php echo $blog->parseMarkdown($config['header_content']); ?>
            </header>
        <?php endif; ?>
        
        <main class="main-content">
            <?php if (empty($posts)): ?>
                <div class="no-posts">
                    <h2>No posts found</h2>
                    <p>Add some markdown files to the <code>posts</code> directory to get started.</p>
                </div>
            <?php else: ?>
                <div class="posts-grid">
                    <?php foreach ($posts as $post): ?>
                        <article class="post-preview">
                            <h2 class="post-title">
                                <a href="post.php?slug=<?php echo urlencode($post['slug']); ?><?php echo $categorySlug ? '&category=' . urlencode($categorySlug) : ''; ?>">
                                    <?php echo htmlspecialchars($post['title']); ?>
                                </a>
                            </h2>
                            <div class="post-meta">
                                <?php if ($post['category']): ?>
                                    <span class="post-category">
                                      <?php echo date('F j, Y', strtotime($post['date'])); ?>  in <a href="?category=<?php echo urlencode($post['category_slug']); ?>"><?php echo htmlspecialchars($post['category']['blog_name']); ?></a>
                                    </span>
                                <?php endif; ?>
                            </div>
                            <div class="post-excerpt">
                                <?php echo $post['excerpt']; ?>
                            </div>
                            <a href="post.php?slug=<?php echo urlencode($post['slug']); ?><?php echo $categorySlug ? '&category=' . urlencode($categorySlug) : ''; ?>" class="read-more">
                                Read more →
                            </a>
                        </article>
                    <?php endforeach; ?>
                </div>
                
                <?php if ($pagination['total'] > 1): ?>
                    <nav class="pagination">
                        <?php if ($pagination['hasPrev']): ?>
                            <a href="?<?php echo $categorySlug ? 'category=' . urlencode($categorySlug) . '&' : ''; ?>page=<?php echo $pagination['prev']; ?>" class="pagination-link prev">
                                ← Previous
                            </a>
                        <?php endif; ?>
                        
                        <span class="pagination-info">
                            Page <?php echo $pagination['current']; ?> of <?php echo $pagination['total']; ?>
                        </span>
                        
                        <?php if ($pagination['hasNext']): ?>
                            <a href="?<?php echo $categorySlug ? 'category=' . urlencode($categorySlug) . '&' : ''; ?>page=<?php echo $pagination['next']; ?>" class="pagination-link next">
                                Next →
                            </a>
                        <?php endif; ?>
                    </nav>
                <?php endif; ?>
            <?php endif; ?>
        </main>
        
        <?php if (!empty($config['footer_content'])): ?>
        <footer class="site-footer">
            <?php echo $blog->parseMarkdown($config['footer_content']); ?>
            <?php include 'includes/debug.php'; ?>
        </footer>
        <?php endif; ?>
    </div>
</body>
</html>