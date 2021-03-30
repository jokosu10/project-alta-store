package controllers

import (
	"net/http"
	"project-alta-store/lib/database"
	"project-alta-store/lib/utils"
	"project-alta-store/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetPaymentsController(c echo.Context) error {
	orderid := c.QueryParam("order")
	if utils.StringIsNotNumber(orderid) {
		id, _ := strconv.Atoi(orderid)
        payment,err := database.GetPaymentByOrderId(id)
		if err!=nil{
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "Fail",
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK,map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "succes get payment",
			"data" : payment,
		})
	} else if len(c.QueryParam("order")) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	} else{
		payment := database.GetAllPayments()
		return c.JSON(http.StatusOK,map[string]interface{}{
			"code":    200,
			"status":  "success",
			"message": "success get all payment",
			"data" : payment,
		})
	}
}


func UpdatePaymentsController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	
	if !utils.StringIsNotNumber(c.Param("id")){
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "Fail",
			"message":  "invalid id supplied",
		})
	}
	var post_body models.Payments_update

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status": "fail",
			"message":  e.Error(),
		})
	}

	if e:= models.Validate.Struct(post_body);e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status" : "fail",
			"message":  e.Error(),
		})
	}
	

	payment,e := database.GetPaymentById(id)
	if  e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    404,
			"status": "fail",
			"message":  e.Error(),
		})
	}

	payment.Payment_status = post_body.Payment_status
	
    err := database.InsertPayments(payment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status": "fail",
			"message":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code" : 200,
		"status" : "success",
		"message": "success update payments ",
	})
}