package interactor

import (
	"github.com/pkg/errors"
	"go-app/domain"
	"go-app/usecase"
)

var (
	ErrUniqEmail = errors.New("unique email error")
)

// UserCretor ユーザ新規作成
type UserCreator struct {
	UserRepository domain.UserRepository
	UniqChecker    *domain.UserEmailUniqChecker
}

// UserCreator
// UserCreatorのコンストラクタ的な
func NewCreateUser(repos domain.UserRepository, checker *domain.UserEmailUniqChecker) *UserCreator {
		return &UserCreator {
		UserRepository: repos,
		UniqChecker:    checker,
	}
}

// Execute ユーザーを新規作成
func (u *UserCreator) Execute(req *usecase.CreateUserRequest) (*usecase.CreateUserResponse, error) {
	isUniq, err := u.UniqChecker.IsUniqueEmail(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !isUniq {
		return nil, errors.WithStack(ErrUniqEmail)
	}
	user, err := u.UserRepository.CreateUser(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &usecase.CreateUserResponse{User: user}, nil
}
