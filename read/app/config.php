<?php

return [
    'redis' => [
        'host' => 'redis',
        'port' => 6379
    ],
    'mysql' => [
        'host' => 'mysql',
        'port' => 3306,
        'db' => 'blog',
        'user' => 'root',
        'passwd' => '123456',
        'charset' => 'utf8mb4',
    ],
    'space' => [
        // view 路径
        'view' => getcwd().'/../app/view/',
    ],
    'description' => 'aqua',
    'es' => [
        'host' => 'elastic',
        'code_index' => 'elastic-code-index',
        'code_type' => 'elastic-code-type'
    ],
    'cas' => [
        'login_url' => 'http://127.0.0.1:8080/#/cas/',
        'auth_url' => 'http://192.168.80.1:2999/api/v1/cas/auth'
    ],
    'grpc' => [
        'host' => '192.168.80.1',
        'port' => 2333
    ]
];
