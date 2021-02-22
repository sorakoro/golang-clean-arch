// +build wireinject

package main

import (
	"time"

	"github.com/google/wire"
	"github.com/sorakoro/golang-clean-arch/adapter/gateway"
	"github.com/sorakoro/golang-clean-arch/domain/usecase"
	"github.com/sorakoro/golang-clean-arch/driver"
)

func InitUserUseCase(timeout time.Duration) (usecase.UserUseCase, func(), error) {
	wire.Build(
		usecase.NewUserUseCase,
		gateway.NewUserRepository,
		driver.NewDBConn,
	)
	return &usecase.UserUseCaseImpl{}, nil, nil
}

func InitArticleUseCase(timeout time.Duration) (usecase.ArticleUseCase, func(), error) {
	wire.Build(
		usecase.NewArticleUseCase,
		gateway.NewArticleRepository,
		driver.NewDBConn,
	)
	return &usecase.ArticleUseCaseImpl{}, nil, nil
}
