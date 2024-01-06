package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("jwt----------")

	tokens, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok || len(tokens) == 0 {
		return fmt.Errorf("unauthorized")
	}
	token := tokens[0]

	if err := parseToken(token); err != nil {
		return fmt.Errorf("unauthorized")
	}
	fmt.Println("token:", token)
	return nil
}

func parseToken(tokenStr string) error {
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
		return fmt.Errorf("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
	}
	return fmt.Errorf("unauthorized")
}
