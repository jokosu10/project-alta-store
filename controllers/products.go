package controllers

import(
	"project-alta-store/lib/database"
	"net/http"
	"project-alta-store/models"
	"github.com/labstack/echo"
)


func GetProductsController(c echo.Context) error{
	product, err := database.GetProducts()

	if err!=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all products",
		"data" : product,
	})
}

func CreateProductsController(c echo.Context) error{
	product := models.Products{
		Categories_id: 1,
		Name: "Yonex b55" ,
		Description: "Fast Racket",
		Quantity: 30,
		Price: 45000,
		Unit: "pcs",
	}
	res,err := database.InsertProducts(product)
	if err!=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"message": "success Create products " + product.Name,
		"data" : res,
	})
}