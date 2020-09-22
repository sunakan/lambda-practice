package interactor

import "go-app/domain"
import "go-app/usecase"
import "github.com/pkg/errors"

// UserCretor ユーザ新規作成
type UserCreator struct {
	UserRepository domain.UserRepository
}

// UserCreator
// UserCreatorのコンストラクタ的な
func NewUserCreator(repos domain.UserRepository) *UserCreator {
	return &UserCreator {
		UserRepository: repos,
	}
}

func NewCreateUser(repos domain.UserRepository) *UserCreator {
	return &UserCreator{
		UserRepository: repos,
	}
}

// Execute
func (u *UserCreator) Execute(req *usecase.CreateUserRequest) (*usecase.CreateUserResponse, error) {
	user, err := u.UserRepository.CreateUser(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &usecase.CreateUserResponse{User: user}, nil
}
