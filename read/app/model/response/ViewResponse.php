<?php

namespace app\model\response;

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
        $config = \kicoe\core\Link::make(\kicoe\core\Config::class);
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

        // inject
        $this->view_vars['setting'] = \kicoe\core\Link::make(\kicoe\core\Cache::class)->getArr('blog:setting');
        // 啊这...
        if (preg_match('/\*\* auto \*\*/', $this->view_vars['setting']['global_css'] ?? '')) {
            $hour = intval(date('H'));
            if ($hour <= 8 || $hour >= 23) {
                $this->view_vars['setting']['global_css'] .= '</style>
                    <link rel="stylesheet" type="text/css" href="/dist/css/dark.min.css">
                    <style>#content { margin-bottom: 40px !important;}';
            }
        }

        extract((array)$this, EXTR_SKIP);
        extract($this->view_vars, EXTR_SKIP);

        // 前台使用gulp压缩了静态文件的版本号
        $temp_header = function(){
            include $this->view_path.'/dist/header.php';
        };
        $temp_footer = function() {
            include $this->view_path.'/dist/footer.php';
        };
        extract(['_temp_header' => $temp_header, '_temp_footer' => $temp_footer]);

        include $view_file;
    }

    public function view(string $path, array $vars = []): static
    {
        $this->view_file = $path;
        if ($vars !== []) {
            $this->view_vars = $vars;
        }
        return $this;
    }

    public function handleException(\Exception $e)
    {
        $this->status($e->getCode());
        $data = ['message' => $e->getMessage()];
        $view = $this->view('error/'.$e->getCode(), $data);
        if (!file_exists($this->getFullPath())) {
            $this->view('error/index', $data)->send();
        } else {
            $view->send();
        }
    }
}
