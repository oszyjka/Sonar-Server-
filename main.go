package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-project/controllers"
	"go-project/database"
)

func main() {
	e := echo.New()
	database.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	const pathProductsId = "/products/:id"

	e.GET("/products", controllers.GetProducts)
	e.GET(pathProductsId, controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT(pathProductsId, controllers.UpdateProduct)
	e.DELETE(pathProductsId, controllers.DeleteProduct)

	e.GET("/carts", controllers.GetCarts)
	e.GET("/carts/:id", controllers.GetCart)
	e.POST("/carts", controllers.CreateCart)

	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.POST("/categories", controllers.CreateCategory)

	e.POST("/payments", controllers.CreatePayment)

	e.Logger.Fatal(e.Start(":8080"))
}
