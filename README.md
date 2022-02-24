博客系统, Golang + PHP + Vue ...

目录结构:

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
|    | -  | - prod.ini *
|    | - web
|    | -  | - app  管理界面 (vue)
|    | -  | -  | - .env.product *
| - docker-compose.yml *
(标*为可手动配置文件)
```

查看构建帮助:

```shell
make help
```

> 编译脚本 ./build.sh, 依赖 `docker-compose`、`vi`

