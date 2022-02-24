package link

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/utils"
)

const STATUS_ACCEPT = 1
const STATUS_REJECT = 2

func GetList(page, pageSize int) (links []*models.Link, pagination *utils.Pagination, err error) {
	var (
		offset     = (page - 1) * pageSize
		whereField = "deleted_at is null"
	)

	links = make([]*models.Link, 0)
	err = db.MysqlXEngine.Limit(pageSize, offset).Desc("id").Where(whereField).Find(&links)
	if err != nil {
		return
	}

	count, err := db.MysqlXEngine.Where(whereField).Count(new(models.Link))
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func FetchDetailByEmail(email string) (l *models.Link, has bool, err error) {
	l = new(models.Link)
	has, err = db.MysqlXEngine.Where("email = ?", email).Where("status = ?", STATUS_ACCEPT).Get(l)
	return
}

func Update(id int64, status int) (err error) {
	link := new(models.Link)
	link.Status = status
	affected, err := db.MysqlXEngine.Id(id).Update(link)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	return
}
