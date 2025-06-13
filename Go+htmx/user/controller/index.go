package controller

import (
	"github.com/labstack/echo/v4"
	"pedro21ribeiro.com/user/service"
)

func SetUpControllers(e *echo.Echo) {
	group := e.Group("/user")

	group.GET("", getAllUsers)
}

func getAllUsers(c echo.Context) error{
	users, err, status := service.GetAllUsers(); if err != nil {
		return c.JSON(status, err.Error())
	}

	return c.JSON(status, users)
}