package core

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type setting struct {
	Server    serverSet
	Account   accountSet
	Database  databaseSet
	OSS       ossSet
	Cache     cacheSet
	System    SystemSet
	JwtSecret string
}

type serverSet struct {
	Addr string
}

type databaseSet struct {
	Driver string
	Source string
}

type accountSet struct {
	Name string
	Pass string
}

type ossSet struct {
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Bucket          string
	RoleArn         string
}

type cacheSet struct {
	Addr string
}

type SystemSet struct {
	Title            string        `json:"title"`
	AttachmentCDN    string        `json:"attachmentCDN"`
	TokenExpiryHours time.Duration `json:"tokenExpiryHours"`
	LastLoginTime    string        `json:"lastLoginTime"`
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
