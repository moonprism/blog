<?php

ini_set('date.timezone','Asia/Shanghai');

require __DIR__ . '/../vendor/autoload.php';

use kicoe\core\Link;
use kicoe\core\Route;

use Monolog\Logger;
use Monolog\Handler\StreamHandler;
use Monolog\Formatter\LineFormatter;

$config = include '../app/config.php';

$whoops = new \Whoops\Run;

$dateFormat = "Y-m-d H:i:s";
$output = "[%datetime%] %level_name%: %message% %context% %extra%\n";
$formatter = new LineFormatter($output, $dateFormat, true, true);
$stream = new StreamHandler($config['log']['path'], $config['log']['level']);
$stream->setFormatter($formatter);

$log = new Logger('app');
$log->pushHandler($stream);

if ($config['debug']) {
    if ($_SERVER['HTTP_KICOE_FETCH_TYPE'] ?? false) {
        $handler = new \Whoops\Handler\JsonResponseHandler();
        $handler->addTraceToOutput(true);
    } else {
        $handler = new \Whoops\Handler\PrettyPageHandler();
    }
    $whoops->pushHandler($handler);
}

$whoops->pushHandler(function ($exception, $inspector, $run) use ($log) {
    $whoops = new \Whoops\Run;
    $whoops->allowQuit(false);
    $whoops->writeToOutput(false);
    $whoops->pushHandler(new \Whoops\Handler\PlainTextHandler());
    $text = $whoops->handleException($exception);
    $log->debug($text);
});

$whoops->register();

Link::bind('log', $log);

$link = new Link($config);

Route::parseAnnotation(\app\controller\ArticleController::class);
Route::parseAnnotation(\app\controller\CommentController::class);
Route::parseAnnotation(\app\controller\CodeController::class);
Route::parseAnnotation(\app\controller\ApiController::class);
Route::parseAnnotation(\app\controller\AuthController::class);

try {
    $link->start();
} catch (Exception $e) {
    $response = Link::make(\kicoe\core\Response::class);
    if ($response instanceof \app\model\response\ViewResponse) {
        $response->handleException($e);
    }
    throw $e;
}
