## kicoe-blog

简单的个人博客系统，PHP + Golang + Vue + ES ...

#### 目录结构

```
| - data      数据
| - docker    各容器的编译目录
| -  | - mysql
| -  | -  | - blog.sql
| - log       日志
| - read      前端展示 (php)
|    | - app
|    | -  | - config.php *
|    | - public
|    | -  | - static
|    | -  | -  | - js 
|    | -  | -  | - | - main.js * 
| - write     后台管理 (go)
|    | - api
|    | - config
|    | -  | - dev.ini *
|    | -  | - prod.ini *
|    | - web     管理界面 (vue)
|    | -  | - src 
|    | -  | - .env.development *
|    | -  | - .env.product *
| - docker-compose.yml *
```

> 标 * 的为可配置文件

#### 运行

首先在本地环境安装好 `docker` `docker-compose` `go` `npm`，运行：
```shell
make help
```
查看构建帮助

> 在docker目录下构建镜像的Dockerfile中有些注释代码可帮助解锁网络环境

#### 设计

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/f-1.png)

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/f-3.png)


![](https://raw.githubusercontent.com/moonprism/cdn/master/image/b-2.png)


---

* [kicoephp](https://github.com/moonprism/kicoephp-src)
* [markdown.js](https://github.com/moonprism/markdown.js)
