package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go-project/database"
	"go-project/models"
)

func GetCategories(c echo.Context) error {
	var categories []models.Category
	result := database.DB.Preload("Products").Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	idCheck, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}
	var category models.Category
	result := database.DB.Preload("Products").First(&category, idCheck)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Category not found")
	}
	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, category)
}
