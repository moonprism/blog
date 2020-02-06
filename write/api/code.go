package api

import (
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/service"
	"git.kicoe.com/blog/write/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type codeListResponse struct {
	Pagination	*utils.Pagination `json:"pagination"`
	Data	[]*model.CodeMeta          `json:"data"`
}
// @Summary get code list
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param page query int true "current page"
// @Success 200 {object} codeListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /code [get]
func CodeList(c echo.Context) error {
	var (
		page int
		pageSize = 10
		err error
	)

	page, err = strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}

	codes, pagination, err := service.GetPaginateCodes(page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &codeListResponse{
		Pagination: pagination,
		Data: codes,
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
	codeUp := new(service.CodeUpdateBody)

	if err := c.Bind(codeUp); err != nil {
		return err
	}

	if err := service.CreateCode(codeUp); err != nil {
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

	code := &service.CodeUpdateBody{
		CodeMeta: new(model.CodeMeta),
		Content:     "",
	}

	if err := c.Bind(code); err != nil {
		return err
	}

	if err := service.UpdateCode(id, code); err != nil {
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

	if err := service.DeleteCode(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary show code
// @Tags Code
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "code id"
// @Success 200 {object} model.CodeDetail
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /code/{id} [get]
func CodeDetail(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil
	}

	code, err := service.GetCodeDetail(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, code)
}
