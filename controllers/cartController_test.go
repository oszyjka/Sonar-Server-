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

const pathCarts = "/carts"

func setupCartRouter() *echo.Echo {
	e := echo.New()
	database.ConnectTestDB()
	database.DB = database.DBTest
	e.GET(pathCarts, GetCarts)
	e.GET("/carts/:id", GetCart)
	e.POST(pathCarts, CreateCart)
	return e
}

func TestCreateCart(t *testing.T) {
	e := setupCartRouter()

	cart := models.Cart{User: "Ana", Total: 150}
	jsonCart, _ := json.Marshal(cart)
	req := httptest.NewRequest(http.MethodPost, pathCarts, bytes.NewBuffer(jsonCart))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestAllCarts(t *testing.T) {
	e := setupCartRouter()

	req := httptest.NewRequest(http.MethodGet, pathCarts, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetCarts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetCart(t *testing.T) {
	e := setupCartRouter()
	cart := models.Cart{User: "Ana", Total: 150}
	jsonCart, _ := json.Marshal(cart)
	req := httptest.NewRequest(http.MethodPost, pathCarts, bytes.NewBuffer(jsonCart))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = CreateCart(c)

	id := "1"
	req = httptest.NewRequest(http.MethodGet, "/carts/1", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(id)
	err := GetCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateInvalidCart(t *testing.T) {
	e := setupCartRouter()

	req := httptest.NewRequest(http.MethodPost, pathCarts, bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := CreateCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetInvalidCartId(t *testing.T) {
	e := setupCartRouter()

	req := httptest.NewRequest(http.MethodGet, "/carts/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")
	err := GetCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetNonExistingCart(t *testing.T) {
	e := setupCartRouter()

	req := httptest.NewRequest(http.MethodGet, "/carts/999999", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999999")
	err := GetCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
