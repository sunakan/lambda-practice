package domain

// Entity
// User Model
// 名前とメールアドレスを持つというビジネスロジックを構造体で表現
type UserModel struct {
	ID uint64
	Name string
	Email string
}

func NewUserModel(name, email string) *UserModel {
	return &UserModel{Name: name, Email: email}
}
