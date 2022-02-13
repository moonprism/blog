<?php

namespace app\model\request;

use kicoe\core\Request;

class LinkRequest extends Request
{
    public string $link;
    public string $email;
    public string $avatar;
    public string $name;
    public string $icon;
    public string $message;

    public string $data;

    public function check()
    {

    }
}
