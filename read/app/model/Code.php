<?php

namespace app\model;

use kicoe\core\Model;

class Code extends Model
{
    public int $id;
    public string $description;
    public string $content;
    public string $lang;
    public string $tags;
    public string $updated_time;
    public string $created_time;
    public ?string $deleted_at;
}