<!DOCTYPE html>
<html>
<?php $config = \kicoe\core\Link::make(\kicoe\core\Config::class) ?>
<head>
    <meta charset="utf-8">
    <title>about | kicoe</title>
    <link rel="shortcut icon" href="/favicon.png">
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game,about,关于">
    <meta name="description" content="<?php echo $config->get('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/markdown.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/md.about.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Space+Mono&family=Noto+Sans+SC&family=Roboto+Mono&display=swap" rel="stylesheet">
    <style><?php echo $setting['global_css'] ?? '' ?></style>
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
<div id="content" class="about-block" style="min-height: 0">
    <div class="markdown" id="markdown"></div>
    <textarea id="text" style="display: none;">
<?php echo trim($article->content); ?> </textarea>
</div>
<br>
</body>
<script src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script type="text/javascript">
    $('markdown').innerHTML = replace_sym(markdown($('text').value, main_markdown_config));
</script>
<script>
<?php echo $setting['global_js'] ?? ''; ?>
</script>
<script type="text/javascript" src="//at.alicdn.com/t/font_2433958_0geqbql8k41k.js"></script>
</html>
