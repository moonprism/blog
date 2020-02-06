package api

import (
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

// @Summary upload image
// @Tags File
// @Accept  multipart/form-data
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param image formData file true " "
// @Success 200 {object} model.FileMeta
// @Failure 403 {object} errors.HTTPError
// @Router /file/image [post]
func UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileModel, err := service.CreateFile(src)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fileModel.FileMeta)
}

type fileListResponse struct {
	Data	[]*model.File	`json:"data"`
}
// @Summary get file list
// @Tags File
// @Produce  json
// @param Authorization header string true "Authorization"
// @param max_time query string false "current time yyyy-MM-dd HH:mm:ss"
// @Param size query int false "fetch list length"
// @Success 200 {object} fileListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /file/image [get]
func FileList(c echo.Context) error {
	var (
		size = 50
		err error
		maxTimeFormat string
	)

	timeFormat := c.QueryParam("max_time")
	if timeFormat != "" {
		maxTime, err := time.Parse("2006-01-02 15:04:05", timeFormat)
		if err != nil {
			return errors.HTTPErrorBadRequest
		}
		maxTimeFormat = maxTime.Format("2006-01-02 15:04:05")
	}


	sizeParam := c.QueryParam("size")
	if sizeParam != "" {
		size, err = strconv.Atoi(c.QueryParam("size"))
		if err != nil {
			return err
		}
	}

	files, err := service.GetFilesWaterfall(size, maxTimeFormat)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &fileListResponse{Data:files})
}


// @Summary delete file
// @Tags File
// @Produce  json
// @param Authorization header string true "Authorization"
// @param id path int true "file id"
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /file/image/{id} [delete]
func DeleteFile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := service.DeleteFile(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}