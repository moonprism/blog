package article

import (
	"git.kicoe.com/blog/write/modules/err/errors"
	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/utils"
	"strconv"
	"strings"
)

func GetList(page, pageSize int) (articles []*models.ArticleInfo, pagination *utils.Pagination, err error) {
	var (
		offset = (page - 1) * pageSize
		//articles = make([]*models.ArticleInfo, pageSize)
		articleIDs []string
		// 扩展用builder.Eq
		whereField = "deleted_at is null"
	)

	articleMetas, err := models.FetchArticleMetas(offset, pageSize, whereField)
	if err != nil {
		return
	}

	for _, meta := range articleMetas {
		articles = append(articles, &models.ArticleInfo{
			ArticleMeta: meta,
			Tags:        []*models.TagMeta{},
		})
		articleIDs = append(articleIDs, strconv.FormatInt(meta.ID, 10))
	}

	if articleIDs != nil {
		articleTagsMap, err := models.FetchArticleTagsMap(strings.Join(articleIDs, ","))
		if err != nil {
			return nil, nil, err
		}

		for _, article := range articles {
			if _, has := articleTagsMap[article.ID]; has {
				article.Tags = articleTagsMap[article.ID]
			}
		}
	}

	count, err := models.CountArticle(whereField)
	if err != nil {
		return
	}

	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetDetail(id int64) (articleDetail *models.ArticleDetail, err error) {
	var (
		has bool
	)

	articleDetail = new(models.ArticleDetail)
	articleDetail.Tags = []*models.TagMeta{}
	articleDetail.Article, has, err = models.FetchArticle(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
		return
	}

	articleTagsMap, err := models.FetchArticleTagsMap(strconv.FormatInt(articleDetail.ID, 10))
	if _, has := articleTagsMap[articleDetail.ID]; has {
		articleDetail.Tags = articleTagsMap[articleDetail.ID]
	}
	return
}

func Create(info *models.ArticleInfo) (err error) {
	article := &models.Article{
		ArticleMeta: info.ArticleMeta,
	}

	affected, err := models.InsertArticle(article)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = models.InsertArticleTags(article.ID, info.Tags)
	return
}

type ArticleUpdateBody struct {
	*models.ArticleInfo
	Content string `json:"content"`
}

func Update(id int64, articleUp *ArticleUpdateBody) (err error) {
	article := &models.Article{
		ArticleMeta: articleUp.ArticleMeta,
		Content:     articleUp.Content,
	}

	affected, err := models.UpdateArticle(id, article, []string{"image"})
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = models.UpdateArticleTags(id, articleUp.Tags)
	return
}

func Delete(id int64) (err error) {
	affected, err := models.DeleteArticle(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}
