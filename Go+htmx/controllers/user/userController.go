package userController

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"pedro21ribeiro.com/dtos/request"
	"pedro21ribeiro.com/services/user"
)

func SetUpControllers(group *echo.Group) {
	group.GET("", getAllUsers)
	group.POST("/login", login)
}

func getAllUsers(c echo.Context) error{
	users, err, status := userService.GetAllUsers(); if err != nil {
		return c.JSON(status, err.Error())
	}

	return c.JSON(status, users)
}

func login(c echo.Context) error{
	data := new(request.LoginDto)

	//Getting the JSON data
	err := c.Bind(data); if err != nil || data.Email == "" || data.Password == "" {
		return c.String(400, "Bad request")
	}

	//generating the token
	token, err := userService.Login(data.Email, data.Password); if err != nil {
		return c.String(400, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 8)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.String(200, token)
}