package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(uId uint, jwtSecret []byte, jwtRefreshSecret []byte) (string, string) {

	accessToken := GenerateAccessClaims(uId, jwtSecret, 15*time.Minute)
	refreshToken := GenerateAccessClaims(uId, jwtRefreshSecret, 30*24*time.Hour)
	return accessToken, refreshToken
}

func GenerateAccessClaims(uId uint, jwtSecret []byte, timeNumber time.Duration) string {
	t := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uId
	claims["exp"] = t.Add(timeNumber).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseToken(tokenString string, jwtSecret []byte) (uint, error) {
	var userId uint
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return userId, fiber.NewError(fiber.StatusUnauthorized, "invalid or missing token")
	}
	if !token.Valid {
		return userId, fiber.NewError(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userId, fiber.NewError(fiber.StatusUnauthorized)
	}
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return userId, fiber.NewError(fiber.StatusUnauthorized, "token is out of date")
	}
	userId = uint(claims["sub"].(float64))
	return userId, nil
}
