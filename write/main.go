package main

import (
	"flag"
	"log"
	"net"
	"os"

	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/protodata"
	"git.kicoe.com/blog/write/service"
	"git.kicoe.com/blog/write/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"git.kicoe.com/blog/write/config"
	"git.kicoe.com/blog/write/database"
	_ "git.kicoe.com/blog/write/docs"
	"git.kicoe.com/blog/write/middlewares"
	"git.kicoe.com/blog/write/router"
	"git.kicoe.com/blog/write/workers"
	"github.com/gookit/color"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Kicoe Blog API
// @version 1.0
// @description This is blog api test server. jwt token prefix: Bearer

// @BasePath /api/v1
var (
	env string
	re  bool
)

func init() {
	flag.StringVar(&env, "env", "dev", "environment for server:[dev|test|prod]")
	flag.BoolVar(&re, "reindex", false, "")
	flag.Parse()
}

func main() {

	// init
	config.InitConfig(env)
	database.InitMysqlEngine()
	utils.InitRiot()

	if len(os.Args) > 1 && os.Args[1] == "reindex" {
		// 重建查询索引
		reindex()
		return
	} else if re {
		reindex()
	}

	// todo check serve health

	go workers.RunSendEmail()

	app := router.Routers()
	app.HideBanner = true

	if env == "prod" {
		app.Debug = false
		app.Use(middleware.Recover())
		os.MkdirAll("log", os.ModePerm)
		f, err := os.OpenFile("log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: f,
		}))
	} else {
		app.Use(middlewares.Test)
		app.Use(middleware.Logger())
		app.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	go func() {
		lis, err := net.Listen("tcp", ":"+config.RPC.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		reflection.Register(s)
		protodata.RegisterCodeServer(s, &grpcS{})
		color.Printf("⇨ grpc server started on [::]<cyan>:%s</>\n", config.RPC.Port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	app.Logger.Fatal(app.Start(":" + config.App.Port))
}

// grpc server

type grpcS struct {
	protodata.CodeServer
}

func (s *grpcS) Search(ctx context.Context, request *protodata.SearchRequest) (*protodata.SearchResponse, error) {
	codes, _, err := service.SearchDoc(request.Text, 1, 30)
	if err != nil {
		return nil, err
	}
	result := &protodata.SearchResponse{}
	for _, code := range codes {
		result.Data = append(result.Data, &protodata.CodeDetail{
			Id:      code.ID,
			Title:   code.Description,
			Lang:    code.Lang,
			Tags:    code.Tags,
			Content: code.Content,
		})
	}
	return result, nil
}

func reindex() {
	codes := make([]*model.Code, 0)
	database.MysqlXEngine.Table("code").Find(&codes)
	for _, code := range codes {
		service.UpsertDoc(code)
		color.Red.Print(code.ID, " ")
		color.Green.Print(" ", code.Description, " ")
		color.BgMagenta.Print(" ", code.Lang, " ")
		color.Cyan.Println(" [", code.Tags, "] ")
		// color.Println(code.Content)
	}
	color.Primary.Tips("success reindex")
}
