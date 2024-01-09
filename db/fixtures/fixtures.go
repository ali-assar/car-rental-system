package fixtures

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddReservation(store db.Store, userID, carID primitive.ObjectID, from, till time.Time) *types.Reservation {
	reservation := &types.Reservation{
		UserID:   userID,
		CarID:    carID,
		FromDate: from,
		TillDate: till,
	}
	insertedReservation, err := store.Reservation.InsertReservation(context.Background(), reservation)
	if err != nil {
		log.Fatal(err)
	}
	return insertedReservation
}

func AddCar(store *db.Store, carType, fuel, model string, year int, price float64, agencyID primitive.ObjectID) *types.Car {
	car := &types.Car{
		Type:     carType,
		Fuel:     fuel,
		Model:    model,
		Year:     year,
		Price:    price,
		AgencyID: agencyID,
	}
	insertedCar, err := store.Car.InsertCar(context.Background(), car)
	if err != nil {
		log.Fatal(err)
	}
	return insertedCar
}

func AddAgency(store *db.Store, name, loc string, rating int, cars []primitive.ObjectID) *types.Agency {
	var carIDs = cars
	if cars == nil {
		carIDs = []primitive.ObjectID{}
	}
	agency := types.Agency{
		Name:     name,
		Location: loc,
		Cars:     carIDs,
		Rating:   rating,
	}
	insertedAgency, err := store.Agency.InsertAgency(context.TODO(), &agency)
	if err != nil {
		log.Fatal(err)
	}
	return insertedAgency
}

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}
