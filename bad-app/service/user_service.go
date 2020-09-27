package service

import (
	"github.com/pkg/errors"
	"bad-app/repository"
)

type UserService struct {
	repo repository.UserRepository
}

// リポジトリ作成
// 作成したリポジトリを直接serviceに
func NewUserService() (*UserService, error) {
	repo, err := repository.NewUserRepository()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &UserService{*repo}, nil
}

