package testController

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"pedro21ribeiro.com/auth/jwt"
)

type Response struct {
	Values string
}

func NewResponse() Response{
	return Response{}
}

type Page struct {
	Response Response
}

func NewPage() Page{
	return Page{
		Response: NewResponse(),
	}
}

func SetUpControllers(group *echo.Group){
	group.GET("", Test)
	group.GET("/page", Index)
}

func Test(c echo.Context) error{
	response := NewResponse()

	cookie, err := c.Cookie("auth"); if err != nil {
		c.Response().Header().Add("HX-Redirect", "/test/page")
		return c.HTML(303,"Sem cookie de auth")
	}

	_ , err = jwt.GetTokenData(cookie.Value); if err != nil {
		c.Response().Header().Add("HX-Redirect", "/test/page")
		return c.HTML(303,"jwt expirado/invalido")
	}


	// auth := headers.Get("AUTH"); if (auth != "123456"){
	// 	fmt.Println("NÃO LOGADO")

	// 	c.Response().Header().Add("HX-Redirect", "/test/page")

	// 	return c.HTML(303,"")
	// }

	fmt.Println("Funcionou")
	response.Values = "funcionou"
	return c.Render(200, "response", response)
}

func Index(c echo.Context) error{
	page := NewPage()
	return c.Render(200, "index", page)
} 