package routes

import (
	"project-alta-store/constants"
	"project-alta-store/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	//Route Products
	jwtAuth.GET("/products", controllers.GetProductsController)
	jwtAuth.POST("/products", controllers.CreateProductsController)
	jwtAuth.PUT("/products/:id", controllers.UpdateProductsController)
	jwtAuth.DELETE("/products/:id", controllers.DeleteProductController)

	//Route Categories
	jwtAuth.GET("/categories", controllers.GetCategoriesController)
	jwtAuth.POST("/categories", controllers.CreateCategoriesController)
	jwtAuth.PUT("/categories/:id", controllers.UpdateCategoriesController)
	jwtAuth.DELETE("/categories/:id", controllers.DeleteCategoriesController)

	//Route  CartItems
	jwtAuth.GET("/cartitems/:id", controllers.GetCartitemsByCartIdController)
	jwtAuth.POST("/cartitems", controllers.CreateCartitemsController)
	jwtAuth.PUT("/cartitems/:id", controllers.UpdateCartitemsController)
	jwtAuth.DELETE("/cartitems/:id", controllers.DeleteCartitemsController)

	//Route Carts
	jwtAuth.PUT("/carts/:id", controllers.UpdateCartsController)

	// route auth
	jwtAuth.POST("/register", controllers.RegisterCustomersController)
	jwtAuth.POST("/login", controllers.LoginCustomersController)

	//checkout items
	jwtAuth.GET("/checkoutitems", controllers.GetCheckoutItemsController)

	//Order Auth
	jwtAuth.POST("/orders", controllers.CreateOrdersController)
	jwtAuth.GET("/orders", controllers.GetOrderController)

	//Payment Auth
	jwtAuth.GET("/payments", controllers.GetPaymentsController)
	jwtAuth.PUT("/payments/:id", controllers.UpdatePaymentsController)

	// update profile
	jwtAuth.PUT("/customers/:id", controllers.UpdateProfileCustomersController)

	return e
}
