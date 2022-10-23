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
        if (!preg_match('/^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,})$/', $this->email) && !$this->data) {
            return '(*゜ロ゜)ノ 真没什么想说的吗?';
        }
    }
}
