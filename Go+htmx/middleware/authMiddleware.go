package middleware

import (
	"fmt"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	userDto "pedro21ribeiro.com/dtos/user"
	secretloader "pedro21ribeiro.com/secretLoader"
)

func GetConfig() echojwt.Config{
	secret := secretloader.GetAuthEnv().SECRET
	
	return echojwt.Config{
		SigningKey: []byte(secret),
		ErrorHandler: ErrorHandler,
		TokenLookup: "cookie: auth",
	}
}

func ErrorHandler(c echo.Context, err error) error {
	c.Response().Header().Set("HX-Redirect", "/user")

	cookie, err := c.Cookie("auth"); if err!= nil {
		return c.Render(401, "user_landing", userDto.NewError())
	}

	fmt.Printf("Cookie com token: %s\n", cookie.Value)

	return c.Render(401, "user_landing", userDto.NewError())
}