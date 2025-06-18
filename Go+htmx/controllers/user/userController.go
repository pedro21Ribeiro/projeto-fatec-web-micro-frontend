package userController

import (
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	userDto "pedro21ribeiro.com/dtos/user"
	"pedro21ribeiro.com/services/user"
)

func SetUpControllers(group *echo.Group) {
	group.GET("", landPage)
	group.POST("/login", login)
}

func landPage(c echo.Context) error {
	return c.Render(200, "user_landing", userDto.NewError())
}

func login(c echo.Context) error{

	//Getting the JSON data
	email := c.FormValue("email")
	password := c.FormValue("password")

	//generating the token
	token, err := userService.Login(email, password); if err != nil {
		formError := userDto.NewError()
		formError.Error = err.Error()
		return c.Render(422, "form_error", formError)
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 8)
	cookie.Path = "/"
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Redirect", "/user")

	return c.String(200, "Sucess")
}