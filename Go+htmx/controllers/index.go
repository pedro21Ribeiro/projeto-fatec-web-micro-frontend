package controllers

import (
	"github.com/labstack/echo/v4"
	bookController "pedro21ribeiro.com/controllers/book"
	testController "pedro21ribeiro.com/controllers/test"
	userController "pedro21ribeiro.com/controllers/user"
)

func SetUpControllers(e *echo.Echo){
	users := e.Group("/user")
	tests := e.Group("/test")
	books := e.Group("/book")

	userController.SetUpControllers(users)
	testController.SetUpControllers(tests)
	bookController.SetUpControllers(books)
}