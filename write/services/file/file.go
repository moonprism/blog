package file

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/modules/oss"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/modules/utils"
	"mime/multipart"
	"os"
)

func Create(file multipart.File) (m *models.File, err error) {
	key := string(utils.RandBytes(20)) + ".jpg"
	os.MkdirAll("static", os.ModePerm)
	filePath := "./static/" + key
	if err = utils.FileCopy(file, filePath); err != nil {
		return
	}

	m = &models.File{
		FileMeta: &models.FileMeta{Key: key},
	}

	affected, err := models.InsertFile(m)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	if !setting.App.EnableStatic {
		bucket, err := oss.OssBlogBucket()
		if err != nil {
			return nil, err
		}
		go func() {
			err = bucket.PutObjectFromFile(key, filePath)
			// todo chan接收err处理
		}()
	}
	return
}

func GetFilesWaterfall(pageSize int, maxTime string) (files []*models.File, err error) {
	files, err = models.FetchFileByCreated(pageSize, maxTime)
	return
}

func Delete(id int) (err error) {
	file, has, err := models.FetchFile(id)
	if err != nil {
		return err
	}

	if !has {
		return errors.ServiceResourceNotFoundError
	}

	realFile := "./static/" + file.Key

	if utils.FileExists(realFile) {
		err = os.Rename(realFile, "./static/DEL-"+file.Key)
		if err != nil {
			return
		}
	}

	affected, err := models.DeleteFile(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	if !setting.App.EnableStatic {
		bucket, err := oss.OssBlogBucket()
		if err != nil {
			return err
		}
		go func() {
			err = bucket.DeleteObject(file.Key)
		}()
	}

	return
}
