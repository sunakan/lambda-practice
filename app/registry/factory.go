package registry

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"go-app/adapter"
	"go-app/domain"
	"go-app/interactor"
	"go-app/usecase"
)

var FactorySingleton *Factory

// Factory 様々なインスタンスを生成する構造体
type Factory struct {
	Envs  *Envs
	cache map[string]interface{}
}

// ClearFactory 使いまわしているインスタンスを削除する
func ClearFactory() {
	FactorySingleton = nil
}

// GetFactory Factoryのインスタンスを取得する
func GetFactory() *Factory {
	if FactorySingleton == nil {
		FactorySingleton = &Factory{
			Envs: NewEnvs(),
		}
	}
	return FactorySingleton
}

// container cacheにインスタンスがある場合はそれを返し、なければ新規作成する
// 返り値は空インターフェース
func (f *Factory) container(key string, builder func() interface{}) interface{} {
	if f.cache == nil {
		f.cache = map[string]interface{}{}
	}
	if f.cache[key] == nil {
		f.cache[key] = builder()
	}
	return f.cache[key]
}

// BuildCreateUser ユーザー作成UseCaseインスタンスを生成
func (f *Factory) BuildCreateUser() usecase.ICreateUser {
	return f.container("CreateUser", func() interface{} {
		return interactor.NewCreateUser(
			f.BuildUserOperator(),
			f.BuildUserEmailUniqChecker())
	}).(usecase.ICreateUser)
}

// BuildUserEmailUniqGenerator ユーザーのメールアドレス重複チェック用のレコード生成機のインスタンスを生成
func (f *Factory) BuildUserEmailUniqGenerator() *adapter.UserEmailUniqGenerator {
	return f.container("UserEmailUniqGenerator", func() interface{} {
		return adapter.NewUserEmailUniqGenerator(
			f.BuildDynamoModelMapper(),
			f.BuildResourceTableOperator(),
			f.Envs.DynamoPKName(),
			f.Envs.DynamoSKName())
	}).(*adapter.UserEmailUniqGenerator)
}

// BuildUserOperator ユーザー情報関連の操作を行うインスタンスを生成
// 実体としてはUserOperatorを返しているが、UserRepositoryとして返している
func (f *Factory) BuildUserOperator() domain.UserRepository {
	return f.container("UserOperator", func() interface{} {
		return &adapter.UserOperator {
			Client:                 f.BuildResourceTableOperator(),
			Mapper:                 f.BuildDynamoModelMapper(),
			UserEmailUniqGenerator: f.BuildUserEmailUniqGenerator(),
		}
	}).(domain.UserRepository)
}

// BuildResourceTableOperator DynamoDBのテーブルに接続するためのインスタンスを生成
func (f *Factory) BuildResourceTableOperator() *adapter.ResourceTableOperator {
	return f.container("ResourceTableOperator", func() interface{} {
		return adapter.NewResourceTableOperator(
			f.BuildDynamoClient(),
			f.Envs.DynamoTableName())
	}).(*adapter.ResourceTableOperator)
}

// BuildDynamoModelMapper ModelからDynamoDBに保存する形式に変換するためのインスタンスを生成
func (f *Factory) BuildDynamoModelMapper() *adapter.DynamoModelMapper {
	return f.container("DynamoModelMapper", func() interface{} {
		return &adapter.DynamoModelMapper{
			Client:    f.BuildResourceTableOperator(),
			TableName: f.Envs.DynamoTableName(),
			PKName:    f.Envs.DynamoPKName(),
			SKName:    f.Envs.DynamoSKName(),
		}
	}).(*adapter.DynamoModelMapper)
}

// BuildUserEmailUniqChecker ユーザーのメールアドレス重複チェックインスタンスを生成
func (f *Factory) BuildUserEmailUniqChecker() *domain.UserEmailUniqChecker {
	return f.container("UserEmailUniqChecker", func() interface{} {
		return domain.NewUserEmailUniqChecker(f.BuildUserOperator())
	}).(*domain.UserEmailUniqChecker)
}

// BuildDynamoClient DynamoDBに接続するためのインスタンスを生成
func (f *Factory) BuildDynamoClient() *adapter.DynamoClient {
	return f.container("DynamoClient", func() interface{} {
		config := &aws.Config{
			Region: aws.String("ap-northeast-1"),
		}

		if f.Envs.DynamoLocalEndpoint() != "" {
			config.Credentials = credentials.NewStaticCredentials("dummy", "dummy", "dummy")
			config.Endpoint = aws.String(f.Envs.DynamoLocalEndpoint())
		}
		return adapter.NewClient(config)
	}).(*adapter.DynamoClient)
}

