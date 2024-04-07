package core

import (
	"os"

	"github.com/BurntSushi/toml"
)

type setting struct {
	Server   SettingServer
	Database SettingDatabase
}

type SettingServer struct {
	Addr string
}

type SettingDatabase struct {
	Driver string
	Source string
}

func NewSetting(f string) (s setting, err error) {
	if _, err = os.Stat(f); err != nil {
		return
	}
	_, err = toml.DecodeFile(f, &s)
	return
}
