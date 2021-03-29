package routes

import(
	"project-alta-store/controllers"
	"github.com/labstack/echo"
)

func Start() *echo.Echo{
	e := echo.New()
	e.GET("/products/:id",controllers.GetProductsByCategoryIdController)
	e.GET("/products",controllers.GetProductsController)
	e.POST("/products",controllers.CreateProductsController)
	e.PUT("/products/:id",controllers.UpdateProductsController)
	e.DELETE("/products/:id",controllers.DeleteProductController)
	return e
}