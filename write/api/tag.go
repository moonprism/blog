package api

import (
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// @Summary get tag list
// @Tags Tag
// @Produce  json
// @param Authorization header string true "Authorization"
// @Success 200 {object} model.TagList
// @Failure 403 {object} errors.HTTPError
// @Router /tag [get]
func TagList(c echo.Context) error {
	var (
		err error
	)

	tags, err := service.GetTagInfos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tags)
}


// @Summary create tag
// @Tags Tag
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param tagMeta body model.TagMeta true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /tag [post]
func CreateTag(c echo.Context) error {
	tagMeta := new(model.TagMeta)

	if err := c.Bind(tagMeta); err != nil {
		return err
	}

	if err := service.CreateTag(tagMeta); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}


// @Summary update tag
// @Tags Tag
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "tag id"
// @Param TagMeta body model.TagMeta true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /tag/{id} [put]
func UpdateTag(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	tag := new(model.TagMeta)
	if err := c.Bind(tag); err != nil {
		return err
	}

	if err := service.UpdateTag(id, tag); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary delete tag
// @Tags Tag
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "tag id"
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /tag/{id} [delete]
func DeleteTag(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := service.DeleteTag(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}


