package api

import (
	"net/http"
	"strconv"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/utils"
	"git.kicoe.com/blog/write/services/article"
	"github.com/labstack/echo/v4"
)

type ArticleListResponse struct {
	Pagination *utils.Pagination     `json:"pagination"`
	Data       []*models.ArticleInfo `json:"data"`
}

// @Summary get article list
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param page query int true "current page"
// @Success 200 {object} ArticleListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /article [get]
func ArticleList(c echo.Context) error {
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

	articles, pagination, err := article.GetList(page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &ArticleListResponse{
		Pagination: pagination,
		Data:       articles,
	})
}

// @Summary create article
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param articleInfo body model.ArticleInfo true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError
// @Router /article [post]
func CreateArticle(c echo.Context) error {
	articleInfo := new(models.ArticleInfo)

	if err := c.Bind(articleInfo); err != nil {
		return err
	}

	if err := article.Create(articleInfo); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary update article
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "article id"
// @Param articleUpdateBody body service.ArticleUpdateBody true " "
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /article/{id} [put]
func UpdateArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	art := &article.ArticleUpdateBody{
		ArticleInfo: new(models.ArticleInfo),
		Content:     "",
	}

	if err := c.Bind(art); err != nil {
		return err
	}

	if err := article.Update(id, art); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary delete article
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "article id"
// @Success 200 {string} string "ok"
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /article/{id} [delete]
func DeleteArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := article.Delete(id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "ok")
}

// @Summary show article
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "article id"
// @Success 200 {object} model.ArticleDetail
// @Failure 403 {object} errors.HTTPError "Unauthenticated"
// @Failure 404 {object} errors.HTTPError "NotFound"
// @Router /article/{id} [get]
func ArticleDetail(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil
	}

	article, err := article.GetDetail(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, article)
}
