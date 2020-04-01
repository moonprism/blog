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
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap" rel="stylesheet">
</head>
<body>
<div class="an"></div>
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
            <?php echo $article['created_time'] ?>
            <?php if($article['tags']){ ?> <b>#</b> <?php } ?>
            <?php foreach($article['tags'] as $tag) { ?>
                <a href="/article/tag/<?php echo $tag['id'] ?>" style="background-color: <?php echo $tag['color'] ?>" class="tag"><?php echo $tag['name'] ?></a>
            <?php } ?>
        </div>
        <div class="summary markdown">
            <textarea style="display:none"><?php echo $article['summary']?></textarea>
        </div>
    </div>
<?php } ?>
    </div>
    <div class="page">
        <ul>
        <?php
            if ($total_page == 1) {

            } else if ($page-1 < $total_page-$page) {
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="'.$url.$i.'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
                }
                if($page+3 < $total_page){
                    echo '<li><span>...</span></li>';
                    echo '<li><a href="'.$url.$total_page.'">'.$total_page.'</a></li>';
                }
            } else {
                if($page-3 > 1){
                    echo '<li><a href="'.$url.'1">1</a></li>';
                    echo '<li><span>...</span></li>';
                }
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="'.$url.$i.'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
                }
            }
        ?>
        </ul>
    </div>
</div>
<!-- <div id="footer">
    <span class="f_en">Powered by<a href="https://github.com/moonprism/kicoephp-src" target="_blank"> kicoephp </a></span>
</div> -->
<script type="text/javascript" src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script type="text/javascript">
    document.querySelectorAll('.summary').forEach((summary) => {
        summary.innerHTML = markdown(summary.getElementsByTagName('textarea')[0].value, main_markdown_config);
    });
</script>
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/hopscotch.min.css">
</body>
</html>
