<?php
// Load configuration if not already loaded
if (!isset($config)) {
    $config = require __DIR__ . '/../config.php';
}

// Default values
$pageTitle = $pageTitle ?? $config['blog_name'];
$pageDescription = $pageDescription ?? $config['default_meta_description'];
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title><?php echo htmlspecialchars($pageTitle); ?></title>
    <link rel="stylesheet" href="<?php echo $config['css_theme']; ?>">
    
    <?php if (!empty($pageDescription)): ?>
    <meta name="description" content="<?php echo htmlspecialchars($pageDescription); ?>">
    <?php endif; ?>
</head>
