package utils

import (
	"git.kicoe.com/blog/write/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var (
	ossClient *oss.Client
	ossOnce sync.Once
)

func NewOssClient() *oss.Client {
	ossOnce.Do(func() {
		client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
		if err != nil {
			panic(err)
		}
		ossClient = client
	})
	return ossClient
}

func OssBlogBucket() (*oss.Bucket, error) {
	bucket, err := NewOssClient().Bucket(config.Oss.BucketName)
	if err != nil {
		return nil ,err
	}
	return bucket, nil
}

func InitOss() {
	NewOssClient()
}