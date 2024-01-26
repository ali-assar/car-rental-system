package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokens, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok || len(tokens) == 0 {
			return ErrAuthorization()
		}

		token := tokens[0]
		claims, err := validateToken(token)
		if err != nil {
			return ErrAuthorization()
		}
		expiresStr, ok := claims["expires"].(string)
		if !ok {
			return fmt.Errorf("invalid credentials")
		}
		expires, err := time.Parse(time.RFC3339, expiresStr)
		if err != nil {
			return fmt.Errorf("invalid credentials %s", err)
		}

		if time.Now().After(expires) {
			return NewError(http.StatusUnauthorized, "token expired")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return ErrAuthorization()
		}
		//set the current authenticated user to the context
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrAuthorization()
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("field to parse jwt token : ", err)
		return nil, ErrAuthorization()
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrAuthorization()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrAuthorization()
	}
	return claims, nil
}
