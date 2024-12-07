在部署前先关闭开发模式：

```toml
isDev = false
```

此时程序将不会代理 `www` 目录下的静态文件，可以通过 `nginx` 服务器来代理它们：

```nginx
server {
    listen 80;
    server_name localhost;
    root /home/Download/blog-linux/www;
    location = / {
      try_files $uri @proxy;
    }
    location / {
      try_files $uri $uri/ @proxy;
    }
    location ^~ /404/ {
      try_files $uri /404/404.html =404;
    }
    location @proxy {
      proxy_pass http://127.0.0.1:3000;
    }
}
```

::: info
Windows 版本的软件包并没有 `www` 目录，而是通过 go embed 打包成了单个可执行文件。
:::