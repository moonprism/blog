package service

import (
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"strconv"
	"strings"
)

func GetPaginateArticles(page, pageSize int) (articles []*model.ArticleInfo, pagination *utils.Pagination, err error) {
	var (
		offset = (page-1)*pageSize
		//articles = make([]*model.ArticleInfo, pageSize)
		articleIDs []string
		// 扩展用builder.Eq
		whereField = "deleted_at is null"
	)

	articleMetas, err := model.FetchArticleMetas(offset, pageSize, whereField)
	if err != nil {
		return
	}

	for _, meta := range articleMetas {
		articles = append(articles, &model.ArticleInfo{
			ArticleMeta: meta,
			Tags:        []*model.TagMeta{},
		})
		articleIDs = append(articleIDs, strconv.FormatInt(meta.ID, 10))
	}

	if articleIDs != nil {
		articleTagsMap, err := model.FetchArticleTagsMap(strings.Join(articleIDs, ","))
		if err != nil {
			return nil, nil, err
		}

		for _, article := range articles {
			if _, has := articleTagsMap[article.ID]; has {
				article.Tags = articleTagsMap[article.ID]
			}
		}
	}

	count, err := model.CountArticle(whereField)
	if err != nil {
		return
	}

	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetArticleDetail(id int64) (articleDetail *model.ArticleDetail, err error) {
	var (
		has bool
	)

	articleDetail = new(model.ArticleDetail)
	articleDetail.Tags = []*model.TagMeta{}
	articleDetail.Article, has, err = model.FetchArticle(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
		return
	}

	articleTagsMap, err := model.FetchArticleTagsMap(strconv.FormatInt(articleDetail.ID, 10))
	if _, has := articleTagsMap[articleDetail.ID]; has {
		articleDetail.Tags = articleTagsMap[articleDetail.ID]
	}
	return
}

func CreateArticleByInfo(info *model.ArticleInfo) (err error) {
	article := &model.Article{
		ArticleMeta: info.ArticleMeta,
	}

	affected, err := model.InsertArticle(article)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = model.InsertArticleTags(article.ID, info.Tags)
	return
}

type ArticleUpdateBody struct {
	*model.ArticleInfo
	Content	string	`json:"content"`
}

func UpdateArticle(id int64, articleUp *ArticleUpdateBody) (err error) {
	article := &model.Article{
		ArticleMeta: articleUp.ArticleMeta,
		Content:     articleUp.Content,
	}

	affected, err := model.UpdateArticle(id, article, []string{"image"})
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = model.UpdateArticleTags(id, articleUp.Tags)
	return
}

func DeleteArticle(id int64) (err error) {
	affected, err := model.DeleteArticle(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
	}
	return
}