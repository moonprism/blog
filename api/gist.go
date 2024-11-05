package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"gorm.io/gorm"
)

func bindGistApi(app *core.App, r chi.Router) {
	api := gistApi{app}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(app.TokenAuth))
		r.Use(jwtauth.Authenticator(app.TokenAuth))
		r.Get("/", api.list)
		r.Post("/", api.create)
		r.Put("/{id}", api.update)
		r.Delete("/{id}", api.delete)
	})
}

type gistApi struct {
	*core.App
}

type gistList struct {
	Data       []*models.Gist    `json:"data"`
	Pagination models.Pagination `json:"pagination"`
}

func (api *gistApi) list(w http.ResponseWriter, r *http.Request) {
	model := api.O.Model(&models.Gist{}).Preload("GistOutput")
	page := 1
	pageSize := 10
	var count int64
	q := r.URL.Query().Get("q")
	if q != "" {
		var params models.SearchURLParams[string]
		err := json.Unmarshal([]byte(q), &params)
		core.P(err)
		page = params.Page + 1
		pageSize = params.PageSize
		// filter
		if params.FilterText != "" {
			model = model.Where(
				"`id` = ? OR `title` LIKE ? OR `lang` LIKE ? OR `content` LIKE ?",
				append(
					[]interface{}{params.FilterText},
					core.CreateSlice[interface{}](3, fmt.Sprintf("%%%s%%", params.FilterText))...,
				)...,
			)
		}
		for k, v := range params.FilterValues {
			if len(v) == 0 {
				continue
			}
			if k == "lang" {
				model = model.Where(fmt.Sprintf("%s IN ?", k), v)
			}
		}
		// sort
		if params.SortKey.ID != "" {
			order := "ASC"
			if params.SortKey.Order == "desc" {
				order = "DESC"
			}
			// inject
			model = model.Order(fmt.Sprintf("%s %s", params.SortKey.ID, order))
		}
	}
	err := model.Count(&count).Error
	core.P(err)
	var gists []*models.Gist
	err = model.Order("id DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&gists).
		Error
	core.P(err)
	json.NewEncoder(w).Encode(gistList{
		Data: gists,
		Pagination: models.Pagination{
			Count: int(count),
		},
	})
}

func (api *gistApi) create(w http.ResponseWriter, r *http.Request) {
	gist := new(models.Gist)
	err := json.NewDecoder(r.Body).Decode(gist)
	core.P(err)
	err = api.O.Create(gist).Error
	core.P(err)
	err = api.O.FtsInsert("gist", gist.ID, models.Gist2TextPoint(gist))
	core.P(err)
	api.JSON(w, gist)
}

func (api *gistApi) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	core.P(err)
	err = api.O.Transaction(func(tx *gorm.DB) error {
		gist := new(models.Gist)
		gist.ID = uint(id)

		if v, ok := data["output"]; ok {
			content := &models.GistOutput{
				GistID: gist.ID,
				HTML:   v.(map[string]interface{})["html"].(string),
			}
			err = tx.Save(content).Error
			if err != nil {
				return err
			}
			delete(data, "output")
		}

		if len(data) == 0 {
			return nil
		}
		err = tx.Model(gist).Updates(data).Error
		if err != nil {
			return err
		}

		err = tx.First(gist).Error
		if err != nil {
			return err
		}
		return api.O.FtsUpdate("gist", gist.ID, models.Gist2TextPoint(gist))
	})
	core.P(err)
	api.JSON(w, id)
}

func (api *gistApi) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	core.P(err)
	gist := new(models.Gist)
	gist.ID = uint(id)
	api.O.Delete(&gist)
	core.P(err)
	err = api.O.FtsDelete("gist", gist.ID)
	core.P(err)
	api.JSON(w, id)
}

func gistsPageRoute(app *core.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app.HTML(w, "gists", getAppSettings(app))
		core.P(err)
	}
}

const gistSearchResoultSplitS = "\n== 🌟 ==\n\n...\n"

