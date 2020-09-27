package repository

import (
	"github.com/guregu/dynamo"
	"fmt"
)

// IFaceは、わかりやすさのためつける
// ホントはとりたい
type UserRepositoryIFace interface {
	CreateUser()
}

type userRepository struct {
	db dynamo.DB
}

func NewUserRepository(db *dynamo.DB) UserRepositoryIFace {
	return &userRepository{*db}
}

func (ur *userRepository) CreateUser() {
	fmt.Println("ユーザを作成します")
}
