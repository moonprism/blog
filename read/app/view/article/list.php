<!DOCTYPE html>
<html>
<?php $config = \kicoe\core\Link::make(\kicoe\core\Config::class) ?>
<head>
    <meta charset="utf-8">
    <title>kicoe</title>
    <link rel="shortcut icon" href="/favicon.png">
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="<?php echo $config->get('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Space+Mono&family=Noto+Sans+SC&family=Roboto+Mono&display=swap" rel="stylesheet">
    <style>
<?php echo $setting['global_css'] ?? '' ?>

    </style>
</head>
<body>
<div class="an" style=" <?php
if (isset($setting['background_image'])) {
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
<div id="content" class="article-block">
<?php if($tag){ ?>
    <div style="border-color: <?php echo $tag->color ?>"><div class="tag-banner">当前标签<span class="t" style="background-color: <?php echo $tag->color; ?>"><?php echo $tag->name; ?></span>, 文章数<span class="c"><?php echo $tag->count; ?><span></div></div>
<?php } ?>
    <div class="article-list">
<?php foreach($article_list as $article){ ?>
        <div class="article">
            <h1><a href="/article/id/<?php echo $article->id ?>"><?php echo stripslashes($article->title) ?></a></h1>
                <div class="mark">
                    <svg class="icon i-date" aria-hidden="true">
                        <use xlink:href="#icondate"></use>
                    </svg>
                    <?php echo date('Y.m.d D', strtotime($article->created_time)) ?>

<?php if($article->getTags()){ ?>
                    <svg class="icon i-tag" aria-hidden="true">
                        <use xlink:href="#icontag"></use>
                    </svg>
<?php } ?>
<?php foreach($article->getTags() as $t) { ?>
                    <span style="border-color:<?php echo $t->color; ?>;color:<?php echo $t->color; ?>">
                        <a href="/article/tag/<?php echo $t->id ?>" style="<?php if ($tag && $t->id == $tag->id) echo 'background-color: inherit;color: inherit;border-color: inherit'; else echo 'background-color: '.$t->color; ?>" class="tag"><?php echo $t->name ?></a>
                    </span>
<?php } ?>
                </div>
                <div class="summary markdown">
                    <script type='text/html' style="display:none"><?php echo $article->summary?></script>
                </div>
            </div>
<?php } ?>
    </div>
<?php include(dirname(__DIR__).'/article/page.php') ?>
</div>
<script src="/dist/js/main.min.js"></script>
<script src="/dist/js/markdown.min.js"></script>
<script>
    document.querySelectorAll('.summary').forEach((summary) => {
        summary.innerHTML = markdown(summary.getElementsByTagName('script')[0].innerText, main_markdown_config);
    });
</script>
<script src="//at.alicdn.com/t/font_1747613_gq90x707awu.js"></script>
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<script>
<?php echo $setting['global_js'] ?? ''; ?>

</script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/hopscotch.min.css">
</body>
</html>
