<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>blog - <?php echo $article_obj->title;?></title>
    <meta name="keywords" content="kicoe,博客,blog<?php if($article_obj->tags!==false){foreach($article_obj->tags as $tag){echo ','.$tag['name'];}} ?>">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <!--[if lt IE 9]><script>window.location.href="/page/hack.html";</script><![endif]-->
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/comment.min.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap" rel="stylesheet">
</head>
<body>
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
            <?php echo $article_obj->created_time ?>
                <b>#</b>
                <?php foreach($article_obj->tags as $tag){ ?>
                    <a href="/article/tag/<?php echo $tag['id'] ?>" style="background-color: <?php echo $tag['color'] ?>" class="tag"><?php echo $tag['name'] ?></a>
                <?php } ?>
        </div>
    </div>
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article_obj->content); ?>
    </textarea>
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
    $('markdown').innerHTML = replace_sym(markdown($('text').value));
</script>
<script type="text/javascript" src="/dist/js/comment.min.js"></script>
<!---------高亮处理---------->
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();bg();</script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/gruvbox-dark.min.css">
<!-------------------------->
</body>
</html>