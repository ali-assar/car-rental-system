package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReserveParams struct {
	FromDate time.Time `json:"fromDate"`
	tillDate time.Time `json:"tillDate"`
}

type CarHandler struct {
	store *db.Store
}

func NewCarHandler(store *db.Store) *CarHandler {
	return &CarHandler{
		store: store,
	}
}

func (h *CarHandler) HandleReserveCar(c *fiber.Ctx) error {
	var params ReserveParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}
	carID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}

	reserving := types.Reservation{
		UserID:   user.ID,
		CarID:    carID,
		FromDate: params.FromDate,
		TillDate: params.tillDate,
	}
	fmt.Println("%+v\n", reserving)
	return nil
}