func gistsSearchRoute(app *core.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = []*models.GistFts{}
		limit := 12
		keyword := r.URL.Query().Get("keyword")
		source := r.URL.Query().Get("source")
		if keyword != "" {
			if len(keyword) > 30 {
				return
			}
			var rows *sql.Rows
			var err error
			if source == "art" {
				rows, err = app.O.FtsSelect("article", keyword, limit, true)
				core.P(err)
				defer rows.Close()
				for rows.Next() {
					row := new(models.GistFts)
					var fulltext string
					// 1,2;3,4;
					var postext string
					rows.Scan(&row.ID, &postext, &fulltext)
					index := strings.IndexByte(fulltext, '\n')
					title := fulltext[:index]
					var sBuilder strings.Builder
					titlePos := len(title) + 1

					// 合并后的位置 相邻查询位置小于50就合并展示[1,2],[5,9] => [1, 9]
					var displayTextPos [][]int
					for i, parts := range strings.Split(postext, ";") {
						if len(parts) == 0 {
							break
						}
						pos, err := core.AtoISlice(strings.Split(parts, ","))
						core.P(err)
						// 对齐查询标记
						startPos := pos[0] + i*core.FtsSearchIdxLen
						if startPos < titlePos {
							continue
						}
						endPos := pos[1] + (i+1)*core.FtsSearchIdxLen
						posArrLen := len(displayTextPos)
						if posArrLen > 0 && displayTextPos[posArrLen-1][1]+50 > startPos {
							displayTextPos[posArrLen-1][1] = endPos
						} else {
							displayTextPos = append(displayTextPos, []int{startPos, endPos})
						}
					}
					lastPos := titlePos
					for _, pos := range displayTextPos {
						// 将开始标记往前推50字符再往前寻找一个换行符
						startPos := max(pos[0]-50, titlePos)
						for i := startPos; i >= titlePos; i-- {
							if fulltext[i:i+2] == "\n\n" || i == titlePos {
								startPos = i
								break
							}
						}
						// 如果开始标记小于上一个结束标记，合并两标记内容
						if startPos <= lastPos {
							startPos = lastPos
						} else {
							sBuilder.WriteString(gistSearchResoultSplitS)
						}
						// 将结束标记往后推50字符再往后寻找一个换行
						endPos := min(pos[1]+50, len(fulltext))
						for i := endPos; i < len(fulltext)+1; i++ {
							endPos = i
							if fulltext[i-2:i] == "\n\n" {
								break
							}
						}
						sBuilder.WriteString(fulltext[startPos:endPos])
						lastPos = endPos
					}
					if lastPos < len(fulltext) {
						sBuilder.WriteString(gistSearchResoultSplitS)
					}
					data = append(data, &models.GistFts{
						ID:      row.ID,
						Title:   title,
						Lang:    "md",
						Content: sBuilder.String(),
					})
				}
			} else {
				rows, err = app.O.FtsSelect("gist", keyword, limit, false)
				core.P(err)
				defer rows.Close()
				for rows.Next() {
					row := new(models.GistFts)
					var fulltext string
					rows.Scan(&row.ID, &fulltext)
					data = append(data, models.Text2Gist(row.ID, &fulltext))
				}
			}
		} else {
			var err error
			page := 1
			pageParam := r.URL.Query().Get("page")
			if pageParam != "" {
				page, err = strconv.Atoi(pageParam)
				core.P(err)
			}
			ids := []uint{}
			idsParam := r.URL.Query().Get("ids")
			if idsParam != "" {
				for _, s := range strings.Split(idsParam, ",") {
					// 转换每个字符串为 uint
					num, err := strconv.ParseUint(s, 10, 0) // 10 是基数，0 表示使用 uint 类型的默认位数
					core.P(err)
					ids = append(ids, uint(num))
				}
			}
			var gists []*models.Gist
			model := app.O.Model(&models.Gist{}).Preload("GistOutput").
				Order("id DESC")

			if len(ids) != 0 {
				page = 1
				model.Where("id in (?)", ids)
			}

			err = model.Offset((page - 1) * limit).
				Limit(limit).Find(&gists).
				Error
			core.P(err)
			for _, v := range gists {
				data = append(data, &models.GistFts{
					ID:      v.ID,
					Title:   v.Title,
					Lang:    v.Lang,
					Content: v.GistOutput.HTML,
				})
			}
		}
		app.JSON(w, data)
	}
}
