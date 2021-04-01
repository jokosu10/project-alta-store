package controllers

import (
	"net/http"
	"net/http/httptest"
	"project-alta-store/config"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEchoCustomer() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

var customerJSON = `{"id":1,"username":"test","email":"test@gmail.com","password":"123456789","address":"test address", "bank_name": "test bank address", "bank_account_number":"0000000000"}`

func TestRegisterCustomersController(t *testing.T) {
	url := "/register"
	e := InitEchoCustomer()
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(customerJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, RegisterCustomersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "["+customerJSON+"]", rec.Body.String())
	}
}
