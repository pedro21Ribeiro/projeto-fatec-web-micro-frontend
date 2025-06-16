package controllers

import (
	"github.com/labstack/echo/v4"
	testController "pedro21ribeiro.com/controllers/test"
	userController "pedro21ribeiro.com/controllers/user"
)

func SetUpControllers(e *echo.Echo){
	users := e.Group("/user")
	tests := e.Group("/test")

	userController.SetUpControllers(users)
	testController.SetUpControllers(tests)
}