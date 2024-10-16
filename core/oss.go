package core

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type oss struct {
	stsClient  *sts20150401.Client
	ExpSeconds int64
	RoleArn    string
}

func newOSS(ak string, sk string, region string, roleArn string) (o *oss, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(ak),
		AccessKeySecret: tea.String(sk),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Sts
	config.Endpoint = tea.String("sts." + region + ".aliyuncs.com")
	stsClient, err := sts20150401.NewClient(config)
	if err != nil {
		return
	}
	o = &oss{
		stsClient:  stsClient,
		ExpSeconds: 3306,
		RoleArn:    roleArn,
	}
	return
}

func (o *oss) GetAssumeRole() (body *sts20150401.AssumeRoleResponseBodyCredentials, err error) {
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		DurationSeconds: (*int64)(&o.ExpSeconds),
		RoleArn:         &o.RoleArn,
		RoleSessionName: tea.String("alice"),
	}
	runtime := &util.RuntimeOptions{}

	result, err := o.stsClient.AssumeRoleWithOptions(assumeRoleRequest, runtime)
	if err != nil {
		return
	}

	body = result.Body.Credentials
	return
}
