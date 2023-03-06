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

func GetAuthCookies(accessToken, refreshToken string) (*fiber.Cookie, *fiber.Cookie) {

	accessCookie := &fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	refreshCookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(10 * 24 * time.Hour),
		HTTPOnly: true,
	}

	return accessCookie, refreshCookie
}
