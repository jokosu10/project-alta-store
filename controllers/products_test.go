package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"project-alta-store/config"
	"project-alta-store/lib/database"
	"project-alta-store/models"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

func insertDB() {
	database.InsertCategories(models.Categories{Name: "Sabun", Description: "BersihBadan"})
	database.InsertCategories(models.Categories{Name: "Olahraga", Description: "BersihBadan"})
	database.InsertProducts(models.Products{Name: "Yonex23", Categories_id: 0, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "redsew", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
	database.InsertProducts(models.Products{Name: "aliran air", Categories_id: 1, Description: "raket", Quantity: 5, Price: float32(52), Unit: "pcs"})
}

func TestGetProductController(t *testing.T) {
	var testCases = []struct {
		url         string
		code        int
		status      string
		expectedLen int
		isParam     bool
	}{
		{
			"/products",
			200,
			"success",
			3,
			false,
		},
		{
			"/products",
			200,
			"success",
			2,
			true,
		},
	}
	e := InitEcho()
	insertDB()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		if testCase.isParam {
			q := make(url.Values)
			q.Set("category", "1")
			req = httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		}
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.url)
		if assert.NoError(t, GetProductsController(c)) {
			assert.Equal(t, testCase.code, rec.Code)
			body := rec.Body.String()
			jsonData := []byte(body)
			var response models.Products_response
			json.Unmarshal(jsonData, &response)
			fmt.Println(response)
			assert.Equal(t, testCase.expectedLen, len(response.Data), "Product not selected based on category")
		}
	}
}

func TestErrorGetProductController(t *testing.T) {
	var testCases = []struct {
		url        string
		code       int
		status     string
		paramName  string
		paramValue string
	}{
		{
			"/products",
			400,
			"fail",
			"category",
			"nina",
		},
		{
			"/products",
			400,
			"fail",
			"category",
			"zzz",
		},
	}

	e := InitEcho()
	// insertDB()
	for _, testCase := range testCases {
		fmt.Println(testCase)
		q := make(url.Values)
		q.Set(testCase.paramName, testCase.paramValue)
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.url)
		fmt.Println(q.Encode())
		assert.Error(t, GetProductsController(c))
		// assert.Equal(t,testCase.code,rec.Code)
		// body := rec.Body.String()
		// jsonData := []byte(body)
		// // var response models.ErrorResponse
		// var response models.ErrorResponse
		// json.Unmarshal(jsonData,&response)
		// fmt.Println(body)
		// assert.Equal(t,testCase.code,response.Code,"Error code must be 400")
		// assert.Equal(t,testCase.status,response.Status,"Error status must be fail")

	}

}
