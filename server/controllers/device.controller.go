package controllers

import (
	deviceDtos "firmware_server/dtos/device"
	"firmware_server/env/routes"
	"firmware_server/server"
	"firmware_server/services/deviceService"
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
			return utils.BadRequestResponse(c,
				"invalid id query value",
			)
		}

		id = &value
	}

	devices, sqlErr := deviceService.GetDevice(id)

	if sqlErr != nil {
		return utils.BadRequestResponse(c, sqlErr.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, devices)
}

func addDevice(c fiber.Ctx) error {
	body, err := utils.ParseBody[deviceDtos.AddDeviceBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if body.Serial == nil || body.Product_Id == nil || body.Firmware_Id == nil {
		return utils.BadRequestResponse(c, "invalid inputs please provide all needed inputs")
	}

	if err := deviceService.AddDevice(*body.Serial, *body.Product_Id, *body.Firmware_Id); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusCreated, "Device added successfully")

}

func updateDevice(c fiber.Ctx) error {
	body, err := utils.ParseBody[deviceDtos.UpdateDeviceBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	err = deviceService.UpdateDevice(*body.Id, &body.Serial, &body.Product_Id, &body.Firmware_Id)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}
	return utils.ResponeConstructor(c, fiber.StatusOK, "Device updated successfully")
}

func deleteDevice(c fiber.Ctx) error {
	body, err := utils.ParseBody[deviceDtos.DeleteDeviceBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if body.Id == nil {
		return utils.BadRequestResponse(c, "id cannot be empty")
	}

	err = deviceService.DeleteDevice(*body.Id)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponeConstructor(c, fiber.StatusOK, fiber.Map{
		"message": "Device deleted successfully",
	})
}

func devicesRoute() {
	var deviceGroup = server.App.Group(routes.Device)
	deviceGroup.Get("", getDevice)
	deviceGroup.Post("", addDevice)
	deviceGroup.Put("", updateDevice)
	deviceGroup.Delete("", deleteDevice)
}
