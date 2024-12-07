## 配置

设置 `[Github] ImagesApi` 为空，博客将使用阿里云 OSS 作为图床。

```toml
[Github]
  ImagesApi = ""
```

接下来配置 OSS：

```toml
[OSS]
  Region = "cn-shanghai" # 地域名
  Bucket = "kicoe-blog"  # 创建的 Bucket 名
  AccessKeyId = ""
  AccessKeySecret = ""
  RoleArn = ""
```

博客管理后台使用 [browser.js](https://help.aliyun.com/zh/oss/developer-reference/browser-js) SDK ，上传文件无需经过自己的服务器。首先浏览器向博客后端接口请求临时访问凭证，后端使用 STS SDK 调用阿里云 AssumeRole 接口，获取凭证返回给浏览器，最后浏览器通过这个临时凭证上传文件。

OSS 配置最后和用户相关的三项可以参考阿里云文档：[创建 RAM 用户](https://help.aliyun.com/zh/ram/user-guide/create-a-ram-user)

::: tip
创建的用户需要具有 `AliyunSTSAssumeRoleAccess` (调用 STS 服务 AssumeRole 接口) 的权限策略。

`RoleArn` 配置项在阿里云后台 RAM 访问控制 - 身份管理 - 角色 中创建具有 `AliyunOSSFullAccess` (管理 OSS) 权限的虚拟用户，详情页里可直接复制其 ARN。
:::
