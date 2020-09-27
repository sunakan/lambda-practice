package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db dynamo.DB
}

// DB作成
// 作成したDBを直接repositoryに
func NewUserRepository() (*UserRepository, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
		Endpoint: aws.String("http://dynamodb-local:8000/"),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	db := dynamo.New(sess)
	return &UserRepository{*db}, nil
}
