# kicoe-blog

[kicoe.com](https://www.kicoe.com) 站点的所有源码

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

> 在docker目录下构建镜像的Dockerfile中有很多注释代码可能帮助解锁网络环境

首先安装好 `go` `npm` `gulp` 环境，可以通过docker镜像，不过我这里依赖了本地环境。

```shell
# Linux
make build
# Win
./build

docker-compose up -d
```

### 开发模式

不用先配置write中的`prod.ini`，只需要先把其他docker容器跑起来，再运行`go run main.go --env dev`启动后台api，进入web目录运行`npm run dev`启动后台界面。具体配置在各自的dev配置中相关联。

read页面通过gulp监听并压缩，在read页面下运行`gulp serve`就可以监听所有js/css的变化并刷新浏览器了

运行build将会编译go与前端程序并放置到对应的docker目录下，之后只要 `docker-compose up` 运行。

### composer

PHP就不copy进容器了
```shell
docker-compose exec read bash
$composer install
```

### 上线

先把各product配置都写好。

```
docker/
read/
docker-compose.yml
```
在build结束后将以上三个文件打包。

服务器上直接`docker-compose up -d`

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