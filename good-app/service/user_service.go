package service

import (
	"good-app/repository"
)

// IFaceは、わかりやすさのためつける
// FromServiceも
// ホントは取る
type UserServiceIFace interface {
	CreateUserFromService()
}

type userService struct {
	repo repository.UserRepositoryIFace
}

// この時点ではDynamoDBかはわからない
func NewUserService(repo repository.UserRepositoryIFace) UserServiceIFace {
    return &userService{repo}
}

func (us *userService) CreateUserFromService() {
	us.repo.CreateUser()
}
