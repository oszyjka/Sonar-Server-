package controllers

import (
	"net/http"

	"go-project/database"
	"go-project/models"

	"github.com/labstack/echo/v4"
)

func CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := database.DB.Create(payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save payment"})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Payment received",
		"data":    payment,
	})
}
