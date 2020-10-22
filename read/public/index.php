<?php
require __DIR__ . '/../vendor/autoload.php';

use kicoe\core\Link;

$link = new Link(
    [
        'cache' => false,
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
            // 自动解析路由注释路径
            'annotation' => getcwd().'/./controller/',
            // view 路径
            'view' => getcwd().'/../app/view/',
        ],
        'description' => 'haha',
    ]
);

\kicoe\core\Route::parseAnnotation(\app\controller\ArticleController::class);
\kicoe\core\Route::parseAnnotation(\app\controller\CommentController::class);
\kicoe\core\Route::parseAnnotation(\app\controller\CodeController::class);

$link->start();
