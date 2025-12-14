package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func ParseBody[T any](c fiber.Ctx) (*T, error) {
	var temp T
	if err := json.Unmarshal(c.Body(), &temp); err != nil {
		return nil, err
	}
	return &temp, nil
}

func ResponeConstructor(c fiber.Ctx, status int, res fiber.Map) error {
	c.SendStatus(status)

	return c.JSON(fiber.Map{
		"res":    res,
		"status": status,
	})
}
