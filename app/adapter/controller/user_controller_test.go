package controller

import (
	"go-app/mocks"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPostUsers_201 新規作成 成功時
func TestPostUsers_201(t *testing.T) {
	// テスト用DynamoDBを設定
	tables := mocks.SetupDB(t)
	defer tables.Cleanup()

	// リクエストパラメータ設定
	body := map[string]interface{}{
		"user_name": "テスト名前",
		"email":     "test@example.com",
	}
	bodyStr, err := json.Marshal(body)
	assert.NoError(t, err)

	// 新規作成処理
	res := PostUsers(events.APIGatewayProxyRequest{
		Body: string(bodyStr),
	})

	// レスポンスコードをチェック
	assert.Equal(t, 201, res.StatusCode)

	// JSONからmap型に変換
	var resBody map[string]interface{}
	err = json.Unmarshal([]byte(res.Body), &resBody)
	assert.NoError(t, err)

	// IDをチェック
	id := uint64(resBody["id"].(float64))
	assert.Equal(t, uint64(1), id)

	// DynamoDBに保存されたデータをチェック
	user, err := tables.UserOperator.GetUserByID(id)
	assert.NoError(t, err)
	assert.Equal(t, body["user_name"].(string), user.Name)
	assert.Equal(t, body["email"].(string), user.Email)
}
