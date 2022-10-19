package code

import (
	"strconv"

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

func ToDoc(code *models.Code) map[string]interface{} {
	return map[string]interface{}{
			"id":      code.ID,
			"desc":    code.Description,
			"lang":    code.Lang,
			"tags":    code.Tags,
			"content": code.Content,
	}
}

func ToDocs(code *models.Code) []map[string]interface{} {
	return []map[string]interface{}{ToDoc(code)}
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
	err = se.NewIndex("code").Insert(ToDocs(code))
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
	err = se.NewIndex("code").Update(ToDocs(code))
	if err != nil {
		return
	}
	return
}

func Delete(id int64) (err error) {
	err = se.NewIndex("code").Delete(strconv.FormatInt(id, 10))
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
	result, count := se.NewIndex("code").Search(text, (page-1)*limit, limit)
	pagination = utils.GeneratePagination(page, limit, count)
	// codes, err = model.FetchCodesByIds(ids)
	if err != nil {
		return
	}
	for _, re := range result {
		code := &models.Code{CodeMeta: &models.CodeMeta{}}
		code.ID, _ = strconv.ParseInt(re["id"].(string), 10, 64)
		code.Description = re["desc"].(string)
		code.Lang = re["lang"].(string)
		code.Tags = re["tags"].(string)
		code.Content = re["content"].(string)
		codes = append(codes, code)
	}
	return
}
