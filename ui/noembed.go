//go:build !embed

package ui

import (
	"embed"
	"errors"
	"io/fs"
)

// 默认不带前端静态文件

//go:embed vanilla/dist/tmpl/*
var vanillaTmplFiles embed.FS

func GetVanillaTmplFS() fs.FS {
	staticFs, _ := fs.Sub(vanillaTmplFiles, "vanilla/dist/tmpl")
	return staticFs
}

func GetVanillaDistFS() fs.FS {
	return nil
}

func ReadAdminDistFile(name string) ([]byte, error) {
	return nil, errors.New("nombed")
}
