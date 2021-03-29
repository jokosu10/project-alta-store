package routes

import (
	"project-alta-store/controllers"

	"github.com/labstack/echo"
)

func Start() *echo.Echo {
	e := echo.New()

	//Route Products
	e.GET("/products/:id",controllers.GetProductsByCategoryIdController)
	e.GET("/products",controllers.GetProductsController)
	e.POST("/products",controllers.CreateProductsController)
	e.PUT("/products/:id",controllers.UpdateProductsController)
	e.DELETE("/products/:id",controllers.DeleteProductController)

	//Route Categories
	e.GET("/categories",controllers.GetCategoriesController)
	e.POST("/categories",controllers.CreateCategoriesController)
	e.PUT("/categories/:id",controllers.UpdateCategoriesController)
	e.DELETE("/categories/:id",controllers.DeleteCategoriesController)



	// route auth
	e.POST("/register", controllers.RegisterUserController)

	return e
}
