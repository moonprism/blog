package api

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/utils"
	"git.kicoe.com/blog/write/services/link"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type LinkListResponse struct {
	Pagination *utils.Pagination `json:"pagination"`
	Data       []*models.Link    `json:"data"`
}

// @Summary get link list
// @Tags Link
// @Produce json
// @Success 200 {object} LinkListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /link [get]
func LinkList(c echo.Context) error {
	var (
		page     = 1
		pageSize = 10
		err      error
	)

	if c.QueryParam("page") != "" {
		page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return err
		}
	}

	links, pagination, err := link.GetList(page, pageSize)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, LinkListResponse{
		Pagination: pagination,
		Data:       links,
	})
}

type LinkUpdateBody struct {
	Status int `json:"status"`
}

// @Summary update link apply
// @Tags Link
// @Produce json
// @Success 200 {object} models.Link
// @Failure 403 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /link/{id} [put]
func UpdateLink(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	l := &LinkUpdateBody{}
	if err := c.Bind(l); err != nil {
		return err
	}

	if err := link.Update(id, l.Status); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}
