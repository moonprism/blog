package service

import (
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"mime/multipart"
	"os"
)

func CreateFile(file multipart.File) (fileModel *model.File, err error) {
	key := utils.RandStringBytes(20) + ".jpg"
	filePath := "./static/" + key
	if err = utils.FileCopy(file, filePath); err != nil {
		return
	}

	fileModel = &model.File{
		FileMeta:    &model.FileMeta{Key:key},
	}

	affected, err := model.InsertFile(fileModel)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	bucket, err := utils.OssBlogBucket()
	if err != nil {
		return
	}

	err = bucket.PutObjectFromFile(key, filePath)
	return
}

func GetFilesWaterfall(pageSize int, maxTime string) (files []*model.File, err error) {
	files, err = model.FetchFileByCreated(pageSize, maxTime)
	return
}

func DeleteFile(id int) (err error) {
	file, has, err := model.FetchFile(id)
	if err != nil {
		return err
	}

	if !has {
		return errors.ServiceResourceNotFoundError
	}

	err = os.Rename("./static/"+file.Key, "./static/DEL-"+file.Key)
	if err != nil {
		return
	}

	affected, err := model.DeleteFile(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	bucket, err := utils.OssBlogBucket()
	if err != nil {
		return
	}

	err = bucket.DeleteObject(file.Key)
	return
}
