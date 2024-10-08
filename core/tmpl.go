package core

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
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

func (tm *tmplManager) Register(name string, file string) (err error) {
	funcMap := template.FuncMap{
		"unsafeHTML": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
	tm.tmplData[name], err = template.New(path.Base(file)).Funcs(funcMap).ParseFiles(file)
	return
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
			err = tm.Register(nameWithoutExt, filepath.Join(dir, entry.Name()))
			if err != nil {
				return err
			}
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
