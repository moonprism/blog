package code

import (
	"strings"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/se"
	"git.kicoe.com/blog/write/modules/utils"
)

func GetList(page, pageSize int) (codes []*models.CodeMeta, pagination *utils.Pagination, err error) {
	var (
		offset     = (page - 1) * pageSize
		whereField = "deleted_at is null"
	)

	codes, err = models.FetchCodeMetas(offset, pageSize, whereField)
	count, err := models.CountCode(whereField)
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetDetail(id int64) (codeDetail *models.CodeDetail, err error) {
	var (
		has bool
	)

	codeDetail = new(models.CodeDetail)
	codeDetail.Code, has, err = models.FetchCode(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

// func ToSearchDoc(code *models.Code) models.SearchCodeParams {
// 	return models.SearchCodeParams{
// 		Description:    code.Description,
// 		Lang:    code.Lang,
// 		Tags:    code.Tags,
// 		Content: code.Content,
// 	}
// }

func ToSearchDoc(code *models.Code) string {
	return strings.Join([]string{
		code.Description,
		code.Lang,
		code.Tags,
		code.Content,
	}, "\n\n\n\n") //...
}

func ParseFromDoc(string) {

}

type UpdateBody struct {
	*models.CodeMeta
	Content string `json:"content"`
}

func Create(codeUp *UpdateBody) (err error) {
	code := &models.Code{
		CodeMeta: codeUp.CodeMeta,
		Content:  codeUp.Content,
	}

	affected, err := models.InsertCode(code)
	if err != nil {
		return
	}

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	// todo transaction
	err = se.Engine.Insert("code", code.ID, ToSearchDoc(code))
	if err != nil {
		return
	}
	return
}

func Update(id int64, codeUp *UpdateBody) (err error) {
	code := &models.Code{
		CodeMeta: codeUp.CodeMeta,
		Content:  codeUp.Content,
	}

	affected, err := models.UpdateCode(id, code, []string{"tags"})
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	code.ID = id
	err = se.Engine.Update("code", code.ID, ToSearchDoc(code))
	if err != nil {
		return
	}
	return
}

func Delete(id int64) (err error) {
	err = se.Engine.Delete("code", id)
	if err != nil {
		return
	}

	affected, err := models.DeleteCode(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	return
}

func SearchDoc(text string, page, limit int) (codes []*models.Code, pagination *utils.Pagination, err error) {
	result, count, err := se.Engine.Search("code", text, (page-1)*limit, limit)
	pagination = utils.GeneratePagination(page, limit, count)
	for _, r := range result {
		rs := strings.SplitN(string(r.([]byte)), "\n\n\n\n", 4)
		code := &models.Code{CodeMeta: &models.CodeMeta{}}
		code.Description = rs[0]
		code.Lang = rs[1]
		code.Tags = rs[2]
		code.Content = rs[3]
		codes = append(codes, code)
	}
	return
}
