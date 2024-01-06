package api

import (
	"fmt"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/gofiber/fiber/v2"
)

type AgencyHandler struct {
	agencyStore db.AgencyStore
	carStore    db.CarStore
}

func NewAgencyHandler(as db.AgencyStore, cs db.CarStore) *AgencyHandler {
	return &AgencyHandler{
		agencyStore: as,
		carStore:    cs,
	}
}

type AgencyQueryParams struct {
	Cars   bool
	Rating int
}

func (a *AgencyHandler) HandleGetAgencies(c *fiber.Ctx) error {
	var qparams AgencyQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}

	agencies, err := a.agencyStore.GetAgencies(c.Context(), nil)
	if err != nil {
		return err
	}
	fmt.Println(qparams)
	return c.JSON(agencies)
}
