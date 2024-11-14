package ui

import (
	"embed"
	"io/fs"
)

//go:embed vanilla/dist/*
var vanillaDistFiles embed.FS

func GetVanillaEmbedFS() fs.FS {
	staticFs, _ := fs.Sub(vanillaDistFiles, "vanilla/dist")
	return staticFs
}
