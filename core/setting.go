package core

import (
	"os"

	"github.com/BurntSushi/toml"
)

type setting struct {
	Server    settingServer
	Account   settingAccount
	Database  settingDatabase
	OSS       OssConfig
	Cache     cacheConfig
	JwtSecret string
}

type settingServer struct {
	Addr string
}

type settingDatabase struct {
	Driver string
	Source string
}

type settingAccount struct {
	Name string
	Pass string
}

type OssConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Bucket          string
	RoleArn         string
}

type cacheConfig struct {
	Addr string
}

var configPath = "./app.toml"

func NewSetting() (s setting, err error) {
	if _, err = os.Stat(configPath); err != nil {
		return
	}
	_, err = toml.DecodeFile(configPath, &s)
	return
}

func ReSetting(s *setting) error {
	file, err := os.OpenFile(configPath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return toml.NewEncoder(file).Encode(s)
}
