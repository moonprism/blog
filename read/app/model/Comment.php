<?php

namespace app\model;

class Comment extends BaseModel
{
    public int $id;
    public int $art_id;
    public int $top_id;
    public int $to_id;
    public string $name;
    public string $email;
    public string $text;
    public string $created_time;
    public ?string $deleted_at;

    public static function cWhere(int $art_id, int $top_id)
    {
        return (new self)->columns('id', 'to_id', 'name', 'email', 'text', 'created_time')
            ->where('art_id', $art_id)
            ->where('top_id', $top_id)
            ->where('comment.deleted_at is null');
    }
}
