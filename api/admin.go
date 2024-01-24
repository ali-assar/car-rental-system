package api

import (
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrAuthorization()
	}
	if !user.IsAdmin {
		return ErrAuthorization()
	}
	return c.Next()
}
