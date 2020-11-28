<?php

namespace app\model\request;

use kicoe\core\Request;

class CommentRequest extends Request
{
    public int $art_id;
    public int $to_id = 0;
    public string $name;
    public string $email;
    public string $text;

    /**
     * @return string error 错误信息
     */
    public function filter():string
    {
        $this->name = htmlspecialchars($this->name);
        $this->email = htmlspecialchars($this->email);
        $this->text = htmlspecialchars($this->text);
        if (!preg_match('/^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,})$/', $this->email)) {
            return 'email 格式错误';
        }
        return '';
    }
}