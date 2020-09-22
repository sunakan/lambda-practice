package domain

// User Respository
// ユーザモデルのリポジトリ
// interfaceの定義のみ
// これで特定のDBに依存していないことになる
type UserRepository interface {
	GetUsers() ([]*UserModel, error)
	CreateUser(newUser *UserModel) (*UserModel, error)
	GetUserByID(id uint64) (*UserModel, error)
}
