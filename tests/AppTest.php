<?php
use PHPUnit\Framework\TestCase;

class AppTest extends TestCase {
    private function runPage($file, $getParams = [], $requestUri = null) {
        $query = http_build_query($getParams);
        $uri = $requestUri ?? ('/' . $file);
        if ($query !== '') {
            $uri .= '?' . $query;
        }
        
        $code = '$_GET = unserialize(base64_decode(\'' . base64_encode(serialize($getParams)) . '\')); ';
        $code .= '$_SERVER["REQUEST_URI"] = "' . $uri . '"; ';
        $code .= 'require "' . $file . '";';
        
        $cmd = 'php -r ' . escapeshellarg($code);
        exec($cmd, $output, $returnVar);
        return implode("\n", $output);
    }
    
    public function testHome() {
        $html = $this->runPage('index.php', []);
        $this->assertStringContainsString('Rodrigo A.', $html);
        $this->assertStringContainsString('srbyte', $html, 'Should contain categories cards');
        $this->assertStringNotContainsString('What are you looking for?', $html, 'Should not contain search form on home');
    }

    public function testCategory() {
        $html = $this->runPage('index.php', ['category' => 'srbyte']);
        $this->assertStringContainsString('srbyte', $html);
        $this->assertStringContainsString('Read more', $html, 'Should show posts list');
    }

    public function testSearchLayout() {
        $html = $this->runPage('index.php', ['q' => 'linux']);
        $this->assertStringContainsString('Search Results for &quot;linux&quot;', $html);
        $this->assertStringContainsString('What are you looking for?', $html);
        $this->assertStringContainsString('Showing results for', $html);
    }

    public function testSearchWorking() {
        $html = $this->runPage('index.php', ['q' => 'tiempo']); // known tag/word in first post
        $this->assertStringContainsString('Read more', $html, 'Should find posts');
        $this->assertStringContainsString('12:34:56', $html, 'Should find the specific post');
    }
    
    public function testPost() {
        $html = $this->runPage('index.php', ['slug' => 'srbyte-12-34-56-7-8-9-y-el-tiempo'], '/post.php');
        $this->assertStringContainsString('12:34:56', $html, 'Post title should be displayed');
        $this->assertStringContainsString('Rodrigo Amaya', $html, 'Author should be displayed');
    }

    public function test404() {
        $html = $this->runPage('index.php', ['slug' => 'does-not-exist'], '/post.php');
        $this->assertStringContainsString('404', $html);
        $this->assertStringContainsString('Not Found', $html);
    }
}
