package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-project/database"
	"go-project/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const pathPayment = "/payments"

func setupPaymentRouter() *echo.Echo {
	e := echo.New()
	database.ConnectTestDB()
	database.DB = database.DBTest
	e.POST(pathPayment, CreatePayment)
	return e
}

func TestCreatePayment(t *testing.T) {
	e := setupPaymentRouter()

	payment := models.Payment{
		Amount:  100.0,
		Method:  "Card",
		CartID:  1657009,
		Comment: "Payment for bill nr 14567",
	}
	jsonPayment, _ := json.Marshal(payment)
	req := httptest.NewRequest(http.MethodPost, pathPayment, bytes.NewBuffer(jsonPayment))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreatePayment(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

}

func TestCreateInvalidPayment(t *testing.T) {
	e := setupPaymentRouter()

	req := httptest.NewRequest(http.MethodPost, pathPayment, bytes.NewBufferString("{invalid}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreatePayment(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
