package router

import (
	"git.kicoe.com/blog/write/api"
	"git.kicoe.com/blog/write/config"
	"git.kicoe.com/blog/write/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Routers() *echo.Echo {
	e := echo.New()

	if config.App.EnableCors {
		e.Use(middleware.CORS())
	}
	// custom err handle
	e.HTTPErrorHandler = errors.CustomHTTPErrorHandler

	// jwt middleware config
	jwtMiddle := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:   []byte(config.Auth.SigningKey),
		Claims:       &jwt.StandardClaims{},
		ErrorHandler: errors.CustomJwtErrorHandler,
	})

	if config.App.EnableStatic {
		e.Static("/", "web/dist")
		e.Static("/static", "static")
	} else {
		e.Static("/", "web")
	}

	// router
	v1 := e.Group("/api/v1")
	{
		v1.GET("/cas/auth", func (c echo.Context) error {
			if c.QueryParam("key") == config.CAS.Key {
				return c.String(http.StatusOK, "success")
			}
			return c.String(http.StatusUnauthorized, "401")
		})

		authG := v1.Group("/auth")
		{
			authG.POST("/login", api.Login)
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

		sysG := v1.Group("/setting", jwtMiddle)
		{
			sysG.GET("", api.SettingDetail)
			sysG.PUT("", api.UpdateSetting)
		}

		commentG := v1.Group("/comment", jwtMiddle)
		{
			commentG.GET("", api.CommentList)
			commentG.DELETE("/:id", api.DeleteComment)
		}
	}

	return e
}
