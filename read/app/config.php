<?php

return [
    'debug' => true,
    'log' => [
        'level' => \Monolog\Logger::DEBUG,
        'path' => getcwd().'/../../log/app-'.date('Y-m-d').'.log',
    ],
    'redis' => [
        'host' => 'redis',
        'port' => 6380
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
        'login_url' => 'http://127.0.0.1:8080/#/cas/',
        'auth_url' => 'http://host.docker.internal:8034/api/v1/cas/auth'
    ],
    'grpc' => [
        'host' => 'host.docker.internal',
        'port' => 2333
    ]
];
