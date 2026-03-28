<main class="main-content">
    <div class="no-posts">
        <h2><?php echo $config['labels']['not_found_title'] ?? '404 &mdash; Post Not Found'; ?></h2>
        <p><?php echo $config['labels']['not_found_message'] ?? 'The post you&rsquo;re looking for doesn&rsquo;t exist.'; ?></p>
        <a href="index.php" class="back-to-home"><?php echo $config['labels']['back_to_all'] ?? '&larr; Back to all posts'; ?></a>
    </div>
</main>
