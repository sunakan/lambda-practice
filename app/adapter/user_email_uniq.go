package adapter

// UserEmailUniq メールアドレス重複チェック用のレコードを表した構造体
type UserEmailUniq struct {
	Email      string `dynamo:"PK"`
	EntityName string `dynamo:"SK"`
	Exists     bool   `dynamo:"Exists"`
	UserID     uint64 `dynamo:"UserID"`
}

type UserEmailUniqGenerator struct {
	Mapper *DynamoModelMapper
	Client *ResourceTableOperator
	PKName string
	SKName string
}

func NewUserEmailUniqGenerator(mapper *DynamoModelMapper, client *ResourceTableOperator, pkName, skName string) *UserEmailUniqGenerator {
	return &UserEmailUniqGenerator{
		Mapper: mapper,
		Client: client,
		PKName: pkName,
		SKName: skName,
	}
}
