package controllers

import (
	"fmt"
	"net/http"
	"project-alta-store/lib/database"
	"project-alta-store/lib/utils"
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

func GetProductsByCategoryIdController(c echo.Context) error {

	id,_ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")){
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "Fail",
			"message":  "invalid id supplied",
		})
	}

	product, err := database.GetProductsByCategoryId(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "fail",
			"message":  err.Error(),
		})
	}

	

	if len(product)==0{
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    404,
			"status": "fail",
			"message":  "product not found",
		})
	}

	res := models.Products_response{
		Code:    "200",
		Status:  "Success",
		Message: "Success",
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

func UpdateProductsController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")){
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "Fail",
			"message":  "invalid id supplied",
		})
	}
	var post_body models.Products_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    500,
			"status": "Fail",
			"message":  e.Error(),
		})
	}
	
	if !utils.IsInt(id){
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "Fail",
			"message":  "invalid id supplied",
		})
	}

	product,e := database.GetProductsById(id)
	if  e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status": "Failed",
			"message":  e.Error(),
		})
	}
	
	product.Name = utils.CompareStrings(product.Name,post_body.Name)
	product.Description = utils.CompareStrings(product.Description,post_body.Description)
	product.Unit =utils.CompareStrings(product.Unit,post_body.Unit)
	product.Quantity = post_body.Quantity
	product.Categories_id = post_body.Categories_id
	product.Price = post_body.Price
	
	
	fmt.Println(product)
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

func DeleteProductController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))

	if !utils.StringIsNotNumber(c.Param("id")){
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "Fail",
			"message":  "invalid id supplied",
		})
	}

	err := database.DeleteProductsById(id)
	if err!=nil{
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status": "Fail",
			"message":  err.Error(),
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"code":    200,
		"status": "success",
		"message":  "product succesfully deleted",
	})
}