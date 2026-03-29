<article class="post-preview">
    <?php
        $postCategorySlug = $post['category_slug'] ?? $categorySlug ?? '';
        $postUrl = 'post.php?slug=' . urlencode($post['slug']) . ($postCategorySlug ? '&category=' . urlencode($postCategorySlug) : '');
    ?>
    <h2 class="post-title">
        <a href="<?php echo htmlspecialchars($postUrl); ?>">
            <?php echo htmlspecialchars($post['title']); ?>
        </a>
    </h2>
    <div class="post-meta">
        <span class="post-date"><?php echo htmlspecialchars(sprintf($config['labels']['posted_on'] ?? 'on %s', date('F j, Y', strtotime($post['date'])))); ?></span>
    </div>
    <?php if (!empty($post['excerpt'])): ?>
    <div class="post-excerpt">
        <?php echo $post['excerpt']; ?>
    </div>
    <?php endif; ?>
    <a href="<?php echo htmlspecialchars($postUrl); ?>" class="read-more">
        <?php echo $config['labels']['read_more'] ?? 'Read more &rarr;'; ?>
    </a>
</article>
