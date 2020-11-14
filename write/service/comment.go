package service

import (
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"git.kicoe.com/blog/write/errors"
)

func GetPaginateComments(page, pageSize int) (comments []*model.Comment, pagination *utils.Pagination, err error) {
	var (
		offset = (page-1)*pageSize
		whereField = "deleted_at is null"
	)

	comments, err = model.FetchComments(offset, pageSize, whereField)
	count, err := model.CountComment(whereField)
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func DeleteComment(id int64) (err error) {
	affected, err := model.DeleteComment(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	return
}