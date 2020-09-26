package adapter

import (
	"fmt"
	"github.com/guregu/dynamo"
	"github.com/memememomo/nomof"
	"github.com/pkg/errors"
	"go-app/domain"
)

// UserOperator ユーザーを操作する構造体
type UserOperator struct {
	Client *ResourceTableOperator
	Mapper *DynamoModelMapper
	UserEmailUniqGenerator *UserEmailUniqGenerator
}

func (u *UserEmailUniqGenerator) NewUserEmailUniqByUser(user *UserResource) *UserEmailUniq {
	return &UserEmailUniq{
		Email:      user.Email,
		EntityName: u.Mapper.GetEntityNameFromStruct(*user),
		Exists:     true,
		UserID:     user.ID(),
	}
}

func (u *UserOperator) getUserResourceByID(id uint64) (*UserResource, error) {
	var user UserResource
	_, err := u.Mapper.GetEntityByID(id, &UserResource{}, &user)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

// GetUserByEmail メールアドレスからユーザー情報を取得する
func (u *UserOperator) GetUserByEmail(email string) (*domain.UserModel, error) {
	table, err := u.Client.ConnectTable()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fb := nomof.NewBuilder()
	fb.Equal("Email", email)
	fb.BeginsWith("PK", u.Mapper.GetEntityNameFromStruct(UserResource{}))

	var usersDynamo []UserResource
	err = table.Scan().Filter(fb.JoinAnd(), fb.Arg...).All(&usersDynamo)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(usersDynamo) == 0 {
		return nil, errors.WithStack(domain.ErrNotFound)
	}

	return &usersDynamo[0].UserModel, nil
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
	fmt.Println("======================================Error")
	fmt.Println(uniq)
	fmt.Println("======================================")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &userResource.UserModel, nil
}

func (u *UserEmailUniqGenerator) BuildQueryCreateByUser(user *UserResource) (*dynamo.Put, error) {
	table, err := u.Client.ConnectTable()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	uniq := u.NewUserEmailUniqByUser(user)

	fb := nomof.NewBuilder()
	fb.AttributeNotExists("Exists")
	fb.Equal("UserID", user.ID())

	query := table.
		Put(uniq).
		If(fb.JoinOr(), fb.Arg...)

	return query, nil
}
