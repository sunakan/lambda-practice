package controller

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/glog"
)

// commonHeaders 各レスポンスに共通で含むヘッダー
func commonHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
}

// Response201Body IDを含めた201レスポンス
type Response201Body struct {
	Message string `json:"message"`
	ID      uint64 `json:"id"`
}

// Response201 IDを含めた201レスポンス
func Response201(id uint64) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    commonHeaders(),
		Body:       fmt.Sprintf(`{"message":"OK","id":%d}`, id),
	}
}

// Response500 500レスポンス
func Response500(err error) events.APIGatewayProxyResponse {
	glog.Errorf("%+v\n", err)
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    commonHeaders(),
		Body:       `{"message":"サーバエラーが発生しました。"}`,
	}
}
