package controllers

import (
	"firmware_server/server"
	"firmware_server/services/productService"

	"github.com/gofiber/fiber/v3"
)

func getProducts(c fiber.Ctx) error {
	products, err := productService.GetProducts()
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(products)
}

func ProductRoute() {
	var ProductGroup = server.App.Group("/products")
	ProductGroup.Get("", getProducts)
}
