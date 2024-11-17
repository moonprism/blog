//go:build embed

package ui

import (
	"embed"
	"io/fs"
)

// -tags embed 编译出单个可执行文件用于开发测试

//go:embed vanilla/dist/tmpl/*
var vanillaTmplFiles embed.FS

func GetVanillaTmplFS() fs.FS {
	staticFs, _ := fs.Sub(vanillaTmplFiles, "vanilla/dist/tmpl")
	return staticFs
}

//go:embed vanilla/dist/*
var vanillaDistFiles embed.FS

func GetVanillaDistFS() fs.FS {
	staticFs, _ := fs.Sub(vanillaDistFiles, "vanilla/dist")
	return staticFs
}

//go:embed admin/build/*
var adminDistFiles embed.FS

func ReadAdminDistFile(name string) ([]byte, error) {
	return adminDistFiles.ReadFile("admin/build/" + name)
}
