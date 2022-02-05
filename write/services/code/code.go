package code

import (
	"bytes"
	"strconv"
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

func convertParams(code *models.Code) (string, string) {
	return strconv.FormatInt(code.ID, 10), strings.Join([]string{
		code.Description,
		code.Lang,
		code.Tags,
		code.Content,
	}, "\n")
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
	se.InsertDoc(convertParams(code))

	return
}

func Update(id int64, codeUp *UpdateBody) (err error) {
	code := &models.Code{
		CodeMeta: codeUp.CodeMeta,
		Content:  codeUp.Content,
	}
	affected, err := models.DeleteCode(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	se.RemoveDoc(strconv.FormatInt(id, 10))

	affected, err = models.InsertCode(code)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	se.InsertDoc(convertParams(code))

	return
}

func Delete(id int64) (err error) {
	affected, err := models.DeleteCode(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	se.RemoveDoc(strconv.FormatInt(id, 10))

	return
}

func SearchDoc(text string, page, limit int) (codes []*models.Code, pagination *utils.Pagination, err error) {
	ids, result, count := se.SearchDoc(text, (page-1)*limit, limit)
	pagination = utils.GeneratePagination(page, limit, int64(count))
	// codes, err = model.FetchCodesByIds(ids)
	if err != nil {
		return
	}
	for _, id := range ids {
		code := &models.Code{CodeMeta: &models.CodeMeta{}}
		resultSlice := bytes.SplitN(result[id], []byte{'\n'}, 4)
		code.ID = id
		code.Description = string(resultSlice[0])
		code.Lang = string(resultSlice[1])
		code.Tags = string(resultSlice[2])
		if len(resultSlice) > 3 {
			code.Content = string(resultSlice[3])
		}
		codes = append(codes, code)
	}
	return
}

func Reindex() {
	for i := 0; ; i++ {
		codes, _, _ := GetList(i, 100)
		for _, code := range codes {
			se.InsertDoc(convertParams(&models.Code{CodeMeta:code}))
		}
		if len(codes) < 100 {
			break
		}
	}
}
