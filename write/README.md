## write

博客后台管理系统

* golang 1.16+

> 虽然 go1.16 还没发布，还是喜欢官方的静态资源打包方法...

查看构建帮助：
```
make help
```

和前台(php)交互:

* `grpc` 用于搜索, 现在全文检索已经从 elasticsearch 迁移到 go 的 roit, php 前台作为 grpc client
* `cas` 用于验证( jwt 无法同域...), 使用 vue 从接口获取 key 并跳转, golang 后台作为认证服务器
* `redis` 用于邮件推送队列等

前端目录位于 **/web**

---

![](https://raw.githubusercontent.com/moonprism/cdn/master/image/b-3.png)
