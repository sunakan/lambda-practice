package registry

import (
	"go-app/adapter"
	"os"
)

// Envs 環境変数
// 暗号化やキャッシュも可能
type Envs struct {
	KMSClient *adapter.AWSKmsClient
	Cache     map[string]string
}

var envs *Envs

// NewEnvs Envs インスタンスを生成
func NewEnvs() *Envs {
	return &Envs {
		KMSClient: adapter.NewAWSKmsClient(),
		Cache:     make(map[string]string),
	}
}

// Env シングルトンを取得する
func Env() *Envs {
	if envs == nil {
		envs = NewEnvs()
	}
	return envs
}

func (c *Envs) env(key string) string {
	return os.Getenv(key)
}

func (c *Envs) DynamoLocalEndpoint() string {
	return c.env("DYNAMO_LOCAL_ENDPOINT")
}

func (c *Envs) DynamoTableName() string {
	return c.env("DYNAMO_TABLE_NAME")
}

func (c *Envs) DynamoPKName() string {
	return c.env("DYNAMO_PK_NAME")
}

func (c *Envs) DynamoSKName() string {
	return c.env("DYNAMO_SK_NAME")
}
