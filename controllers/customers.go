package controllers

import (
	"net/http"
	"project-alta-store/lib/database"
	"project-alta-store/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterCustomersController(c echo.Context) error {
	var customerModel models.Customers_register

	if e := c.Bind(&customerModel); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error registration customers",
			"status":  e.Error(),
		})
	}

	// cek validation struct
	if e := models.Validate.Struct(customerModel); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error registration customers",
			"status":  e.Error(),
		})
	}

	// Generate "hash" to store from user password
	// hash, _ := HashPassword(customerModel.Password)

	var customer models.Customers
	customer.Username = customerModel.Username
	customer.Email = customerModel.Email
	customer.Password = customerModel.Password

	_, err := database.InsertCustomers(customer)

	// cek before insert to database
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Error registration customers",
			"status":  err.Error(),
		})
	}

	//create cart untuk customer 
	newCustomer,_ := database.GetCustomersByName(customer.Username)
	newCart := models.Carts{
		Customers_id: newCustomer.ID,
		Status: "empty",
	}
	_,e := database.InsertCarts(newCart)

	if e!=nil{
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Error registration user",
			"status":  e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success register customers",
		"status":  "success",
	})
}

func LoginCustomersController(c echo.Context) error {
	var customerLogin models.Customers
	c.Bind(&customerLogin)
	customer, _ := database.LoginCustomers(customerLogin)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login customer",
		"status":  "success",
		"data":    customer,
	})
}
