## kicoe-blog

个人博客系统，PHP + Golang + Vue ...

#### 目录结构

```
| - data      数据
| - docker    各容器编译目录
| -  | - mysql
| -  | -  | - blog.sql
| - log       日志
| - read      前台展示 (php)
|    | - app
|    | -  | - config.php *
|    | - public
|    | -  | - static
|    | -  | -  | - js 
|    | -  | -  | - | - main.js * 
| - write     后台管理 (golang)
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

首先在本地环境安装好 `docker` `docker-compose` `go` `npm`

查看构建帮助:

```shell
make help
```

如果使用 tmux, 以下脚本可跑开发环境

```shell
tmux new-session -s blog-dev -d
tmux send-keys -t blog-dev 'make up' C-m
tmux send-keys -t blog-dev 'make read-dev' C-m

tmux split-window -v -t blog-dev
tmux send-key -t blog-dev 'make write-web-dev' C-m

tmux split-window -h -t blog-dev
tmux send-keys -t blog-dev 'cd write' C-m
tmux send-keys -t blog-dev 'make serve-reindex' C-m

tmux a -t blog-dev
```

> 在 docker 目录下构建镜像的 Dockerfile 中有些注释代码可帮助解锁网络环境

#### 设计

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/f-1.png)

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/f-3.png)

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/b-2.png)

---

* [kicoephp](https://github.com/moonprism/kicoephp-src)
* [markdown.js](https://github.com/moonprism/markdown.js)
