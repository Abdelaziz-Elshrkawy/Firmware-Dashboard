package middleware

import (
	"errors"
	"firmware_server/env"
	appJwt "firmware_server/jwt"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJwt(c fiber.Ctx) (jwt.Token, error) {
	jwt := c.Cookies(env.JWTCookieName)
	if jwt == "" {
		return jwt.Token{}, errors.New()
	}
	appJwt.ParseJWT()
}
