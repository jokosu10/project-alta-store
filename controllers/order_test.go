package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-alta-store/lib/database"
	"project-alta-store/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func insertDBTestOrder() {
	database.InsertCategories(models.Categories{Name: "Sabun", Description: "BersihBadan"})
	database.InsertCategories(models.Categories{Name: "Olahraga", Description: "BersihBadan"})
	database.InsertProducts(models.Products{Name: "Yonex23", Categories_id: 0, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Bola Kasti", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Shell", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertCustomers(models.Customers{Username: "Aditya", Email: "as.com", Password: "123", Address: "a"})
	database.InsertCustomers(models.Customers{Username: "rein", Email: "ass.com", Password: "123", Address: "a"})
	database.InsertCustomers(models.Customers{Username: "Liven", Email: "sdass.com", Password: "123", Address: "a"})
	database.InsertCartitems(models.Cartitems{Carts_id: 0, Products_id: 0, Quantity: 5})
	database.InsertCartitems(models.Cartitems{Carts_id: 0, Products_id: 1, Quantity: 5})
	database.InsertCartitems(models.Cartitems{Carts_id: 1, Products_id: 1, Quantity: 4})
	database.InsertCartitems(models.Cartitems{Carts_id: 2, Products_id: 1, Quantity: 3})
	database.InsertCourier(models.Couriers{Name: "tiki"})
	database.InsertCourier(models.Couriers{Name: "JNE"})
}

func TestCreateOrderController(t *testing.T) {
	var testCases = []struct {
		code    int
		status  string
		json    string
		message string
		expectedLen int
	}{
		{
			200,
			"success",
			"{\r\n        \"customers_id\":1,\r\n        \"couriers_id\":1,\r\n        \"payment_method\":\" BCA\",\r\n        \"payment_start_date\":\"2021-03-29 10:42:44.710\",\r\n        \"payment_end_date\" : \"2021-03-30 10:42:44.710\",\r\n        \"payment_status\": \"waiting\",\r\n        \"payment_amount\" : 235000\r\n}\r\n",
			"success insert order",
			1,
		}, {
			200,
			"success",
			"{\r\n        \"customers_id\":2,\r\n        \"couriers_id\":1,\r\n        \"payment_method\":\" BCA\",\r\n        \"payment_start_date\":\"2021-03-29 10:42:44.710\",\r\n        \"payment_end_date\" : \"2021-03-30 10:42:44.710\",\r\n        \"payment_status\": \"waiting\",\r\n        \"payment_amount\" : 235000\r\n}\r\n",
			"success insert order",
			2,
		},
	}
	e := InitEcho()
	insertDBCart()

	for _, testcase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testcase.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/orders")
		if assert.NoError(t, CreateOrdersController(c)) {
			body := rec.Body.String()
			var response models.SuccessResponse
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				fmt.Println("error unmarshall")
			}
			assert.Equal(t, response.Code, testcase.code, "Code not expected")
			assert.Equal(t, response.Status, testcase.status, "Status not expected")
			assert.Equal(t, response.Message, testcase.message, "Message not expectred")
		}
	}
}