package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"git.kicoe.com/blog/write/modules/utils"
	"git.kicoe.com/blog/write/services/file"
	"github.com/urfave/cli/v2"
)

var CmdSync = cli.Command{
	Name:        "sync",
	Usage:       "sync local image from oss",
	Description: "",
	Action:      syncImage,
}

func syncImage(ctx *cli.Context) error {
	Init()
	files, err := file.GetFilesWaterfall(3000, "")
	if err != nil {
		panic(err)
	}
	// 本地备份用的，懒得弄配置文件了
	domainUrl := "https://kicoe-blog.oss-cn-shanghai.aliyuncs.com/"
	localDir := "../data/write/static/"//"/www/static/"
	for _, v := range files {
		fmt.Println(v.Key)
		fileName := localDir+v.Key
		if utils.FileExists(fileName) {
			continue
		}

		resp, err := http.Get(domainUrl+v.Key)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		out, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
