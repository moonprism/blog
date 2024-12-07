请参阅 [Makefile](https://github.com/moonprism/blog/blob/sven/Makefile).

1. `go mod tidy` 和 `npm install`

    ```sh
    $ make init
    ```
2. 编译 linux darwin windows 软件包到 ./dist/

    ```sh
    $ make build
    ```
    
本地开发服务器跨域：
```nginx
  server {
    listen 2999;
    server_name localhost;

    location / {
      add_header 'Access-Control-Allow-Origin' 'http://localhost:5173';
      add_header 'Access-Control-Allow_Credentials' 'true';
      add_header 'Access-Control-Allow-Headers' 'Authorization,Accept,Origin,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Content-Range,Range';
      add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS,PUT,DELETE,PATCH';

      if ($request_method = 'OPTIONS') {
        return 200;
      }

      proxy_pass http://127.0.0.1:3000;
    }
  }
```