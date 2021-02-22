package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	userController "github.com/sorakoro/golang-clean-arch/adapter/controller"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userUsecase, cleanup, _ := InitUserUseCase(timeoutContext)
	userController.NewUserController(e, userUsecase)

	defer cleanup()

	servePort := viper.GetString("server.port")
	err := e.Start(servePort)
	if err != nil {
		panic(err)
	}
}
