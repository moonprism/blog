package tag

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/db"
	"git.kicoe.com/blog/write/modules/err/errors"
)

func GetInfos() (tags []*models.TagInfo, err error) {
	tags = make([]*models.TagInfo, 0)
	tagMetas, err := models.FetchTagMetas()
	if err != nil {
		return
	}

	for _, tagMeta := range tagMetas {
		tagCount, err := models.CountTagReference(tagMeta.ID)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &models.TagInfo{
			TagMeta: tagMeta,
			Count:   tagCount,
		})
	}
	return
}

func Create(tag *models.TagMeta) (err error) {
	affected, err := models.InsertTag(&models.Tag{
		TagMeta: tag,
	})

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func Update(id int64, tag *models.TagMeta) (err error) {
	affected, err := models.UpdateTag(id, &models.Tag{
		TagMeta: tag,
	}, nil)

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func Delete(id int64) (err error) {
	session := db.MysqlXEngine.NewSession()
	defer session.Close()

	affected, err := session.Id(id).Delete(new(models.Tag))

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		session.Rollback()
		return
	}
	if err != nil {
		session.Rollback()
		return
	}

	_, err = session.Where("tag_id = ?", id).Delete(new(models.ArticleTag))
	if err != nil {
		session.Rollback()
		return
	}
	session.Commit()
	return
}
