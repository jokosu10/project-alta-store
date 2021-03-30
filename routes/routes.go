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


	//Route  CartItems
	e.GET("/cartitems/:id",controllers.GetCartitemsByCartIdController)
	e.POST("/cartitems",controllers.CreateCartitemsController)
	e.PUT("/cartitems/:id",controllers.UpdateCartitemsController)
	e.DELETE("/cartitems/:id",controllers.DeleteCartitemsController)

	//Route Carts
	e.PUT("/carts/:id",controllers.UpdateCartsController)

	// route auth
	e.POST("/register", controllers.RegisterCustomersController)
	e.POST("/login", controllers.LoginCustomersController)

	//Order Auth
	e.POST("/orders",controllers.CreateOrdersController)
	e.GET("/orders",controllers.GetOrderController)
	return e
}
