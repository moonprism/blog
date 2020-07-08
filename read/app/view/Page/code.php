<!DOCTYPE html>
<html>
<?php $setting = kicoe\core\cache\Factory::getInstance('redis')->read('blog:setting'); ?>
<head>
    <meta charset="utf-8">
    <title>code | kicoe</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/md.code.min.css">
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
    <div class="search-input">
        <input id="searchInput" autocomplete="off">
    </div>
    <div id="load"></div>
    <div id="searchResult" class="search-result">
        <?php foreach($code_list as $code) { ?>
            <div class="search-item">
                <blockquote>
                    <svg class="icon icon-code" aria-hidden="true">
                        <use xlink:href="#icon-<?php echo $code['lang'] ?>"></use>
                    </svg>
                    <?php echo $code['description'] ?> <span>.<?php echo $code['lang'] ?></span>
                </blockquote>
                <?php if($code['tags']) { ?><div class="tags"><?php echo $code['tags'] ?></div><?php } ?>
                <pre class="c-<?php echo $code['lang'] ?>"><code class="<?php echo $code['lang'] ?>"><?php echo htmlspecialchars($code['content']) ?></code><div class="markdown"></div></pre>
            </div>
        <?php } ?>
        <?php if($next_page) { ?><div class="next"><a href="/page/code/<?php echo $next_page;?>">Next</a></div><?php } ?>
    </div>
</div>
</body>
<script src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script><?php
    if ($setting && $setting['global_js']) {
        echo $setting['global_js'];
    }
?></script>
<script src="//at.alicdn.com/t/font_1747612_imjmg02wt19.js"></script>
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/github-gist.min.css">
<script src="/dist/js/code.min.js"></script>
</html>