package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-alta-store/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var customerJSON = `{"id":1,"username":"test","email":"test@gmail.com","password":"123456789","address":"test address", "bank_name": "test bank address", "bank_account_number":"0000000000"}`

func TestRegisterCustomersController(t *testing.T) {
	var testCases = []struct {
		code    int
		status  string
		message string
	}{
		{
			200,
			"success",
			"success register customers",
		},
	}

	e := InitEcho()

	for _, testcase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(customerJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/register")
		if assert.NoError(t, RegisterCustomersController(c)) {
			body := rec.Body.String()
			var response models.SuccessResponse
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				fmt.Println("error unmarshall")
			}
			assert.Equal(t, response.Code, testcase.code, 200)
			assert.Equal(t, response.Status, testcase.status, "success")
			assert.Equal(t, response.Message, testcase.message, "success register customers")
		}
	}
}
