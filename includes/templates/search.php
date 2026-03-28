<?php // ── SEARCH PAGE: post listing ────────────────── ?>
<div class="search-header" style="text-align: center; margin-bottom: 2em; margin-top: 1em;">
    <h2 class="category-title" style="margin-bottom: 0.5em;">Search</h2>
    <form action="index.php" method="GET" class="standalone-search-form" style="display: flex; justify-content: center; gap: 10px; max-width: 500px; margin: 0 auto;">
        <input type="text" name="q" value="<?php echo htmlspecialchars($query ?? ''); ?>" placeholder="What are you looking for?" required autofocus style="flex: 1; padding: 10px; font-size: 1.1em; border: 1px solid #ccc; border-radius: 6px;">
        <button type="submit" style="padding: 10px 20px; font-size: 1.1em; border: 1px solid #ccc; border-radius: 6px; background: #f0f0f0; cursor: pointer;">🔍 Search</button>
    </form>
    <?php if ($query !== ''): ?>
    <p style="margin-top: 1em; color: #666;">Showing results for "<strong><?php echo htmlspecialchars($query); ?></strong>"</p>
    <?php endif; ?>
</div>

<main class="main-content">
    <?php if ($query === ''): ?>
        <div class="no-posts" style="text-align: center;">
            <p>Enter a keyword above to search through posts.</p>
        </div>
    <?php elseif (empty($posts)): ?>
        <div class="no-posts" style="text-align: center;">
            <p>No posts found matching your query.</p>
        </div>
    <?php else: ?>
        <div class="posts-grid">
            <?php foreach ($posts as $post): ?>
                <article class="post-preview">
                    <h2 class="post-title">
                        <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>&category=<?php echo urlencode($post['category_slug'] ?? ''); ?>">
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
                    <a href="post.php?slug=<?php echo urlencode($post['slug']); ?>&category=<?php echo urlencode($post['category_slug'] ?? ''); ?>" class="read-more">
                        <?php echo $config['labels']['read_more'] ?? 'Read more &rarr;'; ?>
                    </a>
                </article>
            <?php endforeach; ?>
        </div>
        
        <?php if (!empty($pagination) && $pagination['total'] > 1): ?>
            <nav class="pagination">
                <?php if ($pagination['hasPrev']): ?>
                    <a href="?q=<?php echo urlencode($query); ?>&page=<?php echo $pagination['prev']; ?>" class="pagination-link prev">
                        <?php echo $config['labels']['pagination_prev'] ?? '&larr; Newer'; ?>
                    </a>
                <?php endif; ?>
                <span class="pagination-info">
                    <?php echo sprintf($config['labels']['page_indicator'] ?? 'Page %d of %d', $pagination['current'], $pagination['total']); ?>
                </span>
                <?php if ($pagination['hasNext']): ?>
                    <a href="?q=<?php echo urlencode($query); ?>&page=<?php echo $pagination['next']; ?>" class="pagination-link next">
                        <?php echo $config['labels']['pagination_next'] ?? 'Older &rarr;'; ?>
                    </a>
                <?php endif; ?>
            </nav>
        <?php endif; ?>
    <?php endif; ?>
</main>
