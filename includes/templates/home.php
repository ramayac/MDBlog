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
        // Optional blurb
        $indexBlurbFile = $config['posts_dir'] . '/index.md';
        $indexBlurb = file_exists($indexBlurbFile)
            ? $blog->parseMarkdown(file_get_contents($indexBlurbFile))
            : null;
        if ($indexBlurb): ?>
        <div class="index-content">
            <?php echo $indexBlurb['html'] ?? $indexBlurb; // Handling string or array depending on what parseMarkdown returns ?>
        </div>
    <?php endif; ?>

    <?php if (!empty($allCategories)): ?>
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
