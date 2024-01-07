package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	tokens, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok || len(tokens) == 0 {
		return fmt.Errorf("unauthorized")
	}

	token := tokens[0]
	claims, err := validateToken(token)
	if err != nil {
		return fmt.Errorf("unauthorized")
	}
	expiresStr, ok := claims["expires"].(string)
	if !ok {
		return fmt.Errorf("invalid credentials")
	}
	expires, err := time.Parse(time.RFC3339, expiresStr)
	if err != nil {
		return fmt.Errorf("invalid credentials", err)
	}

	if time.Now().After(expires) {
		return fmt.Errorf("token expired")
	}
	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}

		secret := os.Getenv("JWT_SECRET")
		fmt.Println("never show secret", secret)
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("field to parse jwt token : ", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
