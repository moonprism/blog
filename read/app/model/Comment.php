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
}