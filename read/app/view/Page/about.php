<!DOCTYPE html>
<html>
<?php $setting = kicoe\core\cache\Factory::getInstance('redis')->read('blog:setting'); ?>
<head>
    <meta charset="utf-8">
    <title>about | kicoe</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game,about,关于">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/md.about.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+SC&family=Noto+Sans:ital@1&family=Roboto+Mono&display=swap" rel="stylesheet">
    <style><?php
        if ($setting && $setting['global_css']) {
            echo $setting['global_css'];
        }
    ?></style>
</head>
<body>
<div class="an" style=" <?php
$setting = kicoe\core\cache\Factory::getInstance('redis')->read('blog:setting');
if ($setting && $setting['background_image']) {
    echo 'background-image: url('.$setting['background_image'].')';
}
?>"></div>
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
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article_obj->content); ?> </textarea>
</div>
<br>
</body>
<script src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script type="text/javascript">
    $('markdown').innerHTML = replace_sym(markdown($('text').value, main_markdown_config));
</script>
<script><?php
    if ($setting && $setting['global_js']) {
        echo $setting['global_js'];
    }
?></script>
</html>