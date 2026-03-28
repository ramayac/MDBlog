<?php
// Load configuration if not already loaded
if (!isset($config)) {
    $config = require __DIR__ . '/../config.php';
}

// Content Security Policy
if ($config['csp_enabled'] ?? true) {
    header($config['csp_header']);
}

// Default values
$pageTitle       = $pageTitle       ?? $config['blog_name'];
$pageDescription = $pageDescription ?? $config['default_meta_description'];
$ogType          = $ogType          ?? 'website';

// Canonical URL — callers may set $pageCanonical before including this file.
// Falls back to the current request URL, stripping the ?page= parameter.
if (!isset($pageCanonical)) {
    $scheme = (!empty($_SERVER['HTTPS']) && $_SERVER['HTTPS'] !== 'off') ? 'https' : 'http';
    $host   = $_SERVER['HTTP_HOST'] ?? '';
    $uri    = $_SERVER['REQUEST_URI'] ?? '';
    // Remove page= query param so paginated URLs don't create duplicate canonicals
    $uri    = preg_replace('/([?&])page=\d+(&?)/', '$1', $uri);
    $uri    = rtrim($uri, '?&');
    $pageCanonical = $host ? $scheme . '://' . $host . $uri : null;
}

// CSS cache-busting: append the file's mtime as a query string so browsers
// pick up new styles immediately after a deployment.
$cssPath = $config['css_theme'];
$cssVer  = file_exists($cssPath) ? filemtime($cssPath) : '';
?>
<!DOCTYPE html>
<html lang="<?php echo htmlspecialchars($config['lang'] ?? 'en'); ?>">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title><?php echo htmlspecialchars($pageTitle); ?></title>
    <link id="theme-stylesheet" rel="stylesheet" href="<?php echo htmlspecialchars($cssPath); ?>?v=<?php echo $cssVer; ?>">
    <script>
        const themes = {
            'default': 'assets/css/default.style.css',
            'scrum': 'assets/css/scrum.style.css'
        };
        const savedTheme = localStorage.getItem('theme');
        if (savedTheme && themes[savedTheme]) {
            document.getElementById('theme-stylesheet').href = themes[savedTheme] + '?v=<?php echo $cssVer; ?>';
        }
        
        function toggleTheme() {
            const link = document.getElementById('theme-stylesheet');
            let currentThemeKey = localStorage.getItem('theme');
            if (!currentThemeKey) {
                const currentHref = link.getAttribute('href').split('?')[0];
                currentThemeKey = Object.keys(themes).find(key => themes[key] === currentHref) || 'default';
            }
            const newThemeKey = currentThemeKey === 'default' ? 'scrum' : 'default';
            localStorage.setItem('theme', newThemeKey);
            link.href = themes[newThemeKey] + '?v=<?php echo $cssVer; ?>';
        }
    </script>

    <?php if (!empty($pageDescription)): ?>
    <meta name="description" content="<?php echo htmlspecialchars($pageDescription); ?>">
    <?php endif; ?>

    <?php if ($pageCanonical): ?>
    <link rel="canonical" href="<?php echo htmlspecialchars($pageCanonical); ?>">
    <?php endif; ?>

    <!-- Open Graph -->
    <meta property="og:title"       content="<?php echo htmlspecialchars($pageTitle); ?>">
    <meta property="og:type"        content="<?php echo htmlspecialchars($ogType); ?>">
    <?php if ($pageCanonical): ?>
    <meta property="og:url"         content="<?php echo htmlspecialchars($pageCanonical); ?>">
    <?php endif; ?>
    <?php if (!empty($pageDescription)): ?>
    <meta property="og:description" content="<?php echo htmlspecialchars($pageDescription); ?>">
    <?php endif; ?>
</head>
