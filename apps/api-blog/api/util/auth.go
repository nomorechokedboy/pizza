package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(uId uint, jwtSecret []byte, jwtRefreshSecret []byte) (string, string) {

	accessToken := GenerateAccessClaims(uId, jwtSecret)
	refreshToken := GenerateRefreshClaims(uId, jwtRefreshSecret)
	return accessToken, refreshToken
}

func GenerateAccessClaims(uId uint, jwtSecret []byte) string {
	t := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uId
	claims["exp"] = t.Add(15 * time.Minute).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func GenerateRefreshClaims(uId uint, jwtSecret []byte) string {

	t := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = uId
	claims["exp"] = t.Add(30 * 24 * time.Hour).Unix()
	refreshTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}
	return refreshTokenString

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

func RefreshAccessToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("refresh-secret"), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	expirationtime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationtime) {
		return fiber.NewError(fiber.StatusUnauthorized, "token is out of date")
	}
	uId := uint(claims["sub"].(float64))
	accessToken := GenerateAccessClaims(uId, []byte("my-secret"))
	// accessCookie, _ := GetAuthCookies(accessToken, tokenString)
	return c.JSON(accessToken)
}
