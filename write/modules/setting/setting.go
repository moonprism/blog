package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Auth struct {
		SigningKey string `ini:"signing_key"`
	}

	App = struct {
		Debug   bool   `ini:"debug"`
		EnableStatic bool   `ini:"enable_static"`
		Port         string `ini:"port"`
	}{
		Debug: true,
		EnableStatic: true,
		Port: "80",
	}

	Mysql = struct {
		Host     string `ini:"host"`
		Port     string `ini:"port"`
		User     string `ini:"user"`
		Password string `ini:"password"`
		DataBase string `ini:"database"`
	}{
		Host: "127.0.0.1",
		Port: "3306",
	}

	Oss struct {
		Endpoint        string `ini:"endpoint"`
		AccessKeyId     string `ini:"access_key_id"`
		AccessKeySecret string `ini:"access_key_secret"`
		BucketName      string `ini:"bucket_name"`
	}

	Redis = struct {
		Host string `ini:"host"`
		Port string `ini:"port"`
	}{
		Host: "127.0.0.1",
		Port: "6379",
	}

	SMTP struct {
		User string `ini:"user"`
		Port string `ini:"port"`
		Pass string `ini:"pass"`
		Host string `ini:"host"`
		From string `ini:"from"`
	}

	RPC struct {
		Port string `ini:"port"`
	}

	sectionMap = map[string]interface{}{
		"auth":    &Auth,
		"mysql":   &Mysql,
		"app":     &App,
		"oss":     &Oss,
		"redis":   &Redis,
		"smtp":    &SMTP,
		"rpc":     &RPC,
	}
)

func Init(env string) {
	cfg, err := ini.ShadowLoad(fmt.Sprintf("options/%s.ini", env))
	if err != nil {
		panic(err)
	}

	for section, setting := range sectionMap {
		err = cfg.Section(section).MapTo(setting)
		if err != nil {
			panic(err)
		}
	}
}
