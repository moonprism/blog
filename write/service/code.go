package service

import (
	"context"
	"git.kicoe.com/blog/write/errors"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"strconv"
)

func GetPaginateCodes(page, pageSize int) (codes []*model.CodeMeta, pagination *utils.Pagination, err error) {
	var (
		offset = (page-1)*pageSize
		whereField = "deleted_at is null"
	)

	codes, err = model.FetchCodeMetas(offset, pageSize, whereField)
	count, err := model.CountCode(whereField)
	pagination = utils.GeneratePagination(page, pageSize, count)
	return
}

func GetCodeDetail(id int64) (codeDetail *model.CodeDetail, err error) {
	var (
		has bool
	)

	codeDetail = new(model.CodeDetail)
	codeDetail.Code, has, err = model.FetchCode(id)
	if err != nil {
		return
	}
	if !has {
		err = errors.ServiceResourceNotFoundError
	}
	return
}

// ------------- es　随便写下 ---------------
type ElasticCode struct {
	Description	string	`json:"description"`
	Lang	string	`json:"lang"`
	Tags	string	`json:"tags"`
	Content	string	`json:"content"`
}

const ElasticCodeMapping = `
{
	"mappings": {
		"elastic-code-type": {
			"properties": {
				"description": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_max_word"
				},
				"lang": {
					"type": "keyword"
				},
				"tags": {
					"type": "string",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_max_word"
				}
			}
		}
	}
}
`
const (
	ElasticCodeIndex = "elastic-code-index"
	ElasticCodeType = "elastic-code-type"
)

func CreateElasticCode(ctx context.Context, esCode *ElasticCode, id string) (err error) {
	client := utils.NewESClient()
	exists, err := client.IndexExists(ElasticCodeIndex).Do(ctx)
	if err != nil {
		return
	}
	if !exists {
		createIndex, err := client.CreateIndex(ElasticCodeIndex).BodyString(ElasticCodeMapping).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			return errors.NewServiceError(10, "create failed")
		}
	}
	_, err = client.Index().
		Index(ElasticCodeIndex).
		Type(ElasticCodeType).
		Id(id).
		BodyJson(esCode).
		Do(ctx)
	return
}

func UpdateElasticCode(ctx context.Context, esCode *ElasticCode, id string) (err error) {
	client := utils.NewESClient()
	_, err = client.Update().
		Index(ElasticCodeIndex).
		Type(ElasticCodeType).
		Id(id).
		Doc(map[string]interface{}{
			"description": esCode.Description,
			"lang":	esCode.Lang,
			"content": esCode.Content,
			"tags": esCode.Tags,
		}).
		Do(ctx)
	return
}

func DeleteElasticCode(ctx context.Context, id string) (err error) {
	client := utils.NewESClient()
	_, err = client.Delete().
		Index(ElasticCodeIndex).
		Type(ElasticCodeType).
		Id(id).
		Do(ctx)
	return
}

type CodeUpdateBody struct {
	*model.CodeMeta
	Content	string	`json:"content"`
}
func CreateCode(codeUp *CodeUpdateBody) (err error) {
	code := &model.Code{
		CodeMeta:	codeUp.CodeMeta,
		Content:	codeUp.Content,
	}

	affected, err := model.InsertCode(code)
	if err != nil {
		return
	}

	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = CreateElasticCode(context.Background(), &ElasticCode{
		Description: code.Description,
		Lang:        code.Lang,
		Tags:        code.Tags,
		Content:     code.Content,
	}, strconv.FormatInt(code.ID, 10))
	return
}

func UpdateCode(id int64, codeUp *CodeUpdateBody) (err error) {
	code := &model.Code{
		CodeMeta:  codeUp.CodeMeta,
		Content:   codeUp.Content,
	}

	affected, err := model.UpdateCode(id, code, nil)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}

	err = UpdateElasticCode(context.Background(), &ElasticCode{
		Description: code.Description,
		Lang:        code.Lang,
		Tags:        code.Tags,
		Content:     code.Content,
	}, strconv.FormatInt(code.ID, 10))
	return
}

func DeleteCode(id int64) (err error) {
	affected, err := model.DeleteCode(id)
	if err != nil {
		return
	}
	if affected == 0 {
		err = errors.ServiceResourceNotFoundError
		return
	}
	err = DeleteElasticCode(context.Background(), strconv.FormatInt(id, 10))
	return
}
