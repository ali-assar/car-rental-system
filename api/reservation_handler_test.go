package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ali-Assar/car-rental-system/db/fixtures"
)

func TestAdminGetReservation(t *testing.T) {
	db := setup(t)
	defer db.tearDown(t)

	user := fixtures.AddUser(db.Store, "james", "foo", false)
	agency := fixtures.AddAgency(db.Store, "voom voom", "rome", 3, nil)
	car := fixtures.AddCar(db.Store, "sport", "petrol", "BMW", 2017, 100, agency.ID)

	from := time.Now()
	till := from.AddDate(0, 0, 5)
	reservation := fixtures.AddReservation(db.Store, user.ID, car.ID, from, till)
	fmt.Println(reservation)
}
