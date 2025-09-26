<?php

// Start timing
$start_time = microtime(true);

require_once 'includes/Blog.php';

$blog = new Blog();

// Get post slug from URL parameter
$slug = isset($_GET['slug']) ? $_GET['slug'] : '';

if (empty($slug)) {
    header('Location: index.php');
    exit;
}

// Get post data
$post = $blog->getPost($slug);

if (!$post) {
    http_response_code(404);
    echo "Post not found";
    exit;
}

// Get header and footer
$header = $blog->getInclude('header.md');
$footer = $blog->getInclude('footer.md');

// Check for custom JavaScript files
$jsFiles = [];
if (isset($post['frontMatter']['js'])) {
    $jsFiles = is_array($post['frontMatter']['js']) ? $post['frontMatter']['js'] : [$post['frontMatter']['js']];
}

?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title><?php echo htmlspecialchars($post['title']); ?> - MDBlog</title>
    <link rel="stylesheet" href="assets/css/style.css">
    
    <?php if (isset($post['frontMatter']['description'])): ?>
    <meta name="description" content="<?php echo htmlspecialchars($post['frontMatter']['description']); ?>">
    <?php endif; ?>
</head>
<body>
    <div class="container">
        <?php if ($header): ?>
        <header class="site-header">
            <?php echo $header; ?>
        </header>
        <?php endif; ?>
        
        <main class="main-content">
            <article class="post">
                <header class="post-header">
                    <h1 class="post-title"><?php echo htmlspecialchars($post['title']); ?></h1>
                    <div class="post-meta">
                        <time datetime="<?php echo $post['date']; ?>">
                            <?php echo date('F j, Y', strtotime($post['date'])); ?>
                        </time>
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
                <a href="index.php" class="back-to-home">← Back to all posts</a>
            </nav>
        </main>
        
        <?php if ($footer): ?>
        <footer class="site-footer">
            <?php echo $footer; ?>
        </footer>
        <?php endif; ?>
    </div>
    
    <?php foreach ($jsFiles as $jsFile): ?>
        <?php if (file_exists('assets/js/' . $jsFile)): ?>
            <script src="assets/js/<?php echo htmlspecialchars($jsFile); ?>"></script>
        <?php endif; ?>
    <?php endforeach; ?>
    
    <?php
    // Calculate render time
    $end_time = microtime(true);
    $render_time = round(($end_time - $start_time) * 1000, 2);
    ?>
    <!-- Page rendered in <?php echo $render_time; ?>ms -->
</body>
</html>