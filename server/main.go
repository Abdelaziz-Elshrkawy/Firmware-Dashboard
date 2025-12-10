package main

import (
	"firmware_server/controllers"
	"firmware_server/database"
	"fmt"

	fiber "github.com/gofiber/fiber/v3"
)

func main() {

	// initializing database
	database.Connect()


	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {

		return c.SendString("Hello")
	})

	controllers.RegisterControllers(app)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
