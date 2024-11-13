package core

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
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
		"unsafeAttr": func(str string) template.HTMLAttr {
			return template.HTMLAttr(str)
		},
		"unsafeURL": func(str string) template.URL {
			return template.URL(str)
		},
		"unsafeCSS": func(str string) template.CSS {
			return template.CSS(str)
		},
		"formatDate": func(timestamp uint) string {
			t := time.Unix(int64(timestamp), 0)
			dateStr := t.Format("2006/01/02<span>%s</span>Monday")
			// 获取星期几并转换为日文
			var weekdayStr string
			switch t.Weekday() {
			case time.Sunday:
				weekdayStr = "日"
			case time.Monday:
				weekdayStr = "月"
			case time.Tuesday:
				weekdayStr = "火"
			case time.Wednesday:
				weekdayStr = "水"
			case time.Thursday:
				weekdayStr = "木"
			case time.Friday:
				weekdayStr = "金"
			case time.Saturday:
				weekdayStr = "土"
			}
			return fmt.Sprintf(dateStr, weekdayStr)
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
