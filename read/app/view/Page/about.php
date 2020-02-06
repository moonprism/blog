<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>kicoe - About</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game,about,关于">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/static/css/main.css">
    <link rel="stylesheet" type="text/css" href="/static/css/markdown.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap" rel="stylesheet">
    <style>
        .markdown {
            margin: 0 auto;
        }
        .markdown h1 {
            text-align: center
        }
        .markdown h2 a {
            font-size: 1em !important;
        }
        .markdown h1 img {
            height: 67px;
            width: 67px;
            border-radius: 50%;
        }
        .markdown i {
            font-size: .9em;
        }
        .markdown a code {
            font-size: 1em !important;
            color: #455 !important;
        }
        .markdown blockquote {
            margin: 20px 17px;
        }
        .markdown blockquote img {
            height: 22px;
            width: 22px;
            border-radius: 50%;
            position: relative;
            top: 3px;
            margin-left: 7px;
        }
        .markdown blockquote a {
            font-weight: 100;
            font-size: 17px;
            margin-left: 3px;
        }
    </style>
</head>
<body>
<div id="up" class="p_up"><a href="javascript:up()"><b>^</b></a></div>
<div id="header">
    <div class="nav">
        <ul>
            <li> <a href="/">Blog</a> </li>
            <li> <a href="/page/code/">Code</a> </li>
            <li> <a href="/page/link/">Link</a> </li>
            <li> <a href="/page/about/">About</a> </li>
        </ul>
    </div>
</div>
<div id="content" style="min-height: 0">
    <div class="markdown" id="markdown">
        <?php echo $article_obj->html ?>
    </div>
</div>
<br>
</body>
<script src="/static/js/main.js"></script>
</html>