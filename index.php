<?php

// Start timing
$start_time = microtime(true);

require_once 'includes/Blog.php';

$blog = new Blog();

// Get current page from URL parameter
$page = isset($_GET['page']) ? (int)$_GET['page'] : 1;
$page = max(1, $page); // Ensure page is at least 1

// Get posts for current page
$data = $blog->getPosts($page);
$posts = $data['posts'];
$pagination = $data['pagination'];

// Get header and footer
$header = $blog->getInclude('header.md');
$menu = $blog->getInclude('menu.md');
$footer = $blog->getInclude('footer.md');

?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MDBlog</title>
    <link rel="stylesheet" href="assets/css/style.css">
</head>
<body>
    <div class="container">
        <?php if ($menu): ?>
        <nav class="site-menu">
            <?php echo $menu; ?>
        </nav>
        <?php endif; ?>
        <?php if ($header): ?>
        <header class="site-header">
            <?php echo $header; ?>
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
        
        <?php if ($footer): ?>
        <footer class="site-footer">
            <?php echo $footer; ?>
        </footer>
        <?php endif; ?>
    </div>
    
    <?php
    // Calculate render time
    $end_time = microtime(true);
    $render_time = round(($end_time - $start_time) * 1000, 2);
    ?>
    <!-- Page rendered in <?php echo $render_time; ?>ms -->
</body>
</html>