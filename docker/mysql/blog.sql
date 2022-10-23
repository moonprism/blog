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
(1,	'Link',	3,	'',	'',	'# <svg style=\"font-size:23px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-aixin1\"></use></svg>友人帐\n\n@@@\n\n> [![](STzvbpDkfSTIXoHHysUR.jpg)Roki`:sparkles:*玩塔科夫去吧*`](http://blog.weekii.cn/)\n\n> [![](KVrKQpxLKOQSCUndrzrs.jpg)Wrath`:pill:*呐呐呐呐呐*`](https://wrath.cc/)\n\n> [![](xFPCVcPzNHbhRmoiWyad.jpg)Lionad`:rocket:*可爱捏*`](http://www.lionad.art/)\n\n> [![](ivvmJbFUjpOscOSeREMH.jpg)Innei`:evergreen_tree:*致虚极，守静笃*`](https://innei.ren/)\n\n> [![](https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/cc/blogvg.ico)生如夏花`:rabbit:*博客复活!*`](http://www.blog.vg/)\n\n@@@\n\n### <svg style=\"font-size:20px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-RectangleCopy\"></use></svg> Collection\n\n> [<img src=\"https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/gnDtqzYhLVVBIGJkYzVY.jpg\" style=\"border-radius:0;box-shadow:none\">致郁系`:cherry:*页面设计超好看*`](http://www.juroku.net/)\n\n<h1 id=\"board\"><svg style=\"font-size:25px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-pingguo1\"></use></svg>留言板</h1>\n\nxxx\n\n<div style=\"text-align:center\">\n	<canvas id=\"p5\" width=\"267\" height=\"62\"></canvas>\n</div>\n\n>[info] 名称: kicoe\'s Blog\n简介: 白天打游戏的废人\n链接: `https://kicoe.com`\n图标: <https://gravatar.cat.net/avatar/7c6d3737a25a9ec47b5439ec123bd1df>\n符号: :star:\n\n***\n\n---\n\nxxx',	'2020-02-05 23:38:36',	'2022-10-20 09:12:38',	NULL),
(2,	'About',	3,	'',	'',	'# *![](https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640)*\n\n#### <svg style=\"font-size: 19px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mao\"></use></svg> <span style=\"color: #5c6bc0\"> the blog</span>\n\n* 很久之前拥有过一个几十兆内存的服务器.\n\n* 因为内存很低，所以当时只用 PHP 写了个简单的框架&博客系统.\n\n* 换新服务器后用 Golang + Vue 重写了后台，可以在网页上使用 vim 了.\n\n* <a href=\"/link\">link</a>页面欢迎留言，等网站内容足够充实后可能会去主动申请些友链.(大概率还是继续当只小透明~)\n\n* 博客代码完全开源.\n\n#### <svg style=\"font-size: 22px; position: relative; top: 2px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mingmenjuan\"></use></svg> <span style=\"color: #f04a62\">about me</span>\n\n重度游戏废人，「血源诅咒」攻略中\n\n```\n,　　　　　　　　　.　 .　　　　　　　　.          ✦\n　　　　　。　　✦　　　　　　　　　　　　　　　ﾟ　　　　　　　　　。\n　　.　　　　　　　　.　　　:star:　　.　　　　　*　　　　。　　.　\n　 　　　　　　。　　　　　　　　　　　　　　　　　ﾟ　　　.　　　　　　　　　.\n,　　　　　　　　　.　 .　　　　　　　　.\n　　　　　。　　　　　　　　　　　　　　　　 *　　　ﾟ　　　　　　　　　。\n```\n\nkicoe 这个名字是以前注册域名时随便取的，读音应该类似 *kikyo* ? 总之是个没啥实际意义从虚无中诞生的词. kicoe 喜欢看书、漫画，还是个游戏爱好者，可贴以下标签:\n\n* `infp` 调停者...也许该派咱去调停俄乌关系?\n* `crew` 头像是虚拟主播湊あくあ，喜欢看她打游戏~~、整活~~\n* `阿卡林` 透明属性MAX\n* `社恐` 才不是高冷，请让我偷偷注视着你们就好\n\n<div style=\"height: 6px\"></div>\n\n>[fav] 最喜欢的游戏\n>> 黑暗之魂3\n>> 只狼：影逝二度\n>> 塞尔达传说：旷野之息\n>> 女神异闻录5\n>> 文明6\n>> 怪物猎人：世界\n>> 泰拉瑞亚\n>> 艾尔登法环\n>> 特鲁尼克大冒险：不可思议的迷宫\n>> 三国志英杰传\n\n<blockquote class=\"ab-icon\">\n	<p>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://github.com/moonprism\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-github_pre\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://space.bilibili.com/3851114\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-bilibili-\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://book.douban.com/people/moonprism/\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-douban\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://music.163.com/#/user/home?id=76180165\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-wangyiyunyinle\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://steamcommunity.com/id/bumoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-steam1\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://t.me/kicoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-feiji\"></use>\n			</svg>\n		</a>\n	</p>\n</blockquote>\n\n<div style=\"height:3px\"></div>',	'2020-02-05 23:38:36',	'2022-10-19 19:01:53',	NULL);

-- GRANT ALL ON blog.* TO blog@localhost IDENTIFIED BY 'blog';
