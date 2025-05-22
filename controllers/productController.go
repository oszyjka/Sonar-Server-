package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go-project/database"
	"go-project/models"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	idCheck, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}
	var product models.Product
	result := database.DB.First(&product, idCheck)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	result := database.DB.Create(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	idCheck, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}
	var product models.Product
	result := database.DB.First(&product, idCheck)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	updatedProduct := new(models.Product)
	if err := c.Bind(updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.CategoryId = updatedProduct.CategoryId
	database.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	idCheck, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}

	result := database.DB.Delete(&models.Product{}, idCheck)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.NoContent(http.StatusNoContent)
}
