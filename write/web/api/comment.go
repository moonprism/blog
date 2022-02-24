package api

import (
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/utils"
	"git.kicoe.com/blog/write/services/comment"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Pagination *utils.Pagination `json:"pagination"`
	Data       []*models.Comment `json:"data"`
}

// @Summary get comment list
// @Tags Comment
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param page query int true "current page"
// @Success 200 {object} CommentListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /comment [get]
func CommentList(c echo.Context) error {
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

	comments, pagination, err := comment.GetList(page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &CommentListResponse{
		Pagination: pagination,
		Data:       comments,
	})
}

// @Summary delete comment
// @Tags Comment
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "comment id"
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /comment/{id} [delete]
func DeleteComment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := comment.Delete(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}
