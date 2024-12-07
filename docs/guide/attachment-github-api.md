## 配置

在配置文件中如果 `[Github] ImagesApi` 不为空，博客将开启 github 图床：

```toml
[Github]
  ImagesApi = "https://api.github.com/repos/{用户名}/{仓库名}/contents/"
```

::: warning
其实只是开放了一个从该 api 同步文件列表的渠道，才短短几十行代码。什么自动化钩子啊，文件裁剪呀，请交给第三方工具吧！
:::

此时需要将 `[System] AttachmentCDN` 修改为对应的地址：

```toml
[System]
  AttachmentCDN = "https://raw.githubusercontent.com/{用户名}/{仓库名}/master/"
```

如果你的图片是放在仓库的文件夹内，可以在这两个地址后添加路径：

- `https://api.github.com/repos/{用户名}/{仓库名}/contents/dir1/dir2`
- `https://raw.githubusercontent.com/{用户名}/{仓库名}/master/dir1/dir2`

## 后台管理

开启 Github 图床后，在后台文件管理页面的上传按钮将变成 `Sync from github` 同步按钮：

![sync form github](https://raw.githubusercontent.com/moonprism/cdn/master/blog_doc/guide-attachment-github-api-1.png)