<?php

namespace app\model;

use kicoe\core\Model;

class BaseModel extends Model
{
    /**
     * 属性融合
     * @param object $obj
     * @return static
     */
    public static function createByOther(object $obj): self
    {
        $comment = new static();
        foreach (get_object_vars($obj) as $k => $v) {
            $comment->$k = $v;
        }
        return $comment;
    }
}
