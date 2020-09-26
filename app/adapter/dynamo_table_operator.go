package adapter

import (
	"github.com/guregu/dynamo"
)

type TableOperator struct {
	Client    *DynamoClient
	TableName string
}

func NewTableOperator(client *DynamoClient, tableName string) *TableOperator {
	return &TableOperator{
		Client:    client,
		TableName: tableName,
	}
}

func (a *TableOperator) CreateTableForTest(schema interface{}) error {
	return a.Client.CreateTableForTest(a.TableName, schema)
}

func (a *TableOperator) ConnectDB() (*dynamo.DB, error) {
	return a.Client.Connect()
}

func (a *TableOperator) DropTable() error {
	return a.Client.DropTable(a.TableName)
}

func (a *TableOperator) ConnectTable() (*dynamo.Table, error) {
	return a.Client.ConnectTable(a.TableName)
}
