package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReserveParams struct {
	FromDate time.Time `json:"fromDate"`
	TillDate time.Time `json:"tillDate"`
}

type CarHandler struct {
	store *db.Store
}

func (p ReserveParams) validate() error {
	now := time.Now()
	if now.After(p.FromDate) || now.After(p.TillDate) {
		return fmt.Errorf("cannot reserve in the past")
	}

	// Date Range Validation
	if p.FromDate.After(p.TillDate) {
		return fmt.Errorf("fromDate must be before tillDate")
	}

	// Future Date Validation
	if p.FromDate.Before(now) || p.TillDate.Before(now) {
		return fmt.Errorf("reservation dates must be in the future")
	}

	// Minimum Reservation Duration
	minDuration := 24 * time.Hour
	if p.TillDate.Sub(p.FromDate) < minDuration {
		return fmt.Errorf("minimum reservation duration is 24 hours")
	}

	return nil
}

func NewCarHandler(store *db.Store) *CarHandler {
	return &CarHandler{
		store: store,
	}
}

func (h *CarHandler) HandleGetCars(c *fiber.Ctx) error {
	cars, err := h.store.Car.GetCars(c.Context(), db.Map{})
	if err != nil {
		return err
	}
	return c.JSON(cars)
}

func (h *CarHandler) HandleReserveCar(c *fiber.Ctx) error {
	var params ReserveParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.validate(); err != nil {
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

	ok, err = h.isCarAvailableForBooking(c.Context(), carID, params)
	if err != nil {
		return err
	}

	if !ok {
		return c.Status(http.StatusBadRequest).JSON(genericResp{
			Type: "error",
			Msg:  fmt.Sprintf("car %s is already reserved", c.Params("id")),
		})
	}

	reservation := types.Reservation{
		UserID:   user.ID,
		CarID:    carID,
		FromDate: params.FromDate,
		TillDate: params.TillDate,
	}

	insertedReservation, err := h.store.Reservation.InsertReservation(c.Context(), &reservation)
	if err != nil {
		return err
	}

	return c.JSON(insertedReservation)
}

func (h *CarHandler) isCarAvailableForBooking(ctx context.Context, carID primitive.ObjectID, params ReserveParams) (bool, error) {
	filter := bson.M{
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}

	reservations, err := h.store.Reservation.GetReservation(ctx, filter)
	if err != nil {
		return false, err
	}

	return len(reservations) == 0, nil
}
