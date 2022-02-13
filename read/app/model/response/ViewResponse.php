<?php

namespace app\model\response;

use kicoe\core\Cache;
use kicoe\core\Config;
use kicoe\core\Link;
use kicoe\core\Response;

class ViewResponse extends Response
{
    protected string $view_path;
    protected string $view_file;
    protected array $view_vars = [];
    protected string $suffix = '.php';

    public function __construct()
    {
        $config = Link::make(Config::class);
        $this->view_path = $config->get('space.view');
    }

    public function getSuffix(): string
    {
        return $this->suffix;
    }

    public function setSuffix(string $suffix)
    {
        $this->suffix = $suffix;
    }

    public function getFullPath():string
    {
        return $this->view_path.$this->view_file.$this->getSuffix();
    }

    public function send()
    {
        $view_file = $this->getFullPath();
        if (!file_exists($view_file)) {
            throw new \Exception(sprintf('view file "%s" not exists', $view_file));
        }

        extract((array)$this, EXTR_SKIP);
        extract($this->view_vars, EXTR_SKIP);

        $setting = Link::make(Cache::class)->getArr('blog:setting');
        // 前台使用gulp压缩了静态文件的版本号
        $temp_header = function() use ($setting) {
            include $this->view_path.'/dist/header.php';
        };
        $temp_footer = function() use ($setting) {
            include $this->view_path.'/dist/footer.php';
        };
        extract(['_temp_header' => $temp_header, '_temp_footer' => $temp_footer]);

        include $view_file;
    }

    public function view(string $path, array $vars = []): self
    {
        $this->view_file = $path;
        if ($vars !== []) {
            $this->view_vars = $vars;
        }
        return $this;
    }

    public function handleException(\Exception $e)
    {
        $config = Link::make(Config::class);
        $data = ['e' => $e, 'debug' => $config->get('debug')];
        if (!$data['debug']) {
            return;
        }
        $this->view('pages/error', $data)->send();
    }
}
