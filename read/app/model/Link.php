<?php

namespace app\model;

class Link extends BaseModel
{
    public int $id;
    public string $link;
    public string $email;
    public string $avatar;
    public string $name;
    public string $icon;
    public string $message;
    public string $data;
    public string $created_time;

    const STATUS_ACCEPT = 1;
    const STATUS_REJECT = 2;
}
