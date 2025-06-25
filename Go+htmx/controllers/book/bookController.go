package bookController

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	bookDto "pedro21ribeiro.com/dtos/book"
	"pedro21ribeiro.com/middleware"
	bookService "pedro21ribeiro.com/services/book"
)

func SetUpControllers(group *echo.Group) {
	group.Use(echojwt.WithConfig(middleware.GetConfig()))	
	group.RouteNotFound("/*", func (c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/book")
		return c.Redirect(303, "/book")
	})
	group.GET("", Page)
	group.POST("", Create)
}

func Page(c echo.Context) error {
	formData := bookDto.NewFormData()

	return c.Render(http.StatusOK, "book_landing", formData)
}

func Create(c echo.Context) error {
	name := c.FormValue("name")
	isbn := c.FormValue("isbn")
	date := c.FormValue("date")
	author := c.FormValue("author")
	publisher := c.FormValue("publisher")
	imgUrl := c.FormValue("img_url")


	formData := bookService.CreateBook(name, isbn, date, author, publisher, imgUrl)
	return c.Render(http.StatusCreated, "book_form", formData)
}