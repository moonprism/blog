<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>kicoe - Code</title>
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="<?php echo  kicoe\core\Config::prpr('description') ?>">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/main.min.css">
    <link rel="stylesheet" type="text/css" href="/dist/css/md.code.min.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto+Mono&display=swap" rel="stylesheet">
</head>
<body>
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
        <input id="searchInput">
    </div>
    <div id="searchResult" class="search-result">
        <?php foreach($code_list as $code) { ?>
            <div class="search-item">
                <blockquote><?php echo $code['description'] ?> <span>.<?php echo $code['lang'] ?></span></blockquote>
                <div class="tags"><?php echo $code['tags'] ?></div>
                <pre class="c-<?php echo $code['lang'] ?>"><code class="<?php echo $code['lang'] ?>"><?php echo htmlspecialchars($code['content']) ?></code><div class="markdown"></div></pre>
            </div>
        <?php } ?>
    </div>
</div>
</body>
<script src="/dist/js/main.min.js"></script>
<script type="text/javascript" src="/dist/js/markdown.min.js"></script>
<script src="//cdn.staticfile.org/highlight.js/9.18.1/highlight.min.js"></script>
<link rel="stylesheet" type="text/css" href="//cdn.staticfile.org/highlight.js/9.18.1/styles/github-gist.min.css">
<script src="/dist/js/code.min.js"></script>
</html>