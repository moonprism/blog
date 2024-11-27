package core

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type settings struct {
	IsDev     bool
	JwtSecret string
	Server    serverSet
	Database  databaseSet
	FTS       ftsSet
	OSS       ossSet
	Cache     cacheSet
	System    SystemSet
	Github    githubSet
}

type serverSet struct {
	Addr string
}

type databaseSet struct {
	Driver string
	Source string
}

type ftsSet struct {
	Source string
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

type githubSet struct {
	ImagesApi string
}

type SystemSet struct {
	AttachmentCDN    string        `json:"attachmentCDN"`
	TokenExpiryHours time.Duration `json:"tokenExpiryHours"`
}

var configPath = "./app.toml"

func NewSettings() (s settings, err error) {
	if _, err = os.Stat(configPath); err != nil {
		return
	}
	_, err = toml.DecodeFile(configPath, &s)
	return
}
