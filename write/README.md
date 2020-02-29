# write

write是博客的后台管理系统，基于golang和vue。前端目录位于web/

## config

config目录下拷贝一份dev.ini重命名为prod.ini

指定prod配置运行：
```
go run main.go prod
```

## swagger

```go
app.GET("/swagger/*", echoSwagger.WrapHandler)
```