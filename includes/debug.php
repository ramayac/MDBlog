<?php
// Calculate render time
if ($config['show_render_time']) {
    $end_time = microtime(true);
    $render_time = round(($end_time - $start_time) * 1000, 2);
    echo "<small>Page rendered in {$render_time}ms</small>\n";
}
?>