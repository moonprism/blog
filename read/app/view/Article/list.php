<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>kicoe</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+SC&family=Noto+Sans:ital@1&family=Roboto+Mono&display=swap" rel="stylesheet">
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
            <svg t="1585976256378" class="icon-date" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5029" width="16" height="16"><path d="M283.188671 273.805914h17.059559c25.845231 0 35.825073-31.560183 35.825073-53.396418V46.060808c0-21.836235-12.794669-46.060808-38.6399-46.060808h-17.144857c-25.759933 0-37.189838 24.224573-37.189837 46.060808v181.684298c0 14.500625 14.330029 46.060808 40.089962 46.060808z m437.83357 0h17.400749c26.186422 0 37.786922-31.560183 37.786923-53.396418V46.060808C776.209913 24.224573 766.059475 0 739.873053 0H722.472303c-26.186422 0-39.236985 24.224573-39.236984 46.060808v181.684298c0 14.500625 20.300875 46.060808 37.786922 46.060808z m171.277967-187.655143H812.034985V179.125364h80.179925c44.354852 0 44.354852 43.928363 44.354853 66.958768v63.546855H91.268638v-63.546855c0-19.618492 0-66.958767 44.354852-66.958768h78.473969V86.150771h-72.076634C65.764598 86.150771 0 182.537276 0 252.481466v616.70304c0 69.858892 65.6793 154.815494 142.020825 154.815494h756.59142c69.94419 0 124.961266-78.644565 124.961266-154.900791V252.481466c-6.312037-69.94419-55.017076-166.245398-131.358601-166.245398z m40.857643 776.636401c0 44.013661-12.538776 67.811745-40.942941 67.811745H142.19142c-28.404165 0-48.278551-29.854227-48.27855-61.414411V399.193669h839.330279v463.593503zM649.627988 86.150771H376.163265v93.230487h273.550021V86.150771z" p-id="5030"></path></svg>
            <?php echo date('Y.m.d D', strtotime($article['created_time'])) ?>
            <?php if($article['tags']){ ?> 
                <svg t="1585976137037" class="icon-tag" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3386" width="18" height="18"><path d="M555.84 110.08a159.36 159.36 0 0 0-112.64-46.464H286.72A160 160 0 0 0 127.616 222.72v157.184c0 41.984 17.216 82.688 46.464 112.64l385.536 385.536a63.36 63.36 0 0 0 89.728 0l292.032-292.032a63.36 63.36 0 0 0 0-89.728L555.84 110.08zM215.424 451.072a98.56 98.56 0 0 1-29.888-71.872V222.72c0-55.936 45.824-101.12 101.12-101.12h157.184c26.752 0 52.8 10.176 71.872 29.888l390.656 390.016-300.928 300.928-390.016-391.296z m40.064-260.224h127.296v127.296H254.848V190.848h0.64z" p-id="3387"></path></svg>
            <?php } ?>
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
