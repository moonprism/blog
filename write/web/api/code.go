package api

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/services/code"
	"git.kicoe.com/blog/write/modules/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CodeListResponse struct {
	Pagination *utils.Pagination `json:"pagination"`
	Data       []*models.CodeMeta `json:"data"`
}

// @Summary get code list
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param page query int true "current page"
// @Success 200 {object} CodeListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /code [get]
func CodeList(c echo.Context) error {
	var (
		page     = 1
		pageSize = 10
		err      error
	)

	page, err = strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}

	codes, pagination, err := code.GetList(page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &CodeListResponse{
		Pagination: pagination,
		Data:       codes,
	})
}

type CodeDetailListResponse struct {
	Pagination *utils.Pagination `json:"pagination"`
	Data       []*models.Code     `json:"data"`
}

// @Summary search by text
// @Tags Search
// @Produce  json
// @Param text query string true "search text"
// @Success 200 {object} CodeDetailListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /search/code [get]
func SearchCode(c echo.Context) error {
	var (
		page       = 1
		pageSize   = 30
		err        error
		searchText string
	)

	searchText = c.QueryParam("text")

	codes, pagination, err := code.SearchDoc(searchText, page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &CodeDetailListResponse{
		Pagination: pagination,
		Data:       codes,
	})
}

// @Summary create code
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param codeMeta body service.CodeUpdateBody true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /code [post]
func CreateCode(c echo.Context) error {
	codeUp := new(code.UpdateBody)

	if err := c.Bind(codeUp); err != nil {
		return err
	}

	if err := code.Create(codeUp); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary update code
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "code id"
// @Param codeUpdateBody body service.CodeUpdateBody true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /code/{id} [put]
func UpdateCode(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	co := &code.UpdateBody{
		CodeMeta: new(models.CodeMeta),
		Content:  "",
	}

	if err := c.Bind(co); err != nil {
		return err
	}

	if err := code.Update(id, co); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary delete code
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "code id"
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /code/{id} [delete]
func DeleteCode(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := code.Delete(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary show code
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "code id"
// @Success 200 {object} models.CodeDetail
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /code/{id} [get]
func CodeDetail(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil
	}

	code, err := code.GetDetail(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, code)
}
