<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>kicoe - Blog</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
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
<div id="content">
    <div class="article-list">
<?php foreach($article_list as $article){ ?>
    <div class="article">
        <h1><a href="/article/id/<?php echo $article['id'] ?>"><?php echo stripslashes($article['title']) ?></a></h1>
        <div class="mark">
            <?php echo $article['updated_time'] ?>
            <?php if($article['tags']){ ?> <b>#</b> <?php } ?>
            <?php foreach($article['tags'] as $tag) { ?>
                <a style="background-color: <?php echo $tag['color'] ?>" class="tag"><?php echo $tag['name'] ?></a>
            <?php } ?>
        </div>
        <div class="summary">
            <?php if($article['image']){ ?>
                <img src="<?php echo kicoe\core\Config::prpr('image_cdn').$article['image'] ?>">
            <?php } ?>
            <?php echo stripslashes($article['summary']) ?>
        </div>
    </div>
<?php } ?>
    </div>
    <div class="page">
        <ul>
        <?php
        if ($page-1 < $total_page-$page) {
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="/article/page/'.$i.'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
                }
                if($page+3 < $total_page){
                    echo '<li><span>...</span></li>';
                    echo '<li><a href="/article/page/'.$total_page.'">'.$total_page.'</a></li>';
                }
            } else {
                if($page-3 > 1){
                    echo '<li><a href="/article/page/1">1</a></li>';
                    echo '<li><span>...</span></li>';
                }
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="/article/page/'.$i.'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
                }
            }
        ?>
        </ul>
    </div>
</div>
<div id="footer">
    <span class="f_en">Powered by<a href="https://github.com/moonprism/kicoephp-src" target="_blank"> kicoephp </a></span>
</div>
<script type="text/javascript" src="/dist/js/main.min.js"></script>
</body>
</html>
