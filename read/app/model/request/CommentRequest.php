<?php

namespace app\model\request;

use kicoe\core\Request;

class CommentRequest extends Request
{
    public int $top_id = 0;
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
        $name_len = mb_strlen($this->name);
        $text_len = mb_strlen($this->text);

        if ($name_len === 0) {
            return '(,,Ծ‸Ծ,,) 不想告诉我名字吗';
        }
        if ($name_len > 200) {
            return ' (ಥ_ಥ) 名字不要取得太长啦';
        }
        if (!preg_match('/^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,})$/', $this->email)) {
            return '(･_-｡ ) 再检查下邮箱格式吧';
        }
        if ($text_len > 350) {
            return '(ღ˘⌣˘ღ) 留言内容实在太多了';
        }
        if ($text_len === 0) {
            return '( T﹏T ) 没有什么想说的吗';
        }
        return '';
    }
}
