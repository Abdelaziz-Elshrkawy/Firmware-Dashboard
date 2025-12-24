package controllers

import (
	authDtos "firmware_server/dtos/auth"
	"firmware_server/env"
	appJwt "firmware_server/jwt"
	"firmware_server/server"
	authService "firmware_server/services/auth"
	"firmware_server/utils"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func login(c fiber.Ctx) error {
	body, err := utils.ParseBody[authDtos.Creds](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	user, err := authService.Login(*body)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	token, err := appJwt.GenerateJWT(jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
	})

	return utils.ResponseConstructor(c, fiber.StatusOK, fiber.Map{"username": user.Username}, []fiber.Cookie{
		{
			Name:     env.JWTCookieName,
			Expires:  time.Now().Add(time.Hour * 24),
			Path:     "/",
			SameSite: fiber.CookieSameSiteLaxMode,
			HTTPOnly: true,
			Secure:   true,
			Value:    token,
		},
	})
}

func signup(c fiber.Ctx) error {
	body, err := utils.ParseBody[authDtos.Creds](c)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	_, err = authService.SignUp(*body)

	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.ResponseConstructor(c, fiber.StatusOK, fiber.Map{
		"res": "Registered",
	}, nil)
}

func authRoute() {
	var AuthGroup = server.App.Group("auth")
	AuthGroup.Post("login", login)
	AuthGroup.Post("signup", signup)
}
