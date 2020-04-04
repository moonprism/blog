<!DOCTYPE html>
<html>
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
</head>
<body>
<div class="an"></div>
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
            <svg t="1585976256378" class="icon-date" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5029" width="16" height="16"><path d="M283.188671 273.805914h17.059559c25.845231 0 35.825073-31.560183 35.825073-53.396418V46.060808c0-21.836235-12.794669-46.060808-38.6399-46.060808h-17.144857c-25.759933 0-37.189838 24.224573-37.189837 46.060808v181.684298c0 14.500625 14.330029 46.060808 40.089962 46.060808z m437.83357 0h17.400749c26.186422 0 37.786922-31.560183 37.786923-53.396418V46.060808C776.209913 24.224573 766.059475 0 739.873053 0H722.472303c-26.186422 0-39.236985 24.224573-39.236984 46.060808v181.684298c0 14.500625 20.300875 46.060808 37.786922 46.060808z m171.277967-187.655143H812.034985V179.125364h80.179925c44.354852 0 44.354852 43.928363 44.354853 66.958768v63.546855H91.268638v-63.546855c0-19.618492 0-66.958767 44.354852-66.958768h78.473969V86.150771h-72.076634C65.764598 86.150771 0 182.537276 0 252.481466v616.70304c0 69.858892 65.6793 154.815494 142.020825 154.815494h756.59142c69.94419 0 124.961266-78.644565 124.961266-154.900791V252.481466c-6.312037-69.94419-55.017076-166.245398-131.358601-166.245398z m40.857643 776.636401c0 44.013661-12.538776 67.811745-40.942941 67.811745H142.19142c-28.404165 0-48.278551-29.854227-48.27855-61.414411V399.193669h839.330279v463.593503zM649.627988 86.150771H376.163265v93.230487h273.550021V86.150771z" p-id="5030"></path></svg>
            <?php echo date('Y.m.d D', strtotime($article_obj->created_time)) ?>
            <?php if($article_obj->tags){ ?> 
                <svg t="1585976137037" class="icon-tag" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3386" width="18" height="18"><path d="M555.84 110.08a159.36 159.36 0 0 0-112.64-46.464H286.72A160 160 0 0 0 127.616 222.72v157.184c0 41.984 17.216 82.688 46.464 112.64l385.536 385.536a63.36 63.36 0 0 0 89.728 0l292.032-292.032a63.36 63.36 0 0 0 0-89.728L555.84 110.08zM215.424 451.072a98.56 98.56 0 0 1-29.888-71.872V222.72c0-55.936 45.824-101.12 101.12-101.12h157.184c26.752 0 52.8 10.176 71.872 29.888l390.656 390.016-300.928 300.928-390.016-391.296z m40.064-260.224h127.296v127.296H254.848V190.848h0.64z" p-id="3387"></path></svg>
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
        <a href="javascript:comment_list()" class="more_a" id="more_a">查看所有</a>
    </div>
</div>
<script type="text/javascript" src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script type="text/javascript">
    $('markdown').innerHTML = markdown($('text').value, main_markdown_config);
</script>
<script type="text/javascript" src="/dist/js/comment.min.js"></script>
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
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/gruvbox-dark.min.css">
<!-------------------------->
</body>
</html>