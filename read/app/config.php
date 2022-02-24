<?php

return [
    'debug' => true,
    'log' => [
        'level' => \Monolog\Logger::DEBUG,
        'path' => getcwd().'/../../log/app-'.date('Y-m-d').'.log',
    ],
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
    'cas' => [
        'login_url' => 'http://localhost:8033/#/cas/',
        'auth_url' => 'http://write-api:8044/api/v1/cas/auth'
    ],
    'grpc' => [
        'host' => 'write-api',
        'port' => 2333
    ]
];
