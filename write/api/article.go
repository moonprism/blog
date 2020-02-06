package api

import (
	_ "git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/service"
	"git.kicoe.com/blog/write/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type articleListResponse struct {
	Pagination	*utils.Pagination `json:"pagination"`
	Data	[]*model.ArticleInfo          `json:"data"`
}
// @Summary get article list
// @Tags Article
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param page query int true "current page"
// @Success 200 {object} articleListResponse
// @Failure 403 {object} errors.HTTPError
// @Router /article [get]
func ArticleList(c echo.Context) error {
	var (
		page int
		pageSize = 10
		err error
	)

	page, err = strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}

	articles, pagination, err := service.GetPaginateArticles(page, pageSize)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &articleListResponse{
		Pagination: pagination,
		Data: articles,
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
	articleInfo := new(model.ArticleInfo)

	if err := c.Bind(articleInfo); err != nil {
		return err
	}

	if err := service.CreateArticleByInfo(articleInfo); err != nil {
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

	article := &service.ArticleUpdateBody{
		ArticleInfo: new(model.ArticleInfo),
		Content:     "",
		HTML:        "",
	}

	if err := c.Bind(article); err != nil {
		return err
	}

	if err := service.UpdateArticle(id, article); err != nil {
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

	if err := service.DeleteArticle(id); err != nil {
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

	article, err := service.GetArticleDetail(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, article)
}
