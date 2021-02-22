package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sorakoro/golang-clean-arch/domain/entity"
	"gopkg.in/go-playground/validator.v9"
)

// UserRepository UserRepository
type UserRepository interface {
	Store(ctx context.Context, request *AddUserRequest) (*entity.User, error)
	Fetch(ctx context.Context, email string) (*int64, error)
}

// UserUseCase UserUseCase
type UserUseCase interface {
	Store(ctx context.Context, request *AddUserRequest) (*AddUserResponse, error)
	Fetch(ctx context.Context, email string) (*FetchUserResponse, error)
}

// UserUseCaseImpl UserUseCaseImpl
type UserUseCaseImpl struct {
	repository  UserRepository
	contextTime time.Duration
}

// AddUserRequest AddUserRequest
type AddUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// AddUserResponse AddUserResponse
type AddUserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchUserResponse FetchUserResponse
type FetchUserResponse struct {
	ID *int64 `json:"id"`
}

// NewUserUseCase UserUseCaseを作成する
func NewUserUseCase(repository UserRepository, timeout time.Duration) UserUseCase {
	return &UserUseCaseImpl{repository: repository, contextTime: timeout}
}

// Store ユーザーを作成する
func (u *UserUseCaseImpl) Store(ctx context.Context, request *AddUserRequest) (*AddUserResponse, error) {
	if ok, err := isRequestValid(request); !ok {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, u.contextTime)
	defer cancel()
	user, err := u.repository.Store(ctx, request)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	return &AddUserResponse{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}

// Fetch ユーザーを取得する
func (u *UserUseCaseImpl) Fetch(ctx context.Context, email string) (*FetchUserResponse, error) {
	user := entity.User{Email: email}
	if user.IsEmailEmpty() {
		fmt.Fprintln(os.Stderr, "email is empty")
		return nil, nil
	}
	userID, err := u.repository.Fetch(ctx, email)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	return &FetchUserResponse{ID: userID}, nil
}

// isRequestValid バリデーション
func isRequestValid(u *AddUserRequest) (bool, error) {
	validator := validator.New()
	err := validator.Struct(u)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return false, err
	}
	return true, nil
}
