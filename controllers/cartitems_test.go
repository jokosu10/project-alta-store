package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-alta-store/lib/database"
	"project-alta-store/models"
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo"

)



func insertDBCart() {
	database.InsertCategories(models.Categories{Name: "Sabun", Description: "BersihBadan"})
	database.InsertCategories(models.Categories{Name: "Olahraga", Description: "BersihBadan"})
	database.InsertProducts(models.Products{Name: "Yonex23", Categories_id: 0, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Bola Kasti", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "Shell", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertCustomers(models.Customers{Username: "Aditya", Email: "as.com", Password: "123", Address: "a"})
	database.InsertCustomers(models.Customers{Username: "rein", Email: "ass.com", Password: "123", Address: "a"})
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
		},{
			200,
			"success",
			"{\r\n        \"carts_id\":2,\r\n        \"products_id\":2,\r\n        \"quantity\":3\r\n}",
			"success insert cartitems",
		},
	}
	e:= InitEcho()
	insertDBCart()

	for _,testcase := range testCases{
		req := httptest.NewRequest(http.MethodPost,"/",strings.NewReader(testcase.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req,rec)
		c.SetPath("/cartitems")
		if assert.NoError(t,CreateCartitemsController(c)){
			body := rec.Body.String()
			var response models.Cartitems_response_single
			err := json.Unmarshal([]byte(body),&response)
			fmt.Println(response)
			if err!=nil{
				fmt.Println("error unmarshall")
			} 
			assert.Equal(t,response.Code,testcase.code,"Code not expected")
			assert.Equal(t,response.Status,testcase.status,"Status not expected")
			assert.Equal(t,response.Message,testcase.message,"Message not expectred")

		}
	}
}
