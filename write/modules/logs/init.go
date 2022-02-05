package logs

import (
	"os"
	"path/filepath"

	"git.kicoe.com/blog/write/email"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/web/middlewares"
	"github.com/sirupsen/logrus"
)

func Init() {
	email.SetLog(NewFileLog("log/email.log"))
	middlewares.SetLog(NewFileLog("log/web.log"))
}

func NewFileLog(fileName string) *logrus.Logger {
	logger := logrus.New()

	// 设置日志等级
	if setting.App.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}

	// 创建目录
	dirName := filepath.Dir(fileName)
	if err := os.MkdirAll(dirName, 0755); err != nil {
		logrus.Panicf("error create directory %v", err)
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Panicf("error opening file %v", err)
	}
	logger.Out = file

	return logger
}
