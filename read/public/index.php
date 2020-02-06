<?php

ini_set('display_errors',1);            //错误信息
ini_set('display_startup_errors',1);    //php启动错误信息
error_reporting(-1);                    //打印出所有的 错误信息

define('PUB_PATH', __DIR__);
// app path
define('APP_PATH', __DIR__ . '/../app/');
// autoload
require __DIR__ . '/../vendor/autoload.php';
// link-start 
kicoe\core\Load::link_start();