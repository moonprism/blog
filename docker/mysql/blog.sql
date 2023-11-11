CREATE DATABASE `blog`;
USE `blog`;

-- todo 
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `avatar` varchar(500) NOT NULL DEFAULT '',
    `password` char(60) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(64) NOT NULL,
    -- draft publish system ...
    `status` tinyint(3) NOT NULL,
    `image` varchar(64) NOT NULL DEFAULT '',
    `summary` varchar(500) NOT NULL DEFAULT '',
    `content` text NOT NULL,
    `rune` int(10) UNSIGNED NOT NULL DEFAULT 0,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `key` varchar(255) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(40) NOT NULL,
    `color` varchar(20) NOT NULL DEFAULT '#000000',
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    KEY `name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `art_id` int(11) NOT NULL,
    `tag_id` int(11) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `art_id` (`art_id`),
    KEY `tag_id` (`tag_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `art_id` int(11) NOT NULL,
    `to_id` int(11) NOT NULL,
    `top_id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `text` text NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    KEY `art_top` (`art_id`, `top_id`),
    KEY `to_id` (`to_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `link`;
CREATE TABLE `link`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `link` varchar(500) NOT NULL,
    `email` varchar(500) NOT NULL,
    `status` tinyint(1) unsigned NOT NULL DEFAULT 0,
    `avatar` varchar(500) NOT NULL,
    `name` varchar(255) NOT NULL,
    `icon` varchar(500) NOT NULL,
    `message` varchar(500) NOT NULL,
    `data` text NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    KEY `email_status` (`email`, `status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `code`;
CREATE TABLE `code`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `description` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `lang` varchar(20) NOT NULL,
    `tags` varchar(255) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `article` (`id`, `title`, `status`, `image`, `summary`, `content`, `created_time`, `updated_time`, `deleted_at`) VALUES
(1,	'Link',	3,	'',	'',	'# <svg style=\"font-size:23px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-aixin1\"></use></svg>友人帐\n\n@@@\n\n> [![](STzvbpDkfSTIXoHHysUR.jpg)~~Roki~~`💢*域名不续费了是吧*`]()\n\n> [![](KVrKQpxLKOQSCUndrzrs.jpg)Wrath`:pill:*呐呐呐呐呐*`](https://wrath.cc/)\n\n> [![](xFPCVcPzNHbhRmoiWyad.jpg)Lionad`:rocket:*可爱可爱可爱捏*`](http://www.lionad.art/)\n\n> [![](ivvmJbFUjpOscOSeREMH.jpg)Innei`:evergreen_tree:*致虚极，守静笃*`](https://innei.ren/)\n\n> [![](https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/cc/blogvg.ico)生如夏花`:rabbit:*博客复活!*`](http://www.blog.vg/)\n\n@@@+\n\n### <svg style=\"font-size:20px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-RectangleCopy\"></use></svg> Collection<span style=\"font-size:13px\">（单链）</span>\n\n> [<img src=\"https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/gnDtqzYhLVVBIGJkYzVY.jpg\" style=\"border-radius:0;box-shadow:none\">致郁系`:cherry:*页面设计超好看*`](http://www.juroku.net/)\n\n> [![](https://blog.cha.moe/image/avatar.svg)抹茶爱吃猹`🐱*一隻没有性别的猫耳生物，或许并不属于地球。*`](https://blog.cha.moe/)\n\n>[d] [![](https://blog.rxliuli.com/medias/avatar.jpg)rxliuli`🌀*这里是吾辈的博客，主要记录一些技术相关的东西呢 ヾ(＠^∇^＠)ノ*`](https://blog.rxliuli.com/about/)\n\n> [![](https://blog.k8s.li/avatar.png)木子的博客`🌰*垃圾佬、搬砖社畜、运维工程师*`](https://blog.k8s.li)\n\nxxx\n***\n<p style=\"text-align:center\" id=\"p5\">\n	<a><span>t</span>ake your heart...</a>\n</p>\n\n<div style=\"text-align:center\">\n	<canvas id=\"p5\" width=\"267\" height=\"62\"></canvas>\n</div>\nxxx\n\n<h1 id=\"board\"><svg style=\"font-size:25px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-pingguo1\"></use></svg>留言板</h1>\n\nxxx\n\n>[info] 名称: kicoe\'s Blog\n简介: 白天打游戏的废人\n链接: `https://kicoe.com`\n图标: <https://gravatar.cat.net/avatar/7c6d3737a25a9ec47b5439ec123bd1df>\n符号: ⭐️\n\n***\n\n---\n\nxxx',	'2020-02-05 23:38:36',	'2023-10-18 04:44:07',	NULL),
(2,	'About',	3,	'',	'',	'# *![](https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640)*\n\n#### <svg style=\"font-size: 19px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mao\"></use></svg> <span style=\"color: #5c6bc0\"> the blog</span>\n\n* 很久之前拥有过一个几十兆内存的服务器.\n\n* 用 PHP 写了个简单的博客框架，用来分享技术/游戏/生活类的文章.\n\n* 换新服务器后用 Golang + Vue 重写了后台，可以在网页上使用 vim 了.\n\n* <a href=\"/link\">link</a>页面欢迎留言(~~勾搭~~，等网站内容足够充实后可能会去主动申请些友链.(大概率还是继续当只小透明~)\n\n* 擅自添加了一些喜欢的博客单链，介意的话请务必联系[✉️](mailto:ankicoe@gmail.com)\n\n* 代码完全开源.\n\n#### <svg style=\"font-size: 22px; position: relative; top: 2px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mingmenjuan\"></use></svg> <span style=\"color: #f04a62\">about me</span>\n\n梦境里的宇宙，星光飘忽，命运蠕行其中,突破光锥联结万物. 我亦是梦中一束幻影，在黑暗背景内的一颗行星上，转瞬即逝.\n\n```\n,　　　　　　　　　.　 .　　　　　　　　.          ✦\n　　　　　。　　✦　　　　　　　　　　　　　　　ﾟ　　　　　　　　　。\n　　.　　　　　　　　.　　　:star:　　.　　　　　*　　　　。　　.　\n　 　　　　　　。　　　　　　　　　　　　　　　　　ﾟ　　　.　　　　　　　　　.\n,　　　　　　　　　.　 .　　　　　　　　.\n　　　　　。　　　　　　　　　　　　　　　　 *　　　ﾟ　　　　　　　　　。\n```\n\n>[info] 愿能在清醒的世界中找到自己的价值\n\n\nxxx\n```\n( (　`´︶´¯`︶´`´︶︶´*★ 　^v^　 ┊❅　 °☆ . 　 ☆ :. ☆ 　　\n　 ) ) 　⦅‖ ͇͇ ͇͇▃▇͇͇͌̿̿⌂͇͇▌..*　★　　☆ . 　 ★ ^v^ 　°　❅ 　☆ . 　 ★ ^v^ 　°　\n_̅̏̏̏̏̋̋̏̏▅̅̏̏̏̋̏ ╱◥███████╲　ˆ...^v^ ˆˆ︵.︵...^v^︵¸ ❅　 ° ╱◥◣ 　☆ . *　★　*　\n╱◥◣ ◥████◣▓∩▓∩║ 　 MERRY CHRISTMAS &　│∩ │◥███◣ ╱◥███◣\n│╱◥█◣║∩∩∩ ║╲◥███╲ 　 HAPPY NEW YEAR ╱◥◣ ◥████◣▓∩▓│∩\n││∩│ ▓ ║∏ 田║▓ 田田 ▓ ∩║ ii--ii--ii ⏃☃⏃ ii--ii--ii--ii │╱◥█◣║∩∩∩ ║◥█\n☸๑۩๑๑۩๑๑۩๑๑۩๑๑۩๑๑۩๑๑۩„„„„„„„„☸„„„„„„„☸„„„„„„„☸๑۩๑๑۩๑๑۩๑๑۩๑๑۩๑๑۩๑๑۩\n```\nxxx\n\n最开始 kicoe 这个名字是很久以前注册域名时随便取的，读音应该类似 *kikyo* ? 总之是个没啥实际意义从虚无中诞生的词. 也可以叫我月棱(moonprism)<span class=\"mask\" style=\"padding: 0px 1px\">酱</span>. 博主是社会底层、重度游戏废人、手残、baka，还能贴下列标签:\n\n`二次元` ☆～（ゝ。∂)\n`infp` 调停者...也许该派咱去调停俄乌关系(｡ŏ_ŏ)\n`阿卡林` 透明属性MAX\n`社恐` 才不是高冷，请让我偷偷注视着你们就好\n`严重的天真分子` ?\n`crew` crew真的是\n`解谜游戏爱好者` (逃课==解谜)&&(碎片化叙事==解谜要素)，是这样的没错\n\n<div style=\"height: 6px\"></div>\n\nxxx\n* PSN: [aquan445](https://psnprofiles.com/aquan445)\n* NS: SW-7753-9648-8871\nxxx\n\n>[fav] 最喜欢的游戏\n>> 黑暗之魂\n>> 血源诅咒\n>> 塞尔达传说：旷野之息\n>> 只狼：影逝二度\n>> 怪物猎人：世界\n>> 女神异闻录5\n>> 泰拉瑞亚\n>> 火焰纹章：回声\n>> 特鲁尼克大冒险2：不可思议的迷宫\n>> 艾尔登法环\n\n<blockquote class=\"ab-icon\">\n	<p>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://github.com/moonprism\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-github_pre\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://space.bilibili.com/3851114\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-bilibili-\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://book.douban.com/people/moonprism/\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-douban\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://music.163.com/#/user/home?id=76180165\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-wangyiyunyinle\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://steamcommunity.com/id/bumoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-steam1\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://t.me/kicoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-feiji\"></use>\n			</svg>\n		</a>\n	</p>\n</blockquote>\n\n<div style=\"height:3px\"></div>',	'2020-02-05 23:38:36',	'2023-10-18 14:35:16',	NULL);

-- GRANT ALL ON blog.* TO blog@localhost IDENTIFIED BY 'blog';
