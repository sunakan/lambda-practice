package adapter

import (
	"go-app/domain"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

// UserOperator ユーザーを操作する構造体
type UserOperator struct {
	Client *ResourceTableOperator
	Mapper *DynamoModelMapper
}

func (u *UserOperator) getUserResourceByID(id uint64) (*UserResource, error) {
	var user UserResource
	_, err := u.Mapper.GetEntityByID(id, &UserResource{}, &user)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

// Execute IDからユーザー情報を取得する
func (u *UserOperator) GetUserByID(id uint64) (*domain.UserModel, error) {
	userResource, err := u.getUserResourceByID(id)
	if err != nil {
		if err.Error() == dynamo.ErrNotFound.Error() {
			return nil, errors.WithStack(domain.ErrNotFound)
		}
		return nil, errors.WithStack(err)
	}
	return &userResource.UserModel, nil
}

// CreateUser ユーザーを新規作成する
func (u *UserOperator) CreateUser(userModel *domain.UserModel) (*domain.UserModel, error) {
	conn, err := u.Client.ConnectDB()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userResource := NewUserResource(userModel, u.Mapper)

	tx := conn.WriteTx()

	r, err := u.Mapper.BuildQueryCreate(userResource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	uniq, err := u.UserEmailUniqGenerator.BuildQueryCreateByUser(userResource)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = tx.Put(r).Put(uniq).Run()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &userResource.UserModel, nil
}
