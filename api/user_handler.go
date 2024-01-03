package api

import (
	"context"
	"log"

	"github.com/Ali-Assar/reservation-system/db"
	"github.com/Ali-Assar/reservation-system/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	u := types.User{
		FirstName: "ali",
		LastName:  "assar",
	}
	return c.JSON(u)
}
