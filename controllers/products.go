package controllers

import (
	"fmt"
	"net/http"
	"project-alta-store/lib/database"
	"project-alta-store/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetProductsController(c echo.Context) error {
	product, err := database.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := models.Products_response{
		Code:    "200",
		Message: "Success",
		Status:  "Success",
		Data:    product,
	}
	return c.JSON(http.StatusOK, res)
}

func CreateProductsController(c echo.Context) error {

	var post_body models.Products_post
	// c.Bind(&post_body)
	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    500,
			"message": "Fail insert data",
			"status":  e.Error(),
		})
	}
	if e:= models.Validate.Struct(post_body);e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    500,
			"message": "Fail insert data",
			"status":  e.Error(),
		})
	}

	var product models.Products
	product.Categories_id = post_body.Categories_id
	product.Name = post_body.Name
	product.Description = post_body.Description
	product.Quantity = post_body.Quantity
	product.Price = post_body.Price
	product.Unit = post_body.Unit



	_, err := database.InsertProducts(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Fail insert data",
			"status":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create products ",
		"data":    product,
	})
}

// func UpdateProductsController(c echo.Context) error {
// 	id,_ := strconv.Atoi(c.Param("id"))

// 	var post_body models.Products_update

// 	if e := c.Bind(&post_body); e != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
// 			"code":    500,
// 			"message": "Fail Update data",
// 			"status":  e.Error(),
// 		})
// 	}


	
// 	// product:= models.Products{
// 	// 	ID : id,
// 	// 	Categories_id : post_body.Categories_id,
// 	// 	Name : post_body.Name,
// 	// 	Description : post_body.Description,
// 	// 	Quantity : post_body.Quantity,
// 	// 	Price : post_body.Price,
// 	// 	Unit : post_body.Unit,
// 	// }

// }