package service

import (
	"bytes"
	"strconv"
	"strings"

	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
)

func GetPaginateCodes(page, pageSize int) (codes []*model.CodeMeta, pagination *utils.Pagination, err error) {
	var (
		offset = (page-1)*pageSize
		whereField = "deleted_at is null"
	)

	codes, err = model.FetchCodeMetas(offset, pageSize, whereField)
	count, err := model.CountCode(whereField)
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetCodeDetail(id int64) (codeDetail *model.CodeDetail, err error) {
	var (
		has bool
	)

	codeDetail = new(model.CodeDetail)
	codeDetail.Code, has, err = model.FetchCode(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func UpsertDoc(code *model.Code){
	utils.InsertOrUpdateDoc(strconv.FormatInt(code.ID, 10), strings.Join([]string{
		code.Description,
		code.Lang,
		code.Tags,
		code.Content,
	}, "\n"))
}

func deleteDoc(id string) {
	utils.RemoveDoc(id)
}

type CodeUpdateBody struct {
	*model.CodeMeta
	Content	string	`json:"content"`
}
func CreateCode(codeUp *CodeUpdateBody) (err error) {
	code := &model.Code{
		CodeMeta:	codeUp.CodeMeta,
		Content:	codeUp.Content,
	}

	affected, err := model.InsertCode(code)
	if err != nil {
		return
	}

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	UpsertDoc(code)

	return
}

func UpdateCode(id int64, codeUp *CodeUpdateBody) (err error) {
	code := &model.Code{
		CodeMeta:  codeUp.CodeMeta,
		Content:   codeUp.Content,
	}

	affected, err := model.UpdateCode(id, code, nil)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	UpsertDoc(code)

	return
}

func DeleteCode(id int64) (err error) {
	affected, err := model.DeleteCode(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	deleteDoc(strconv.FormatInt(id, 10))

	return
}

func SearchDoc(text string, page, limit int) (codes []*model.Code, pagination *utils.Pagination, err error) {
	ids, result, count := utils.SearchDoc(text, (page-1)*limit, limit)
	pagination = utils.GeneratePagination(page, limit, int64(count))
	// codes, err = model.FetchCodesByIds(ids)
	if err != nil {
		return
	}
	for _, id := range ids {
		code := &model.Code{ CodeMeta: &model.CodeMeta{} }
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
