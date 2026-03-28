<?php
$pageTitle = $pageTitle ?? $config['blog_name'];
$pageDescription = $pageDescription ?? $config['default_meta_description'];
$ogType = $ogType ?? 'website';

include __DIR__ . '/../head.php';
?>
<body>
    <div class="container">
        <?php if (!empty($menu)): ?>
        <nav class="site-menu">
            <?php 
            $menuHtml = [];
            foreach ($menu as $item) {
                if (isset($item['onclick'])) {
                    $menuHtml[] = '<a href="' . htmlspecialchars($item['url']) . '" onclick="' . htmlspecialchars($item['onclick']) . '" title="' . htmlspecialchars($item['title'] ?? '') . '">' . htmlspecialchars($item['label']) . '</a>';
                } else {
                    $menuHtml[] = '<a href="' . htmlspecialchars($item['url']) . '">' . htmlspecialchars($item['label']) . '</a>';
                }
            }
            echo implode(' | ', $menuHtml);
            ?>
        </nav>
        <?php endif; ?>

        <?php echo $content ?? ''; ?>

        <?php
            $footerVersionInfo = $blog->getVersionInfo();
        ?>
        <?php if (!empty($config['footer_content']) || !empty($footerVersionInfo['commit'])): ?>
        <footer class="site-footer">
            <?php if (!empty($config['footer_content'])): ?>
                <?php echo $blog->parseMarkdown($config['footer_content']); ?>
            <?php endif; ?>
            <?php if (!empty($footerVersionInfo['commit'])): ?>
            <small class="site-version">
                <?php if (!empty($footerVersionInfo['version'])): ?><?php echo htmlspecialchars($footerVersionInfo['version']); ?> &middot; <?php endif; ?>
                <?php echo htmlspecialchars($footerVersionInfo['commit']); ?>
                <?php if (!empty($footerVersionInfo['date'])): ?>&middot; <?php echo htmlspecialchars($footerVersionInfo['date']); ?><?php endif; ?>
            </small>
            <?php endif; ?>
            <?php include __DIR__ . '/../debug.php'; ?>
        </footer>
        <?php endif; ?>
    </div>
    
    <?php if (!empty($jsFiles)): ?>
        <?php foreach ($jsFiles as $jsFile): ?>
            <?php if (file_exists('assets/js/' . $jsFile)): ?>
                <script src="assets/js/<?php echo htmlspecialchars($jsFile); ?>"></script>
            <?php endif; ?>
        <?php endforeach; ?>
    <?php endif; ?>
</body>
</html>