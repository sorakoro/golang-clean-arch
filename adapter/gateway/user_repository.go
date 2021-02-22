package gateway

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sorakoro/golang-clean-arch/domain/entity"
	"github.com/sorakoro/golang-clean-arch/domain/usecase"
)

// UserRepositoryImpl UserRepositoryImpl
type UserRepositoryImpl struct {
	db *sqlx.DB
}

// NewUserRepository UserRepositoryを作成する
func NewUserRepository(db *sqlx.DB) usecase.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Store ユーザーを作成する
func (r *UserRepositoryImpl) Store(ctx context.Context, request *usecase.AddUserRequest) (*entity.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES (:name, :email, :password)"
	res, err := r.db.NamedExecContext(ctx, query, &request)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	lastInsertID, _ := res.LastInsertId()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	return &entity.User{ID: lastInsertID, Name: request.Name, Email: request.Email, Password: request.Password}, nil
}

// Fetch ユーザーを取得する
func (r *UserRepositoryImpl) Fetch(ctx context.Context, email string) (*int64, error) {
	user := entity.User{}
	query := "SELECT id from users WHERE email = ?"
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}
	var userID *int64
	userID = &user.ID
	return userID, nil
}
