package routes

import (
	"project-alta-store/constants"
	"project-alta-store/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() *echo.Echo {
	e := echo.New()

	//Route Products
	e.GET("/products", controllers.GetProductsController)
	e.POST("/products", controllers.CreateProductsController)
	e.PUT("/products/:id", controllers.UpdateProductsController)
	e.DELETE("/products/:id", controllers.DeleteProductController)

	//Route Categories
	e.GET("/categories", controllers.GetCategoriesController)
	// e.POST("/categories", controllers.CreateCategoriesController)
	e.PUT("/categories/:id", controllers.UpdateCategoriesController)
	e.DELETE("/categories/:id", controllers.DeleteCategoriesController)

	//Route  CartItems
	e.GET("/cartitems/:id", controllers.GetCartitemsByCartIdController)
	e.POST("/cartitems", controllers.CreateCartitemsController)
	e.PUT("/cartitems/:id", controllers.UpdateCartitemsController)
	e.DELETE("/cartitems/:id", controllers.DeleteCartitemsController)

	//Route Carts
	e.PUT("/carts/:id", controllers.UpdateCartsController)

	// route auth
	e.POST("/register", controllers.RegisterCustomersController)
	e.POST("/login", controllers.LoginCustomersController)

	// route jwt
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.POST("/categories", controllers.CreateCategoriesController)

	//Order Auth
	jwtAuth.POST("/orders", controllers.CreateOrdersController)
	jwtAuth.GET("/orders", controllers.GetOrderController)

	//Payment Auth
	jwtAuth.GET("/payments", controllers.GetPaymentsController)
	jwtAuth.PUT("/payments/:id", controllers.UpdatePaymentsController)

	return e
}
