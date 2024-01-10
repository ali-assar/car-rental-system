package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/Ali-Assar/car-rental-system/api/middleware"
	"github.com/Ali-Assar/car-rental-system/db/fixtures"
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
)

func TestUserGetReservation(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	var (
		//nonAuthUser        = fixtures.AddUser(db.Store, "james", "wrong", false)
		user               = fixtures.AddUser(db.Store, "james", "foo", false)
		agency             = fixtures.AddAgency(db.Store, "voom voom", "rome", 3, nil)
		car                = fixtures.AddCar(db.Store, "sport", "petrol", "BMW", 2017, 100, agency.ID)
		from               = time.Now()
		till               = from.AddDate(0, 0, 5)
		reservation        = fixtures.AddReservation(db.Store, user.ID, car.ID, from, till)
		app                = fiber.New()
		route              = app.Group("/", middleware.JWTAuthentication(db.User))
		reservationHandler = NewReservationHandler(db.Store)
	)

	route.Get("/:id", reservationHandler.HandleGetReservation)
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", reservation.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 but got, %d", resp.StatusCode)
	}
	var reservationResp *types.Reservation
	if err := json.NewDecoder(resp.Body).Decode(&reservationResp); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(reservation.ID, reservationResp.ID) {
		t.Fatalf("the inserted reservation ID is %s\n but the fetched reservation ID is %s\n", reservation.ID, reservationResp.ID)
	}
}

func TestAdminGetReservation(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)
	var (
		adminUser          = fixtures.AddUser(db.Store, "admin", "admin", true)
		user               = fixtures.AddUser(db.Store, "james", "foo", false)
		agency             = fixtures.AddAgency(db.Store, "voom voom", "rome", 3, nil)
		car                = fixtures.AddCar(db.Store, "sport", "petrol", "BMW", 2017, 100, agency.ID)
		from               = time.Now()
		till               = from.AddDate(0, 0, 5)
		reservation        = fixtures.AddReservation(db.Store, user.ID, car.ID, from, till)
		app                = fiber.New()
		admin              = app.Group("/", middleware.JWTAuthentication(db.User), middleware.AdminAuth)
		reservationHandler = NewReservationHandler(db.Store)
	)
	_ = reservation
	admin.Get("/", reservationHandler.HandleGetReservations)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 but got, %d", resp.StatusCode)
	}

	var reservations []*types.Reservation
	if err := json.NewDecoder(resp.Body).Decode(&reservations); err != nil {
		fmt.Println("Error decoding response body:", err)
		t.Fatal(err)
	}
	if len(reservations) != 1 {
		t.Fatalf("expected at least 1 reservation but got %d", len(reservations))
	}
	if !reflect.DeepEqual(reservation.ID, reservations[0].ID) {
		t.Fatalf("the inserted reservation ID is %s\n but the fetched reservation ID is %s\n", reservation.ID, reservations[0].ID)
	}

	// testing  reservations cannot be fetched by a user
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))

	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected not 200 code but got, %d", http.StatusOK)
	}
}
