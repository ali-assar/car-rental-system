package api

import (
	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AgencyHandler struct {
	store *db.Store
}

func NewAgencyHandler(store *db.Store) *AgencyHandler {
	return &AgencyHandler{
		store: store,
	}
}

func (a *AgencyHandler) HandleGetCars(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"agencyID": oid}
	cars, err := a.store.Car.GetCars(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(cars)
}

func (a *AgencyHandler) HandleGetAgency(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	agency, err := a.store.Agency.GetAgencyByID(c.Context(), oid)
	if err != nil {
		return err
	}
	return c.JSON(agency)
}

func (a *AgencyHandler) HandleGetAgencies(c *fiber.Ctx) error {
	agencies, err := a.store.Agency.GetAgencies(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(agencies)
}
