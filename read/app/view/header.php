<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="UTF-8">
    <title>kicoe's Blog</title>
    <link rel="shortcut icon" href="/favicon.png">
    <meta name="keywords" content="kicoe,博客,blog,代码,code,游戏,game">
    <meta name="description" content="静默空白.">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <link rel="stylesheet" type="text/css" href="/dist/css/app.min.css">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Space+Mono&family=Work+Sans:wght@400;500&display=swap" rel="stylesheet">
    <style>
<?php echo $setting['global_css'] ?? '' ?>

    </style>
</head>
<body>
<div class="an" style="<?php
if (isset($setting['background_image'])) {
    echo 'background-image: url('.$setting['background_image'].')';
}
?>"></div>
<div id="up"></div>
<div id="header">
    <div class="nav">
        <ul>
            <li> <a href="/">Blog</a> </li>
            <li> <a href="/code">Code</a> </li>
            <li> <a href="/link">Link</a> </li>
            <li> <a href="/about">About</a> </li>
        </ul>
    </div>
</div>
