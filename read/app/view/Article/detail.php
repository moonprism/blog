<!DOCTYPE html>
<html>
<?php $setting = \kicoe\core\Link::make(\kicoe\core\Cache::class)->getArr('blog:setting'); ?>
<?php $config = \kicoe\core\Link::make(\kicoe\core\Config::class) ?>
<head>
    <meta charset="utf-8">
    <title><?php echo $article->title;?></title>
    <meta name="keywords" content="kicoe,博客,blog<?php foreach($article->getTags() as $tag){ echo ','.$tag->name; } ?>">
    <meta name="description" content="<?php echo $config->get('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/comment.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Space+Mono&family=Noto+Sans+SC&family=Roboto+Mono&display=swap" rel="stylesheet">
    <style><?php echo $setting['global_css'] ?? '' ?></style>
</head>
<body>
<div class="an" style=" <?php
if (isset($setting['background_image'])) {
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
        <h1> <a><?php echo $article->title ?></a> </h1>
        <div class="mark">
            <?php $math_flag = false ?>
            <svg class="icon i-date" aria-hidden="true">
                <use xlink:href="#icondate"></use>
            </svg>
            <?php echo date('Y.m.d D', strtotime($article->created_time)) ?>
            <?php if($article->getTags()){ ?>
            <svg class="icon i-tag" aria-hidden="true">
                <use xlink:href="#icontag"></use>
            </svg>
            <?php } ?>
            <?php foreach($article->getTags() as $tag){ ?>
                <span style="border-color:<?php echo $tag->color ?? 'none'; ?>;color:<?php echo $tag->color ?? 'none'; ?>">
                <?php if ($math_flag != true) $math_flag = (strtolower($tag->name) === 'math'); ?>
                    <a href="/article/tag/<?php echo $tag->id ?>" style="background-color: <?php echo $tag->color ?>" class="tag"><?php echo $tag->name ?></a>
                </span>
            <?php } ?>
        </div>
    </div>
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article->content); ?> </textarea>

    <div class="agreement">
        parse by <a target="_blank" href="https://github.com/moonprism/markdown.js">markdown.js</a><br>
    </div>

    <div class="comment">
        <span id="to_name"></span>
        <div id="com_up">
            <input type="hidden" name="art_id" value="<?php echo $article->id ?>" >
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
<script src="//at.alicdn.com/t/font_1922445_olfo3nycsle.js"></script>
<?php if ($math_flag) { ?>
    <script type="text/javascript" src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
</script>
<?php } ?>
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();bg();</script>
<script><?php echo $setting['global_css'] ?? ''; ?></script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/gruvbox-dark.min.css">
</body>
</html>
