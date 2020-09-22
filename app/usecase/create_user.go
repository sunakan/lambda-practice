package usecase

// usecase/create_user.go
// ユーザ新規作成UseCase
// ここでは具体的な処理を実装しない
// Interactorで実装をする

import "go-app/domain"

// ICreateUser
// 新規作成リクエストを受け取り、新規作成レスポンスを返す
type ICreateUser interface {
	Execute(req *CreateUserRequest) (*CreateUserResponse, error)
}

// 新規作成リクエスト
type CreateUserRequest struct {
	Name string
	Email string
}

// 新規作成レスポンス
type CreateUserResponse struct {
	User *domain.UserModel
}

func (u *CreateUserResponse) GetUserID() uint64 {
	return u.User.ID
}

// 新規作成リクエストからユーザモデルを作成して参照を返す
func (u *CreateUserRequest) ToUserModel() *domain.UserModel {
	return domain.NewUserModel(u.Name, u.Email)
}
