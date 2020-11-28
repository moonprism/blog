<?php

namespace app\model;

use kicoe\core\Model;

class Comment extends Model
{
    public int $id;
    public int $art_id;
    public int $to_id;
    public string $name;
    public string $email;
    public string $text;
    public string $created_time;

    /**
     * 属性融合
     * @param object $obj
     * @return self
     */
    public static function createByOther(object $obj) {
        $comment = new self();
        foreach (get_object_vars($obj) as $k => $v) {
            $comment->$k = $v;
        }
        return $comment;
    }
}