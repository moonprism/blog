package comment

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/utils"
)

func GetList(page, pageSize int) (comments []*models.Comment, pagination *utils.Pagination, err error) {
	var (
		offset     = (page - 1) * pageSize
		whereField = "deleted_at is null"
	)

	comments, err = models.FetchComments(offset, pageSize, whereField)
	count, err := models.CountComment(whereField)
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetDetail(id int64) (detail *models.Comment, err error) {
	var has bool

	detail = new(models.Comment)
	detail, has, err = models.FetchComment(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func Delete(id int64) (err error) {
	affected, err := models.DeleteComment(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	return
}
