package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
)

type (
	AuthConf struct {
		SigningKey	string	`ini:"signing_key"`
	}

	AppConf struct {
		EnableCors	bool	`ini:"enable_cors"`
		EnableStatic	bool	`ini:"enable_static"`
		Port	string	`ini:"port"`
	}

	MysqlConf struct {
		Host	string	`ini:"host"`
		Port	string	`ini:"port"`
		User	string	`ini:"user"`
		Password	string	`ini:"password"`
		DataBase	string	`ini:"database"`
	}

	ElasticConf struct {
		Host	string	`ini:"host"`
	}

	OssConf struct {
		Endpoint	string	`ini:"endpoint"`
		AccessKeyId	string	`ini:"access_key_id"`
		AccessKeySecret	string	`ini:"access_key_secret"`
		BucketName	string	`ini:"bucket_name"`
	}

	RedisConf struct {
		Host	string	`ini:"host"`
		Port	string	`ini:"port"`
	}

	SMTPConf struct {
		User	string	`ini:"user"`
		Port	string	`ini:"port"`
		Pass	string	`ini:"pass"`
		Host	string	`ini:"host"`
	}
)

var (
	Auth = new(AuthConf)
	Mysql = new(MysqlConf)
	App = new(AppConf)
	Elastic = new(ElasticConf)
	Oss = new(OssConf)
	Redis = new(RedisConf)
	SMTP = new(SMTPConf)
	sectionMap = map[string]interface{}{
		"auth":	Auth,
		"mysql":	Mysql,
		"app":	App,
		"elastic": Elastic,
		"oss":	Oss,
		"redis": Redis,
		"smtp": SMTP,
	}

	once sync.Once
)

func InitConfig(env string) {
	once.Do(func() {
		cfg, err := ini.ShadowLoad(fmt.Sprintf("config/%s.ini", env))
		if err != nil {
			panic(err)
		}

		for section, config := range sectionMap {
			err = cfg.Section(section).MapTo(config)
			if err != nil {
				panic(err)
			}
		}
	})
}
