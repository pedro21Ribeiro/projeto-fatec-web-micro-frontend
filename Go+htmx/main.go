package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pedro21ribeiro.com/controllers"
	"pedro21ribeiro.com/dbConector"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {

	if dbConector.InstanciateDB() != nil {
		fmt.Println("Error connecting to DB, aborting now")
		return
	}
	fmt.Println("Database connected")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	controllers.SetUpControllers(e)
	e.Static("/css", "static/css")
	e.Static("/js", "static/js")

	e.Logger.Fatal(e.Start(":42069"))
}
