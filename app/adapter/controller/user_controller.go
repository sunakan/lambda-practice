package controller

import (
	"go-app/registry"
	"go-app/usecase"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// RequestPostUser PostUserのリクエスト
type RequestPostUser struct {
	Name  string `json:"user_name"`
	Email string `json:"email"`
}

// ユーザ新規作成
func PostUsers(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	// JSON形式から構造体に変換
	var req RequestPostUser
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}
	// 新規作成処理
	creator := registry.GetFactory().BuildCreateUser()
	res, err := creator.Execute(&usecase.CreateUserRequest {
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return Response500(err)
	}
	// 201レスポンス
	return Response201(res.GetUserID())
}
