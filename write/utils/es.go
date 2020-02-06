package utils

import (
	"git.kicoe.com/blog/write/config"
	"gopkg.in/olivere/elastic.v5"
	"sync"
)

var (
	elasticClient *elastic.Client
	esOnce sync.Once
)

func NewESClient() *elastic.Client {
	esOnce.Do(func() {
		client, err := elastic.NewClient(
			elastic.SetURL("http://"+config.Elastic.Host+":9200/"),
		)
		if err != nil {
			panic(err)
		}
		elasticClient = client
	})
	return elasticClient
}

func InitEs() {
	NewESClient()
}