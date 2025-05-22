package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"go-project/controllers"
	"go-project/database"
	"go-project/models"
)

const pathProducts = "/products"
const pathProductsId = "/products/:id"
const pathTestId = "/products/1"

func setupTestServer() *echo.Echo {
	e := echo.New()
	database.ConnectTestDB()
	database.DB = database.DBTest

	e.GET(pathProducts, controllers.GetProducts)
	e.GET(pathProductsId, controllers.GetProduct)
	e.POST(pathProducts, controllers.CreateProduct)
	e.PUT(pathProductsId, controllers.UpdateProduct)
	e.DELETE(pathProductsId, controllers.DeleteProduct)

	return e
}

func TestCreateProduct(t *testing.T) {
	e := setupTestServer()

	product := models.Product{Name: "T-Shirt", Price: 75, CategoryId: 1}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, pathProducts, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetAllProducts(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, pathProducts, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TesGetProduct(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, pathTestId, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controllers.GetProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	e := setupTestServer()

	updated := models.Product{Name: "T-Shirt", Price: 80.5, CategoryId: 1}
	body, _ := json.Marshal(updated)
	req := httptest.NewRequest(http.MethodPut, pathTestId, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	_ = json.NewDecoder(rec.Body).Decode(&updated)

	if assert.NoError(t, controllers.UpdateProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteProduct(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodDelete, pathTestId, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controllers.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestGetInvalidId(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	_ = controllers.GetProduct(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)

}

func TestGetInvalidFormat(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	_ = controllers.GetProduct(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestCreateInvalidProduct(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodPost, pathProducts, bytes.NewBuffer([]byte(`{invalid}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = controllers.CreateProduct(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateInvalidProduct(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodPut, "/products/9998", bytes.NewBuffer([]byte(`{"name": "A"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	_ = controllers.UpdateProduct(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteNonExistingProduct(t *testing.T) {
	e := setupTestServer()

	req := httptest.NewRequest(http.MethodDelete, "/products/9997", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	_ = controllers.DeleteProduct(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
