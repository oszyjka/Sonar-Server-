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

const pathCategories = "/categories"

func setupCategoryTestServer() *echo.Echo {
	e := echo.New()
	database.ConnectTestDB()
	database.DB = database.DBTest

	e.GET(pathCategories, controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.POST(pathCategories, controllers.CreateCategory)

	return e
}

func TestCreateCategorie(t *testing.T) {
	e := setupCategoryTestServer()

	category := models.Category{Name: "Clothes"}
	body, _ := json.Marshal(category)
	req := httptest.NewRequest(http.MethodPost, pathCategories, bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateCategory(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetAllCategories(t *testing.T) {
	e := setupCategoryTestServer()

	req := httptest.NewRequest(http.MethodGet, pathCategories, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetCategories(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TesGetCategorie(t *testing.T) {
	e := setupCategoryTestServer()

	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, controllers.GetCategory(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestGetNonExistingCategory(t *testing.T) {
	e := setupCategoryTestServer()

	req := httptest.NewRequest(http.MethodGet, "/categories/999", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("999")

	_ = controllers.GetCategory(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)

}

func TestGetInvalidCategory(t *testing.T) {
	e := setupCategoryTestServer()

	req := httptest.NewRequest(http.MethodGet, "/categories/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("abc")

	_ = controllers.GetCategory(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestInvalidCreateCategory(t *testing.T) {
	e := setupCategoryTestServer()

	req := httptest.NewRequest(http.MethodPost, pathCategories, bytes.NewBuffer([]byte(`{invalid}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = controllers.CreateCategory(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
