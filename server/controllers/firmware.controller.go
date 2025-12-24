package controllers

import (
	firmwareDtos "firmware_server/dtos/firmware"
	"firmware_server/env/routes"
	"firmware_server/server"
	"firmware_server/services/firmwareService"
	"firmware_server/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func getFirmwares(c fiber.Ctx) error {
	var id *uint
	var product_id *uint
	idQuery := c.Query("id")
	productIdQuery := c.Query("product_id")

	if productIdQuery == "" {
		return utils.BadRequestResponse(c, "invalid query values")
	} else {
		value, err := strconv.Atoi(productIdQuery)
		if err != nil {
			return utils.BadRequestResponse(c,
				"invalid id query value",
			)
		}
		uintValue := uint(value)
		product_id = &uintValue
	}

	if idQuery != "" {
		value, err := strconv.Atoi(idQuery)
		if err != nil {
			return utils.BadRequestResponse(c,
				"invalid id query value",
			)
		}

		uintValue := uint(value)
		id = &uintValue
	}
	println(id)
	firmwares, err := firmwareService.GetFirmwares(*product_id, id)

	if err != nil {
		return err
	}

	return utils.ResponseConstructor(c, fiber.StatusOK, firmwares,nil)
}

// func updateFirmwareVersion(){

// }

func addFirmware(c fiber.Ctx) error {
	body, err := utils.ParseBody[firmwareDtos.AddFirmwareBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	err = firmwareService.AddFirmware(*body)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponseConstructor(c, fiber.StatusOK, "Firmware added successfully",nil)
}

func updateFirmware(c fiber.Ctx) error {
	body, err := utils.ParseBody[firmwareDtos.UpdateFirmwareBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	err = firmwareService.UpdateFirmware(*body)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponseConstructor(c, fiber.StatusOK, "Firmware updated successfully",nil)
}

func deleteFirmware(c fiber.Ctx) error {
	body, err := utils.ParseBody[firmwareDtos.DeleteFirmwareBody](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	err = firmwareService.DeleteFirmware(*body)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponseConstructor(c, fiber.StatusOK, "firmware deleted successfully",nil)
}

func firmwareRoute() {
	firmwareGroup := server.App.Group(routes.Firmware)
	firmwareGroup.Get("", getFirmwares)
	firmwareGroup.Post("", addFirmware)
	firmwareGroup.Put("", updateFirmware)
	firmwareGroup.Delete("", deleteFirmware)
}
