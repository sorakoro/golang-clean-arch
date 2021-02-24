package entity

import (
	"errors"

	"github.com/sorakoro/golang-clean-arch/domain"
)

// User User
type User struct {
	ID       int64
	Email    string
	Name     string
	Password string
}

// CheckEmailEmpty メールアドレスが空か確認する
func (u *User) CheckEmailEmpty() error {
	if len(u.Email) == 0 {
		return &domain.AppError{ErrType: domain.ErrEmailEmpty, Err: errors.New("email is empty")}
	}
	return nil
}
