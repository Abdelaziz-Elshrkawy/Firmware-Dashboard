package controllers

import (
	"firmware_server/env/routes"
	"firmware_server/server"
	"firmware_server/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func getDevice(c fiber.Ctx) error {
	var id *int
	query := c.Query("id")

	if query != "" {
		value, err := strconv.Atoi(query)
		if err != nil {
			return utils.ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
				"error": "invalid id query value",
			})
		}

		id = &value
	}

	// devices, sqlErr :=

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{
		"message": "",
	})
}

func devicesRoute() {
	var deviceGroup = server.App.Group(routes.Device)
	deviceGroup.Get("", getDevice)
}
