package aliyun

import (
	"codearena/pkg/config"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
	Bucket *oss.Bucket
)

func init() {
	var err error
	aliyun := config.GetConfig().Aliyun
	client, err = oss.New(aliyun.Endpoint, aliyun.AccessKeyID, aliyun.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	Bucket, err = client.Bucket(aliyun.BucketName)
	if err != nil {
		panic(err)
	}
}

func GetObjectUrl(objKey string) string {
	aliyun := config.GetConfig().Aliyun
	return fmt.Sprintf("https://%s.%s/%s", aliyun.BucketName, aliyun.Endpoint, objKey)
}
