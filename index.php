<?php

// Start timing
$start_time = microtime(true);

// Load configuration
$config = require 'config.php';

require_once 'includes/Blog.php';

$blog = new Blog();

// Get current page from URL parameter
$page = isset($_GET['page']) ? (int)$_GET['page'] : 1;
$page = max(1, $page); // Ensure page is at least 1

// Get posts for current page
$data = $blog->getPosts($page);
$posts = $data['posts'];
$pagination = $data['pagination'];

// Get menu
$menu = $blog->getInclude('menu.md');

?>
<?php
$pageTitle = $config['blog_name'];
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
            <h1><?php echo htmlspecialchars($config['blog_name']); ?></h1>
        </div>
        
        <?php if (!empty($config['header_content'])): ?>
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
                                <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>">
                                    <?php echo htmlspecialchars($post['title']); ?>
                                </a>
                            </h2>
                            <div class="post-meta">
                                <time datetime="<?php echo $post['date']; ?>">
                                    <?php echo date('F j, Y', strtotime($post['date'])); ?>
                                </time>
                            </div>
                            <div class="post-excerpt">
                                <?php echo $post['excerpt']; ?>
                            </div>
                            <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>" class="read-more">
                                Read more →
                            </a>
                        </article>
                    <?php endforeach; ?>
                </div>
                
                <?php if ($pagination['total'] > 1): ?>
                    <nav class="pagination">
                        <?php if ($pagination['hasPrev']): ?>
                            <a href="?page=<?php echo $pagination['prev']; ?>" class="pagination-link prev">
                                ← Previous
                            </a>
                        <?php endif; ?>
                        
                        <span class="pagination-info">
                            Page <?php echo $pagination['current']; ?> of <?php echo $pagination['total']; ?>
                        </span>
                        
                        <?php if ($pagination['hasNext']): ?>
                            <a href="?page=<?php echo $pagination['next']; ?>" class="pagination-link next">
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