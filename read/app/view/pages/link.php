<!DOCTYPE html>
<html>
<?php $setting = \kicoe\core\Link::make(\kicoe\core\Cache::class)->getArr('blog:setting'); ?>
<?php $config = \kicoe\core\Link::make(\kicoe\core\Config::class) ?>
<head>
    <meta charset="utf-8">
    <title>link | kicoe</title>
    <meta name="keywords" content="kicoe,博客,blog,留言板,友情链接">
    <meta name="description" content="<?php echo $config->get('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/md.link.min.css">
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
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article->content); ?> </textarea>
    <div class="comment">
        <span id="to_name"></span>
        <div id="com_up">
            <input type="hidden" name="art_id" value="<?php echo $article->id ?>" >
            <input type="hidden" name="to_id" value="0" >
            <label for="name">🙎 Name:</label><input id="name" name="name" autocomplete="off" placeholder="your name." type="text" >
            <label for="email">✉️ Email:</label><input id="email" name="email" autocomplete="off" placeholder="Email" type="text" >
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
<script src="//at.alicdn.com/t/font_1922445_olfo3nycsle.js"></script>
<script type="text/javascript">
    $('markdown').innerHTML = markdown($('text').value, main_markdown_config);
</script>
<script><?php echo $setting['global_css'] ?? ''; ?></script>
<script type="text/javascript" src="/dist/js/comment.min.js"></script>
</body>
</html>
