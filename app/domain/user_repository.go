package domain

// User Respository
// ユーザモデルのリポジトリ
// interfaceの定義のみ
// これで特定のDBに依存していないことになる
type UserRepository interface {
	CreateUser(newUser *UserModel) (*UserModel, error)
//	GetUsers() ([]*UserModel, error)
	GetUserByID(id uint64) (*UserModel, error)
	GetUserByEmail(email string) (*UserModel, error)
}
