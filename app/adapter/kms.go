package adapter

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"os"
)

// AWSKmsClient AWS SDKから KMS を利用して暗号化・復号化する
type AWSKmsClient struct {
	Client *kms.KMS
	KeyID  string
}

// NewAWSKmsClient AWSKmsClient インスタンスを生成
func NewAWSKmsClient() *AWSKmsClient {
	client := kms.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion("ap-northeast-1"))
	return &AWSKmsClient{
		Client: client,
		KeyID:  os.Getenv("KMS_KEY_ID"),
	}
}
