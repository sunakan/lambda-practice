package controller

import (
	"app4/registry"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// こんな感じのjsonを期待
// {"email":"test@example.com","user_name":"田中太郎"}
type RequestPostUser struct {
	Name  string `json:"user_name"`
	Email string `json:"email"`
}

// ユーザ新規作成
func PostUsers(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var req RequestPostUser
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}
	// ここから
	//creator := registry.GetFactory().BuildCreateUser()
	//res, err := creator.Execute(&usecase.CreateUserRequest {
	//	Name:  req.Name,
	//	Email: req.Email,
	//})
	return Response500(nil)
}
