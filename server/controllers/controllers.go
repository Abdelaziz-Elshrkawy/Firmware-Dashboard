package controllers

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app fiber.Router){
	ProductRoute(app)
}