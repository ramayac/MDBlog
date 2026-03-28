<?php // ── CATEGORY PAGE: post listing + pagination ────────────────── ?>
<h2 class="category-title">
    <?php echo htmlspecialchars($currentCategory['blog_name']); ?>
    <?php if (!empty($currentCategory['header_content'])): ?>
        <br><small><?php echo htmlspecialchars($currentCategory['header_content']); ?></small>
    <?php endif; ?>
</h2>

<main class="main-content">
    <?php if (empty($posts)): ?>
        <div class="no-posts">
            <p><?php echo htmlspecialchars($config['labels']['no_posts_in_category'] ?? 'No posts found in this category.'); ?></p>
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
                        <span class="post-date"><?php echo htmlspecialchars(sprintf($config['labels']['posted_on'] ?? 'on %s', date('F j, Y', strtotime($post['date'])))); ?></span>
                    </div>
                    <?php if (!empty($post['excerpt'])): ?>
                    <div class="post-excerpt">
                        <?php echo $post['excerpt']; ?>
                    </div>
                    <?php endif; ?>
                    <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>&category=<?php echo urlencode($categorySlug); ?>" class="read-more">
                        <?php echo $config['labels']['read_more'] ?? 'Read more &rarr;'; ?>
                    </a>
                </article>
            <?php endforeach; ?>
        </div>

        <?php if ($pagination['total'] > 1): ?>
            <nav class="pagination">
                <?php if ($pagination['hasPrev']): ?>
                    <a href="?category=<?php echo urlencode($categorySlug); ?>&page=<?php echo $pagination['prev']; ?>" class="pagination-link prev">
                        <?php echo $config['labels']['pagination_prev'] ?? '&larr; Newer Posts'; ?>
                    </a>
                <?php endif; ?>
                <span class="pagination-info">
                    <?php echo sprintf($config['labels']['page_indicator'] ?? 'Page %d of %d', $pagination['current'], $pagination['total']); ?>
                </span>
                <?php if ($pagination['hasNext']): ?>
                    <a href="?category=<?php echo urlencode($categorySlug); ?>&page=<?php echo $pagination['next']; ?>" class="pagination-link next">
                        <?php echo $config['labels']['pagination_next'] ?? 'Older Posts &rarr;'; ?>
                    </a>
                <?php endif; ?>
            </nav>
        <?php endif; ?>
    <?php endif; ?>
</main>
