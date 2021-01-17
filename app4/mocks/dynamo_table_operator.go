package mocks

import (
	"testing"
	"fmt"
	"os"
	"app4/registry"
)

// Operatorの意味：操作するもの
type DynamoTableOperator struct {
//	Operator *adapter.ResourceTableOperator
//	UserOperator domain.UserRepository
}

// DynamoDBのTableのMockを作成
func SetupDB(t *testing.T) *DynamoTableOperator {
	t.Helper()
	operator := &DynamoTableOperator{}
	registry.ClearFactory()
	f := registry.GetFactory()
	fmt.Println("=============[ Factory ]")
	fmt.Println(f)
	fmt.Println("=============")
	os.Setenv("DYNAMO_TABLE_NAME", "hoge-table")
	return operator
}

// テーブルを削除
func (d *DynamoTableOperator) Cleanup() {
	fmt.Println("Mock Dynamo Table Cleanupするぜー")
	//d.Operator.DropTable()
}
