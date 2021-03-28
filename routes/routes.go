package routes

import(
	"project-alta-store/controllers"
	"github.com/labstack/echo"
)

func Start() *echo.Echo{
	e := echo.New()
	e.GET("/products",controllers.GetProductsController)
	e.POST("/products",controllers.CreateProductsController)

	return e
}