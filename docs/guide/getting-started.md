# 快速开始

## 安装与启动

在 [release](https://github.com/moonprism/blog/releases) 页面下载您当前系统对应的软件包并解压缩：

```
.
├─ app.toml
├─ dict/
│  ├─ *.utf8
│  └─ libsimple.so
├─ www/
│  ├─ 404/ ...
│  └─ v/ ...
└─ main
```

::: tip
可执行文件 `./main` 在 windows 系统下为 `.\main.exe`
:::

1. 初始化数据库：

    ```sh
    $ ./main db migrate
    ```

2. 设置用户名与密码：

    ```sh
    $ ./main admin passwd
    ```

3. 启动服务：

    ```sh
    $ ./main serve
    ```

默认服务器地址是 `http://localhost:3000`，在浏览器访问 URL 应该就能看到相应的页面了。

此时博客因为没有数据所以首页内容一片空白，访问 `http://localhost:3000/404/login` 输入之前设置的帐号和密码登录后台添加文章吧。

::: tip
在文章编辑器中使用 `/gi` 命令可以快速插入 github 风格的一些 markdown 附加语法哦
```md
> [!NOTE]
> Useful information that users should know, even when skimming content.

> [!TIP]
> Helpful advice for doing things better or more easily.

> [!IMPORTANT]
> Key information users need to know to achieve their goal.

> [!WARNING]
> Urgent info that needs immediate user attention to avoid problems.

> [!CAUTION]
> Advises about risks or negative outcomes of certain actions.
...
```
:::

## 配置文件

配置文件(`./app.toml`)用于设置博客的一些基础参数，其中比较重要的有:

```toml
IsDev = true
JwtSecret = "z8xgbmykvNViZkhwZnc1Tr/JnKNqNfbBThPYxiVBLgE="

[Server]
  Addr = ":3000"

[Database]
  Driver = "sqlite"
  Source = "test.db"

[FTS]
  Source = "test_fts.db"
```

- `IsDev` 开启情况下，程序将会通过以下两种方式主动代理前端静态文件：
    1. 如果在编译时开启了 `-tags embed` 选项，将会从嵌入的 go embed 中读取内容(windows)
    2. 否则将这些静态资源访问直接映射到当前程序运行目录的 `www` 文件夹下(linux、darwin)
- `JwtSecret` 应该是一个随机的、高熵字符串，建议至少 256 位(32 字节)，shell 生成：
  ```sh
  openssl rand -base64 32
  ```
- `[Server]` 设置服务器地址
- `[Database]` 数据库目前支持 SQLite 和 MySQL(MariaDB)
  
  所以默认配置下程序无需任何外部服务即可运行，当然想用 MySQL 也可以修改配置为：
  ```toml
  Driver = "mysql"
  Source = "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4"
  ```
- `[FTS]` 全文搜索数据库文件

## 下一步

- 接下来就是我为什么要写这个文档的核心：[图床配置](/guide/attachment-github-api)
  ::: warning
  博客可以使用 Github 和阿里云OSS 作为图床，有点麻烦，
  怕以后自己都忘了
  :::
- 想要了解如何使用 nginx 反向代理这个程序，使其运行在遥远的服务器上，请参见 [部署](/guide/deploy)。
- 最后，如果您闲得无聊，也可以查阅本项目的 [开发指南](/guide/development)