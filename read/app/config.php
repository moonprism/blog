<?php

return [
    'db' => [
        'hostname'    => 'mysql',
        'database'    => 'blog',
        'username'    => 'root',
        'password'    => '123456'
    ],
    'elastic' => [
        'host'  => 'elastic',
    ],
    'redis' => [
        'host'  => 'redis',
    ],
    'route' => include 'route.php',
    'route_cache' => 'redis',
    'cp'    => 'cache',
    'test'  => false,
    'description' => 'dd',
    'cache_comment_prefix' => 'blog-comment-',
    'elastic_code_index' => 'elastic-code-index',
    'elastic_code_type' => 'elastic-code-type',
    'image_cdn' => 'https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/',
];
