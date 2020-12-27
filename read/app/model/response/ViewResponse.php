<?php

namespace app\model\response;

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

    public function getSuffix()
    {
        return $this->suffix;
    }

    public function setSuffix(string $suffix)
    {
        $this->suffix = $suffix;
    }

    public function send()
    {
        if (!isset($this->view_file)) {
            parent::send();
        }
        $view_file = $this->view_path.$this->view_file.$this->getSuffix();
        if (!file_exists($view_file)) {
            throw new \Exception(sprintf('view file "%s" not exists', $view_file));
        }
        extract((array)$this, EXTR_SKIP);
        extract($this->view_vars, EXTR_SKIP);

        include $view_file;
        parent::send();
    }

    public function view(string $path, array $vars = [])
    {
        $this->view_file = $path;
        if ($vars !== []) {
            $this->view_vars = $vars;
        }
        return $this;
    }
}