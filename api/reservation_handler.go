package api

import (
	"net/http"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ReservationHandler struct {
	store *db.Store
}

func NewReservationHandler(store *db.Store) *ReservationHandler {
	return &ReservationHandler{
		store: store,
	}
}

func (h *ReservationHandler) HandleGetReservations(c *fiber.Ctx) error {
	reservation, err := h.store.Reservation.GetReservation(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(reservation)
}

func (h *ReservationHandler) HandleGetReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	reservation, err := h.store.Reservation.GetReservationByID(c.Context(), id)
	if err != nil {
		return err
	}
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}
	if reservation.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON((genericResp{
			Type: "error",
			Msg:  "not authorized",
		}))
	}
	return c.JSON(reservation)
}
