package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	secretloader "pedro21ribeiro.com/secretLoader"
)

func GetConfig() echojwt.Config{
	secret := secretloader.GetAuthEnv().SECRET
	
	return echojwt.Config{
		SigningKey: []byte(secret),
		ErrorHandler: ErrorHandler,
		TokenLookup: "cookie:auth",
	}
}

func ErrorHandler(c echo.Context, err error) error {
	c.Response().Header().Set("HX-Redirect", "/user")
	
	fmt.Printf("\nRedirecting becaause: %s\n", err.Error())

	return c.Redirect(http.StatusSeeOther, "/user")
}