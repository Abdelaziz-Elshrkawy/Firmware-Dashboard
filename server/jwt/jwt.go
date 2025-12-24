package appJwt

import (
	"firmware_server/env"
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	Secret     []byte
	Issuer     string
	Audience   string
	ExpireTime time.Duration
}

var JWT = JWTConfig{
	Secret:     []byte(env.JWTSecret),
	Issuer:     "bedo-firmware-api",
	Audience:   "bedo-firmware-client-side",
	ExpireTime: time.Hour * 24,
}

func GenerateJWT(extraClaims jwt.MapClaims) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"iat": now.Unix(),
		"exp": now.Add(JWT.ExpireTime).Unix(),
		"iss": JWT.Issuer,
		"aud": JWT.Audience,
	}

	maps.Copy(claims, extraClaims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT.Secret)
}

func ParseJWT(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return JWT.Secret, nil
	},
		jwt.WithAudience(JWT.Audience),
		jwt.WithIssuer(JWT.Issuer),
	)

	return token, claims, err
}
