package utils

import (
	"firmware_server/env"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func ParseBody[T any](c fiber.Ctx) (*T, error) {
	var temp T
	// if err := json.Unmarshal(c.Body(), &temp); err != nil {
	// 	return nil, err
	// }

	if err := c.Bind().Body(&temp); err != nil {
		return nil, err
	}

	// if err := dtos.DtoValidator.Struct(temp); err != nil {
	// 	validationErrors := dtos.FormatValidationErrors(err)
	// 	return nil, validationErrors
	// }

	return &temp, nil
}

func ResponseConstructor(c fiber.Ctx, status int, res any, cookies []fiber.Cookie) error {
	c.Status(status)

	for i := range cookies {
		c.Cookie(&cookies[i])
	}

	return c.JSON(fiber.Map{
		"res":    res,
		"status": status,
	})
}

func BadRequestResponse(c fiber.Ctx, msg string) error {
	return ResponseConstructor(c, fiber.StatusBadRequest, fiber.Map{
		"error": msg,
	}, nil)
}

func CreatePassword(passwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd+env.PasswordSecret), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

func ComparePassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password+env.PasswordSecret))
}
