<?php

require __DIR__ . '/../vendor/autoload.php';

use kicoe\core\Link;

$config = include '../app/config.php';

$link = new Link($config);

\kicoe\core\Route::parseAnnotation(\app\controller\ArticleController::class);
\kicoe\core\Route::parseAnnotation(\app\controller\CommentController::class);
\kicoe\core\Route::parseAnnotation(\app\controller\CodeController::class);
\kicoe\core\Route::parseAnnotation(\app\controller\ApiController::class);

$link->start();
