package oss

import (
	"git.kicoe.com/blog/write/modules/setting"
	aliyunoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var (
	ossClient *aliyunoss.Client
	ossOnce   sync.Once
)

func NewOssClient() *aliyunoss.Client {
	ossOnce.Do(func() {
		client, err := aliyunoss.New(setting.Oss.Endpoint, setting.Oss.AccessKeyId, setting.Oss.AccessKeySecret)
		if err != nil {
			panic(err)
		}
		ossClient = client
	})
	return ossClient
}

func OssBlogBucket() (*aliyunoss.Bucket, error) {
	bucket, err := NewOssClient().Bucket(setting.Oss.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
