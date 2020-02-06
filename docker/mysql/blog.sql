CREATE DATABASE `blog`;
USE `blog`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `password` char(32) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(64) NOT NULL,
    -- draft publish system ...
    `status` tinyint(3) NOT NULL,
    `image` varchar(64) NOT NULL DEFAULT '',
    `summary` varchar(500) NOT NULL DEFAULT '',
    `content` text NOT NULL,
    `html`  text NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `key` varchar(255) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

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
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `art_id` int(9) NOT NULL,
    `tag_id` int(9) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `art_id` (`art_id`),
    KEY `tag_id` (`tag_id`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `art_id` int(9) NOT NULL,
    `to_id` int(9) NOT NULL,
    `name` varchar(50) NOT NULL,
    `email` varchar(30) NOT NULL,
    `text` varchar(255) NOT NULL,
    `created_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`),
    KEY `art_id` (`art_id`),
    KEY `to_id` (`to_id`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

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
)ENGINE=MyISAM DEFAULT CHARSET=utf8;

INSERT INTO `article` (`title`, `status`, `image`, `summary`, `content`, `html`, `created_time`, `updated_time`, `deleted_at`)
VALUES ('Link', 3, '', '', '', '', now(), now(), NULL);

INSERT INTO `article` (`title`, `status`, `image`, `summary`, `content`, `html`, `created_time`, `updated_time`, `deleted_at`)
VALUES ('About', 3, '', '', '', '', now(), now(), NULL);

-- GRANT ALL ON blog.* TO blog@localhost IDENTIFIED BY 'blog';