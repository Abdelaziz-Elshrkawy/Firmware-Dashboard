package controllers

import (
	"firmware_server/services"

	"github.com/gofiber/fiber/v3"
)



func getProducts() rune {
	return services.UserService{}.GetProducts()
}



func ProductRoute(app fiber.Router) {
	var ProductRoute = fiber.New().Group("/products")
	ProductRoute.Get("", getProducts)
}