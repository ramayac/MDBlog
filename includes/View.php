<?php

class View {
    private static $basePath = __DIR__ . '/templates/';

    /**
     * Render a view within the layout.
     */
    public static function render($view, array $data = []) {
        $content = self::renderTemplate($view, $data);
        $layoutData = array_merge($data, ['content' => $content]);
        echo self::renderTemplate('layout.php', $layoutData);
    }

    /**
     * Render a raw template file and return the output.
     */
    public static function renderTemplate($templateFile, array $data = []) {
        extract($data);
        
        // Start output buffering
        ob_start();
        $path = self::$basePath . $templateFile;
        if (file_exists($path)) {
            include $path;
        } else {
            echo "Template {$templateFile} not found.";
        }
        return ob_get_clean();
    }
}
