package adapter

type ResourceTableOperator struct {
	TableOperator
}

type ResourceSchema struct {
	PK string `dynamo:"PK,hash"`
	SK string `dynamo:"SK,range"`
}

func NewResourceTableOperator(client *DynamoClient, tableName string) *ResourceTableOperator {
	return &ResourceTableOperator{
		TableOperator: *NewTableOperator(client, tableName),
	}
}

func (r *ResourceTableOperator) CreateTableForTest() error {
	return r.TableOperator.CreateTableForTest(&ResourceSchema{})
}
