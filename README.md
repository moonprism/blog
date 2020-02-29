# kicoe-blog

博客站点的所有源码

## 目录

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
|    | -  | -  | - | - markdown.js * 
| - write     后台管理 (go)
|    | - api
|    | - config
|    | -  | - dev.ini *
|    | -  | - prod.ini *
|    | - web     管理界面 (vue)
|    | -  | - src 
|    | -  | - .env.development *
|    | -  | - .env.product *
| - docker-compose.yml
```
标 * 的为可配置文件

## 运行

> 在docker目录下构建镜像的Dockerfile中有些注释代码可帮助解锁网络环境

首先在本地环境安装好 `go` `npm` `gulp`

```shell
# Linux
make build
# Win
./build

docker-compose up -d
```

### 开发

不用先配置write中的`prod.ini`，只需要先把其他docker容器跑起来。

* 进入write目录运行`go run main.go --env dev`启动后台api
* 进入write/web目录运行`npm run dev`启动后台web界面，监听相关代码修改自动刷新`127.0.0.1:8081`
* 进入read目录运行`gulp serve`启动前台web界面，监听相关代码修改自动刷新`127.0.0.1:2000`

```
// 切换dev分支，设置提交模板
git config commit.template .git-commit-template.txt
```

### 上线

先把各product配置文件配置好，运行build。在build结束后将以下三个文件打包上传服务器。
```
docker/
read/
docker-compose.yml
```

服务器上运行`docker-compose up -d`

> `.github/workflows`目录下配置了持续集成，借助github actions自动编译上线

## 依赖

* [kicoephp](https://github.com/moonprism/kicoephp-src)
* [markdown.js](https://github.com/moonprism/markdown.js)

---

* [echo](https://github.com/labstack/echo)
* [jwt-go](https://github.com/dgrijalva/jwt-go)
* [xorm](https://github.com/go-xorm/xorm)
* [swagger](https://github.com/go-swagger/go-swagger)
* [go elastic](gopkg.in/olivere/elastic.v5)
* [gulp](https://github.com/gulpjs/gulp)
* [vue](https://github.com/vuejs/vue)
* [axios](https://github.com/axios/axios)
* [element-ui](https://github.com/ElemeFE/element)
* [codemirror](https://github.com/vuejs/vue)
* [vue route](https://github.com/vuejs/vue-router)
* [vuex](https://github.com/vuejs/vuex)
* [vue-advanced-cropper](https://github.com/Norserium/vue-advanced-cropper)
* [vue-waterfall-plugin](https://github.com/heikaimu/vue-waterfall-plugin)