package routes

import (
	"project-alta-store/controllers"

	"github.com/labstack/echo"
)

func Start() *echo.Echo {
	e := echo.New()

	// route auth
	e.POST("/register", controllers.RegisterCustomersController)
	e.POST("/login", controllers.LoginCustomersController)

	// route products
	e.GET("/products", controllers.GetProductsController)
	e.POST("/products", controllers.CreateProductsController)

	return e
}
