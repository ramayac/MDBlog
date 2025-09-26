<?php

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
    echo "Post not found";
    exit;
}

// Get Menu
$menu = $blog->getInclude('menu.md');

// Check for custom JavaScript files
$jsFiles = [];
if (isset($post['frontMatter']['js'])) {
    $jsFiles = is_array($post['frontMatter']['js']) ? $post['frontMatter']['js'] : [$post['frontMatter']['js']];
}

?>
<?php
$pageTitle = htmlspecialchars($post['title']) . ' - ' . $config['blog_name'];
$pageDescription = isset($post['frontMatter']['description']) ? $post['frontMatter']['description'] : '';
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
                        <time datetime="<?php echo $post['date']; ?>">
                            <?php echo date('F j, Y', strtotime($post['date'])); ?>
                        </time>
                        <?php if ($post['category']): ?>
                            <span class="post-category">
                                in <a href="index.php?category=<?php echo urlencode($post['category_slug']); ?>"><?php echo htmlspecialchars($post['category']['blog_name']); ?></a>
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
        
        <?php if (!empty($config['footer_content'])): ?>
        <footer class="site-footer">
            <?php echo $blog->parseMarkdown($config['footer_content']); ?>
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