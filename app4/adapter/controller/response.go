package controller

import (
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

// Response500 500レスポンス
func Response500(err error) events.APIGatewayProxyResponse {
	glog.Errorf("%+v\n", err)
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    commonHeaders(),
		Body:       `{"message":"サーバエラーが発生しました。"}`,
	}
}
