<?php

namespace app\model\response;

use kicoe\core\Response;

class ApiResponse extends Response
{
    public int $code = 200;
    public string $message = '';

    public $data = null;

    public function setBodyStatus(int $code, string $message):self
    {
        $this->code = $code;
        $this->message = $message;
        return $this;
    }
}