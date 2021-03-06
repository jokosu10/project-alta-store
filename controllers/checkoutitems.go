package controllers

import (
	"net/http"
	"project-alta-store/lib/database"
	"project-alta-store/lib/utils"
	"project-alta-store/models"
	"strconv"
	"github.com/labstack/echo"
)

func GetCheckoutItemsController(c echo.Context) error {

	if !utils.StringIsNotNumber(c.QueryParam("order")) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Fail",
			"message": "invalid id supplied",
		})
	} 

	id, _ := strconv.Atoi(c.QueryParam("order"))
		items, err := database.GetCheckOutItemByOrderId(id)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"status":  "fail",
				"message": err.Error(),
			})
		}

		res := models.Checkoutitem_response{
			Code:    200,
			Status:  "Success",
			Message: "Success",
			Data:    items,
		}
		return c.JSON(http.StatusOK, res)
}