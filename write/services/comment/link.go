package comment

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/utils"
)

func GetLinkList(page, pageSize int) (links []*models.Link, pagination *utils.Pagination, err error) {
	var (
		offset     = (page - 1) * pageSize
		whereField = "deleted_at is null"
	)

	links = make([]*models.Link, 0)
	err = db.MysqlXEngine.Table("link").Limit(pageSize, offset).Desc("id").Where(whereField).Find(&links)
	if err != nil {
		return
	}

	count, err := db.MysqlXEngine.Where(whereField).Count(new(models.Link))
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func UpdateLink(id int64, status int) (err error) {
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
