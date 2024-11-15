//go:build !embed

package ui

import (
	"embed"
	"errors"
	"io/fs"
)

// 默认不带admin前端的静态文件

//go:embed vanilla/dist/*
var vanillaDistFiles embed.FS

func GetVanillaEmbedFS() fs.FS {
	staticFs, _ := fs.Sub(vanillaDistFiles, "vanilla/dist")
	return staticFs
}

func ReadAdminDistFile(name string) ([]byte, error) {
	return nil, errors.New("nombed")
}
