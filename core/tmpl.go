package core

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
	"time"
)

type tmplManager struct {
	tmpl *template.Template
}

var tmplFuncMap = template.FuncMap{
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

func NewTmplManager(name string) *tmplManager {
	return &tmplManager{
		tmpl: template.New(name).Funcs(tmplFuncMap),
	}
}

func (tm *tmplManager) RegistFS(f fs.FS) (err error) {
	entries, err := fs.ReadDir(f, ".")
	if err != nil {
		return err
	}
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".html" {
			files = append(files, entry.Name())
		}
	}
	tm.tmpl, err = tm.tmpl.ParseFS(f, files...)
	return
}

func (tm *tmplManager) Execute(tmplName string, w io.Writer, data any) error {
	return tm.tmpl.ExecuteTemplate(w, tmplName+".html", data)
}
