package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"go-project/database"
	"go-project/models"
)

func GetCarts(c echo.Context) error {
	var carts []models.Cart
	result := database.DB.Find(&carts)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, carts)
}

func GetCart(c echo.Context) error {
	idCheck, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id"})
	}
	var cart models.Cart
	result := database.DB.First(&cart, idCheck)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	return c.JSON(http.StatusOK, cart)
}

func CreateCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	result := database.DB.Create(cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusCreated, cart)
}
