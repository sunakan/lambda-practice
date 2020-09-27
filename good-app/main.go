package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/guregu/dynamo"
	"good-app/repository"
	"good-app/service"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
		Endpoint: aws.String("http://dynamodb-local:8000/"),
	})
	if err != nil {
		fmt.Println("============================[ Error ]")
		fmt.Println(err)
		return
	}
	db      := dynamo.New(sess)
	repo    := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	fmt.Println("===================[ OK ]")
	fmt.Println(service)
}
