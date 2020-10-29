<?php

require __DIR__ . '/../vendor/autoload.php';

use kicoe\core\Link;
use kicoe\core\Route;

$config = include '../app/config.php';

$link = new Link($config);

Route::parseAnnotation(\app\controller\ArticleController::class);
Route::parseAnnotation(\app\controller\CommentController::class);
Route::parseAnnotation(\app\controller\CodeController::class);
Route::parseAnnotation(\app\controller\ApiController::class);

$link->start();
