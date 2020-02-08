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
    <link rel="stylesheet" type="text/css" href="/static/css/md.about.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap" rel="stylesheet">
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