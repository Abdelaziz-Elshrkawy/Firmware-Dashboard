package utils

import (
	"github.com/gofiber/fiber/v3"
)

func ParseBody[T any](c fiber.Ctx) (*T, error) {
	var temp T
	// if err := json.Unmarshal(c.Body(), &temp); err != nil {
	// 	return nil, err
	// }

	if err := c.Bind().Body(&temp); err != nil {
		return nil, err
	}

	// if err := dtos.DtoVlidator.Struct(temp); err != nil {
	// 	validationErrors := dtos.FormatValidationErrors(err)
	// 	return nil, validationErrors
	// }

	return &temp, nil
}

func ResponeConstructor(c fiber.Ctx, status int, res any) error {
	c.SendStatus(status)

	return c.JSON(fiber.Map{
		"res":    res,
		"status": status,
	})
}

func BadRequestResponse(c fiber.Ctx, msg string) error {
	return ResponeConstructor(c, fiber.StatusBadRequest, fiber.Map{
		"error": msg,
	})
}
