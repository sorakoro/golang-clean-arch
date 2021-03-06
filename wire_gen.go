// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/sorakoro/golang-clean-arch/adapter/gateway"
	"github.com/sorakoro/golang-clean-arch/domain/usecase"
	"github.com/sorakoro/golang-clean-arch/driver"
	"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitUserUseCase(timeout time.Duration) (usecase.UserUseCase, func(), error) {
	db, cleanup, err := driver.NewDBConn()
	if err != nil {
		return nil, nil, err
	}
	userRepository := gateway.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, timeout)
	return userUseCase, func() {
		cleanup()
	}, nil
}

func InitArticleUseCase(timeout time.Duration) (usecase.ArticleUseCase, func(), error) {
	db, cleanup, err := driver.NewDBConn()
	if err != nil {
		return nil, nil, err
	}
	articleRepository := gateway.NewArticleRepository(db)
	articleUseCase := usecase.NewArticleUseCase(articleRepository, timeout)
	return articleUseCase, func() {
		cleanup()
	}, nil
}
