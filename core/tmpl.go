package core

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type tmplManager struct {
	tmplData map[string]*template.Template
}

func NewTmplManager() *tmplManager {
	return &tmplManager{
		tmplData: make(map[string]*template.Template),
	}
}

func (tm *tmplManager) Register(name string, file string) {
	tm.tmplData[name] = template.Must(template.ParseFiles(file))
}

// RegisterFs 挂载目录中所有html文件
func (tm *tmplManager) RegisterDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".html" {
			fileName := entry.Name()
			nameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]
			fmt.Println("Regist tmpl:", nameWithoutExt)
			tm.Register(nameWithoutExt, filepath.Join(dir, entry.Name()))
		}
	}
	return nil
}

func (tm *tmplManager) Execute(tmplName string, wr io.Writer, data any) error {
	tmpl, exists := tm.tmplData[tmplName]
	if !exists {
		return fmt.Errorf("template {%s} not found", tmplName)
	}
	return tmpl.Execute(wr, data)
}
