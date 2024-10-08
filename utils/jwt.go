package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromJWT(c *fiber.Ctx) string {
	// Find id from jwt
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	return userId
}

func VerificationJWT(jwtToken string) (string, error) {
	// Verification
	secret_key := []byte(configs.NewConfig().JWTSecret)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret_key, nil
	})

	// Check JWT
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid && err == nil {
		return claims["user_id"].(string), nil
	} else {
		parts := strings.Split(err.Error(), ":")
		err = errors.New(parts[0])
		switch err.Error() {
		case jwt.ErrTokenMalformed.Error():
			return "", exceptions.ErrInvalidToken
		case jwt.ErrTokenInvalidClaims.Error():
			return "", exceptions.ErrExpiredToken
		default:
			return "", err
		}
	}
}

func GenerateJWT(user_id string) string {
	// Generate JWT token
	expireAt := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     expireAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(configs.NewConfig().JWTSecret))
	if err != nil {
		return err.Error()
	}

	return tokenString
}
