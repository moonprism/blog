package routers

import (
	"git.kicoe.com/blog/write/modules/err"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/web/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "git.kicoe.com/blog/write/docs"
)

func Routers() *echo.Echo {
	e := echo.New()

	if setting.App.Debug {
		e.Use(middleware.CORS())
	}
	// custom err handle
	e.HTTPErrorHandler = err.CustomHTTPErrorHandler

	// jwt middleware config
	jwtMiddle := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:   []byte(setting.Auth.SigningKey),
		Claims:       &jwt.StandardClaims{},
		ErrorHandler: err.CustomJwtErrorHandler,
	})

	if setting.App.EnableStatic {
		e.Static("/static", "static")
	}

	// router
	v1 := e.Group("/api/v1")
	{
		authG := v1.Group("/auth")
		{
			authG.POST("/login", api.Login)
			authG.GET("/captcha", api.Captcha)
		}

		casG := v1.Group("/cas")
		{
			casG.GET("/key", api.CasKey, jwtMiddle)
			casG.GET("/auth", api.CasAuth)
		}

		articleG := v1.Group("/article", jwtMiddle)
		{
			articleG.GET("", api.ArticleList)
			articleG.POST("", api.CreateArticle)
			articleG.PUT("/:id", api.UpdateArticle)
			articleG.DELETE("/:id", api.DeleteArticle)
			articleG.GET("/:id", api.ArticleDetail)
		}

		tagG := v1.Group("/tag", jwtMiddle)
		{
			tagG.GET("", api.TagList)
			tagG.POST("", api.CreateTag)
			tagG.PUT("/:id", api.UpdateTag)
			tagG.DELETE("/:id", api.DeleteTag)
		}

		codeG := v1.Group("/code", jwtMiddle)
		{
			codeG.GET("", api.CodeList)
			codeG.POST("", api.CreateCode)
			codeG.GET("/:id", api.CodeDetail)
			codeG.PUT("/:id", api.UpdateCode)
			codeG.DELETE("/:id", api.DeleteCode)
		}

		fileG := v1.Group("/file", jwtMiddle)
		{
			fileG.POST("/image", api.UploadImage)
			fileG.GET("/image", api.FileList)
			fileG.DELETE("/image/:id", api.DeleteFile)
		}

		sysG := v1.Group("/system/info")
		{
			sysG.GET("", api.GetSystemInfo)
			sysG.PUT("/theme", api.UpdateSystemTheme, jwtMiddle)
			sysG.PUT("/notice", api.UpdateSystemNotice, jwtMiddle)
		}

		commentG := v1.Group("/comment", jwtMiddle)
		{
			commentG.GET("", api.CommentList)
			commentG.DELETE("/:id", api.DeleteComment)
		}

		linkG := v1.Group("/link", jwtMiddle)
		{
			linkG.GET("", api.LinkList)
			linkG.PUT("/:id", api.UpdateLink)
		}

		searchG := v1.Group("/search")
		{
			searchG.GET("/code", api.SearchCode)
		}
	}

	if setting.App.Debug {
        e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return e
}
