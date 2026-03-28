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
                    <span class="post-author"><?php echo sprintf($config['labels']['author_by'] ?? 'By %s', htmlspecialchars($post['frontMatter']['author'])); ?></span>
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
            <a href="index.php?category=<?php echo urlencode($categorySlug); ?>" class="back-to-home">
                <?php echo sprintf($config['labels']['back_to_category'] ?? '&larr; Back to %s', htmlspecialchars($post['category']['blog_name'])); ?>
            </a>
        <?php else: ?>
            <a href="index.php" class="back-to-home">
                <?php echo $config['labels']['back_to_all'] ?? '&larr; Back to all posts'; ?>
            </a>
        <?php endif; ?>
    </nav>
</main>
