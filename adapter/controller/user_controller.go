package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sorakoro/golang-clean-arch/domain"
	"github.com/sorakoro/golang-clean-arch/domain/usecase"
)

// UserController UserController
type UserController struct {
	usecase usecase.UserUseCase
}

// NewUserController UserControllerを作成する
func NewUserController(e *echo.Echo, usecase usecase.UserUseCase) {
	controller := &UserController{usecase: usecase}
	e.POST("/user", controller.Store)
	e.GET("/user", controller.Fetch)
}

// Store ユーザーを作成する
func (uc *UserController) Store(c echo.Context) error {
	var req usecase.AddUserRequest
	err := c.Bind(&req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	ctx := c.Request().Context()
	res, err := uc.usecase.Store(ctx, &req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		err, ok := err.(*domain.AppError)
		if !ok || err.ErrType == domain.ErrDatabase {
			return c.NoContent(http.StatusInternalServerError)
		}
		if err.ErrType == domain.ErrValidation {
			return c.NoContent(http.StatusBadRequest)
		}
	}
	return c.JSON(http.StatusCreated, res)
}

// Fetch ユーザーを取得する
func (uc *UserController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	email := c.QueryParams().Get("email")
	res, err := uc.usecase.Fetch(ctx, email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}
