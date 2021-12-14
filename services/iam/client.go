package iam

import "github.com/marmotedu/medu-sdk-go/sdk"

const (
	defaultEndpoint = "127.0.0.1:9090"
	serviceName     = "iam"
)

// Client SDK 客户端，继承基础层的 Client。
type Client struct {
	sdk.Client
}

func NewClient(config *sdk.Config, credential *sdk.Credential) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		config = sdk.NewConfig().WithEndpoint(defaultEndpoint)
	}

	client.Init(serviceName).WithCredential(credential).WithConfig(config)
	return
}

// NewClientWithSecret 创建客户端：传入密钥、用户名、密码配置。
func NewClientWithSecret(secretID, secretKey string) (client *Client, err error) {
	client = &Client{}
	config := sdk.NewConfig().WithEndpoint(defaultEndpoint)
	client.Init(serviceName).WithSecret(secretID, secretKey).WithConfig(config)
	return
}
