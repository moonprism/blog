<!DOCTYPE html>
<html>
<?php $setting = kicoe\core\cache\Factory::getInstance('redis')->read('blog:setting'); ?>
<head>
    <meta charset="utf-8">
    <title><?php echo $article_obj->title;?></title>
    <meta name="keywords" content="kicoe,博客,blog<?php if($article_obj->tags!==false){foreach($article_obj->tags as $tag){echo ','.$tag['name'];}} ?>">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/comment.min.css">
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
<div id="up"><a href="javascript:up()"><b>^</b></a></div>
<div id="bg"></div>
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
<div id="content">
    <div class="article">
        <h1> <a><?php echo $article_obj->title ?></a> </h1>
        <div class="mark">
            <?php $asciinema_flag = false ?>
            <?php $math_flag = false ?>
            <svg class="icon i-date" aria-hidden="true">
                <use xlink:href="#icondate"></use>
            </svg>
            <?php echo date('Y.m.d D', strtotime($article_obj->created_time)) ?>
            <?php if($article_obj->tags){ ?> 
            <svg class="icon i-tag" aria-hidden="true">
                <use xlink:href="#icontag"></use>
            </svg>
            <?php } ?>
            <?php foreach($article_obj->tags as $tag){ ?>
                <?php if ($asciinema_flag != true) $asciinema_flag = (strtolower($tag['name']) === 'asciinema'); ?>
                <?php if ($math_flag != true) $math_flag = (strtolower($tag['name']) === 'math'); ?>
                <a href="/article/tag/<?php echo $tag['id'] ?>" style="background-color: <?php echo $tag['color'] ?>" class="tag"><?php echo $tag['name'] ?></a>
            <?php } ?>
        </div>
    </div>
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article_obj->content); ?> </textarea>

    <div class="agreement">
        parse by <a target="_blank" href="https://github.com/moonprism/markdown.js">markdown.js</a><br>
    </div>

    <div class="comment">
        <span id="to_name"></span>
        <div id="com_up">
            <input type="hidden" name="art_id" value="<?php echo $article_obj->id ?>" >
            <input type="hidden" name="to_id" value="0" >
            <label for="name">Name:</label><input id="name" name="name" autocomplete="off" placeholder="your name." type="text" >
            <label for="email">Email:</label><input id="email" name="email" autocomplete="off" placeholder="Email" type="text" >
            <div class="exp-p" id="exp_p" style="display: none;"></div>
            <textarea name="content"></textarea>
            <button id="exp" class="exp">(｡・`ω´･)</button>
            <input type="submit" class="submit" value="留言" >
        </div>
        <div class="com-list" id="com_list"></div>
        <a href="javascript:comment_list()" class="more_a" id="more_a">更多评论</a>
    </div>
</div>
<script type="text/javascript" src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script type="text/javascript">
    $('markdown').innerHTML = markdown($('text').value, main_markdown_config);
</script>
<script type="text/javascript" src="/dist/js/comment.min.js"></script>
<script src="//at.alicdn.com/t/font_1747613_gq90x707awu.js"></script>
<?php if ($asciinema_flag) { ?>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/asciinema-player@2.6.1/resources/public/css/asciinema-player.css">
    <script src="https://cdn.jsdelivr.net/npm/asciinema-player@2.6.1/resources/public/js/asciinema-player.min.js"></script>
<?php } ?>
<?php if ($math_flag) { ?>
    <script type="text/javascript" src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
</script>
<?php } ?>
<!---------高亮处理---------->
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();bg();</script>
<script><?php
    if ($setting && $setting['global_js']) {
        echo $setting['global_js'];
    }
?></script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/gruvbox-dark.min.css">
<!-------------------------->
</body>
</html>