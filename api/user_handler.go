package api

import (
	"github.com/Ali-Assar/reservation-system/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {

	u := types.User{
		FirstName: "ali",
		LastName:  "assar",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("ali")
}
