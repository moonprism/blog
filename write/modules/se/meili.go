package se

import (
	"git.kicoe.com/blog/write/modules/setting"
	"github.com/davecgh/go-spew/spew"
	"github.com/meilisearch/meilisearch-go"
)

var (
	client *meilisearch.Client
)

func Init() {
	client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://" + setting.Meili.Host + ":" + setting.Meili.Port,
		APIKey: "",
	})
}

type Index struct {
	index *meilisearch.Index
}

func NewIndex(index string) *Index {
	return &Index{
		index: client.Index(index),
	}
}

func (i *Index) Insert(docs []map[string]interface{}) error {
	_, err := i.index.AddDocuments(docs)
	if err != nil {
		return err
	}
	return nil
}

func (i *Index) Update(docs []map[string]interface{}) error {
	_, err := i.index.UpdateDocuments(docs)
	if err != nil {
		return err
	}
	return nil
}

func (i *Index) Delete(id string) error {
	_, err := i.index.DeleteDocument(id)
	if err != nil {
		return err
	}
	return nil
}

func (i *Index) Search(text string, offset, limit int) (result []map[string]interface{}, count int64) {
	searchRaw, err := i.index.Search(text, &meilisearch.SearchRequest{
		Limit:                 int64(limit),
		Offset:                int64(offset),
		AttributesToHighlight: []string{"*"},
	})
	spew.Dump(searchRaw)
	if err != nil {
		return
	}
	for _, r := range searchRaw.Hits {
		result = append(result, r.(map[string]interface{})["_formatted"].(map[string]interface{}))
	}
	count = searchRaw.NbHits
	return
}
