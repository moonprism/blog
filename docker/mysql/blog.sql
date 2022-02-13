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
(1,	'Link',	3,	'',	'',	'# <svg style=\"font-size:23px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-aixin1\"></use></svg>友人帐\n\n>  [![](https://gravatar.cat.net/avatar/7c6d3737a25a9ec47b5439ec123bd1df) kicoe  `🍥*白天打游戏的废人*`](https://www.kicoe.com/)\n\n### <svg style=\"font-size:20px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-RectangleCopy\"></use></svg> Collection\n\n>  [![](https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640) moonprism  ` 🍉*另一人格？*`](https://www.kicoe.com/)\n\n<h1 id=\"board\"><svg style=\"font-size:23px;position:relative;top:1px;margin-right:8px\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-liuyanban\"></use></svg>留言板</h1>',	'2020-02-05 23:38:36',	'2022-02-13 11:26:09',	NULL),
(2,	'About',	3,	'',	'',	'# *![](https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640)*\n\n#### <svg style=\"font-size: 19px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mao\"></use></svg> <span style=\"color: #5c6bc0\"> the blog</span>\n\n* 很久之前拥有过一个几十M内存的服务器.\n\n* 因为内存很低，所以当时只用PHP写了个简单的框架&博客系统.\n\n* 换新服务器后用Golang+Vue重写了后台，可以在网页上使用vim了.\n\n* <a href=\"/page/link/#board\">link</a>页面欢迎留言，等网站内容足够充实后可能会去主动申请些友链.(大概率还是继续当只小透明~)\n\n* 博客代码完全开源\n\n#### <svg style=\"font-size: 22px; position: relative; top: 2px;\" class=\"icon\" aria-hidden=\"true\"><use xlink:href=\"#icon-mingmenjuan\"></use></svg> <span style=\"color: #f04a62\">about me</span>\n\n重度游戏废人，「只狼：影逝二度」攻略中，化身修罗吧！\n\n```\n__  __     __     ______     ______     ______    \n/\\ \\/ /    /\\ \\   /\\  ___\\   /\\  __ \\   /\\  ___\\ \n\\ \\  _\"-.  \\ \\ \\  \\ \\ \\____  \\ \\ \\/\\ \\  \\ \\  __\\   \n \\ \\_\\ \\_\\  \\ \\_\\  \\ \\_____\\  \\ \\_____\\  \\ \\_____\\ \n  \\/_/\\/_/   \\/_/   \\/_____/   \\/_____/   \\/_____/ .com\n \n```\n<div style=\"float: right; position: relative; bottom: 42px; right: 30px;\">⭐️</div>\n\n>[info] 有死之荣 无生之辱\n\n<blockquote class=\"ab-icon\">\n	<p>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://github.com/moonprism\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-github_pre\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://space.bilibili.com/3851114\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-bilibili-\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://book.douban.com/people/moonprism/\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-douban\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://music.163.com/#/user/home?id=76180165\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-wangyiyunyinle\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://steamcommunity.com/id/bumoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-steam1\"></use>\n			</svg>\n		</a>\n		<a target=\"_blank\" rel=\"noopener\" href=\"https://t.me/kicoe\">\n			<svg class=\"icon\" aria-hidden=\"true\">\n				<use xlink:href=\"#icon-feiji\"></use>\n			</svg>\n		</a>\n	</p>\n</blockquote>\n\n<div style=\"height:3px\"></div>',	'2020-02-05 23:38:36',	'2022-02-13 11:57:21',	NULL);

-- GRANT ALL ON blog.* TO blog@localhost IDENTIFIED BY 'blog';
