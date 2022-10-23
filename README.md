博客系统

PHP + Golang

## Features

* 简单，快速
* markdown
* 设计清楚、可爱
* 轻量的的全文搜索
* 评论系统
* 基于 codemirror 的后台 vim 编辑器

## Directory

```
| - docker    各容器编译目录
| -  | - mysql
| -  | -  | - blog.sql
| - read      前台展示 (php)
|    | - app
|    | -  | - config.php *
|    | - public
|    | -  | - static
|    | -  | -   | - js 
|    | -  | -   | -  | - main.js * 
| - write     后台管理 (golang)
|    | - config
|    | -  | - prod.ini *
|    | - web
|    | -  | - app  管理界面 (vue)
|    | -  | -  | - .env.product *
| - docker-compose.yml
> *为可配置文件
```

## Usage

先安装好 `docker` `docker-compose`，查看构建帮助:

```shell
make help
```

## Tech Stack

* PHP
  - [kicoephp](https://github.com/moonprism/kicoephp-src)
* Golang
  - echo
  - riot
* JS
  - [markdown.js](https://github.com/moonprism/markdown.js)
  - vue
  - element-ui
  - gulp

## Preview

![](https://github.com/moonprism/cdn/blob/master/image/f-1.png?raw=true)
