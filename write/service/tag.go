package service

import (
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
)

func GetTagInfos() (tags []*model.TagInfo, err error) {
	tags = make([]*model.TagInfo, 0)
	tagMetas, err := model.FetchTagMetas()
	if err != nil {
		return
	}

	for _, tagMeta := range tagMetas {
		tagCount, err := model.CountTagReference(tagMeta.ID)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &model.TagInfo{
			TagMeta: tagMeta,
			Count:	tagCount,
		})
	}
	return
}

func CreateTag(tag *model.TagMeta) (err error) {
	affected, err := model.InsertTag(&model.Tag{
		TagMeta:     tag,
	})

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func UpdateTag(id int64, tag *model.TagMeta) (err error) {
	affected, err := model.UpdateTag(id, &model.Tag{
		TagMeta:     tag,
	}, nil)

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

func DeleteTag(id int64) (err error) {
	affected, err := model.DeleteTag(id)

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	if err != nil {
		return
	}

	err = model.DeleteAllArticleTags(id)
	return
}
