package entity

// User User
type User struct {
	ID       int64
	Email    string
	Name     string
	Password string
}

// IsEmailEmpty メールアドレスが空か確認する
func (u *User) IsEmailEmpty() bool {
	return len(u.Email) == 0
}
