package utils

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractUserID(c *fiber.Ctx) (string, error) {
	user := c.Locals("user").(*jwt.Token)
	if user == nil {
		return "", errors.New("missing or invalid JWT token")
	}

	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	return userID, nil
}

func IsAuthenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, err := ExtractUserID(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		return c.Next()
	}
}
func GenerateToken(userID string, JWTSecret []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token expires in 72 hours

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
