package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"project-alta-store/lib/database"
	"project-alta-store/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func insertDBCart() {
	database.InsertCategories(models.Categories{Name: "Sabun", Description: "BersihBadan"})
	database.InsertCategories(models.Categories{Name: "Olahraga", Description: "BersihBadan"})
	database.InsertProducts(models.Products{Name: "Yonex23", Categories_id: 0, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Bola Kasti", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Shell", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertCustomers(models.Customers{Username: "Aditya", Email: "as.com", Password: "123", Address: "a"})
	database.InsertCustomers(models.Customers{Username: "rein", Email: "ass.com", Password: "123", Address: "a"})
	database.InsertCartitems(models.Cartitems{Carts_id: 0, Products_id: 0, Quantity: 5})
	database.InsertCartitems(models.Cartitems{Carts_id: 0, Products_id: 1, Quantity: 5})
	database.InsertCartitems(models.Cartitems{Carts_id: 0, Products_id: 1, Quantity: 4})
	database.InsertCartitems(models.Cartitems{Carts_id: 1, Products_id: 1, Quantity: 3})

}

func TestCreateCartItemController(t *testing.T) {
	var testCases = []struct {
		code    int
		status  string
		json    string
		message string
	}{
		{
			200,
			"success",
			"{\r\n        \"carts_id\":1,\r\n        \"products_id\":1,\r\n        \"quantity\":3\r\n}",
			"success insert cartitems",
		}, {
			200,
			"success",
			"{\r\n        \"carts_id\":2,\r\n        \"products_id\":2,\r\n        \"quantity\":3\r\n}",
			"success insert cartitems",
		},
	}
	e := InitEcho()
	insertDBCart()

	for _, testcase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testcase.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/cartitems")
		if assert.NoError(t, CreateCartitemsController(c)) {
			body := rec.Body.String()
			var response models.Cartitems_response_single
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

func TestGetCartItemController(t *testing.T) {
	var testCases = []struct {
		code        int
		status      string
		message     string
		cartId      string
		ExpectedLen int
	}{
		{
			200,
			"success",
			"Success Get Cartitems",
			"0",
			3,
		}, {
			200,
			"success",
			"Success Get Cartitems",
			"1",
			1,
		},{
			400,
			"fail",
			"invalid id supplied",
			"Nadi",
			1,
		},
	}
	e := InitEcho()
	insertDBCart()

	for _, testCase := range testCases {
		q := make(url.Values)
		q.Set("cart", testCase.cartId)
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/cartitems")
		err:=GetCartitemsByCartIdController(c) 
	    if err==nil{
			assert.Equal(t, testCase.code, rec.Code)
			body := rec.Body.String()
			jsonData := []byte(body)
			var response models.CartItems_response_detail
			json.Unmarshal(jsonData, &response)
			assert.Equal(t, testCase.code, response.Code, "Code must be same")
			assert.Equal(t, testCase.message, response.Message, "Message must be same")
			assert.Equal(t, testCase.status, response.Status, "Status must be same")
			assert.Equal(t, testCase.ExpectedLen, len(response.Data), "Length of returned data must be same")
		}else{
			body := err.Error()
			jsonData := []byte(body)
			var response models.ErrorResponse
			json.Unmarshal(jsonData, &response)
			assert.Equal(t, testCase.code, response.Code, "Code must be same")
			assert.Equal(t, testCase.message, response.Message, "Message must be same")
			assert.Equal(t, testCase.status, response.Status, "Status must be same")
		}
	}
}

func TestDeleteCartItemController(t *testing.T){
	var testCases = []struct {
		code        int
		status      string
		message     string
		deleteId      string
	}{
		{
			200,
			"success",
			"cartitems succesfully deleted",
			"2",
		},
		{
			200,
			"success",
			"cartitems succesfully deleted",
			"1",
		},
	}
	e := InitEcho()
	insertDBCart()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/"+testCase.deleteId, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/cartitems")
		err:=GetCartitemsByCartIdController(c) 
	    if err==nil{
			assert.Equal(t, testCase.code, rec.Code)
			body := rec.Body.String()
			jsonData := []byte(body)
			var response models.SuccessResponse
			json.Unmarshal(jsonData, &response)
			assert.Equal(t, testCase.code, response.Code, "Code must be same")
			assert.Equal(t, testCase.message, response.Message, "Message must be same")
			assert.Equal(t, testCase.status, response.Status, "Status must be same")
		}
	}
}
