package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ali-Assar/car-rental-system/api"
	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client           *mongo.Client
	carStore         db.CarStore
	agencyStore      db.AgencyStore
	userStore        db.UserStore
	reservationStore db.ReservationStore
	ctx              = context.Background()
)

func seedUser(isAdmin bool, fname, lname, email, password string) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     email,
		FirstName: fname,
		LastName:  lname,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin
	insertedUser, err := userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
	return insertedUser
}

func seedCar(carType, fuel, model string, year int, price float64, agencyID primitive.ObjectID) *types.Car {

	car := &types.Car{
		Type:     carType,
		Fuel:     fuel,
		Model:    model,
		Year:     year,
		Price:    price,
		AgencyID: agencyID,
	}
	insertedCar, err := carStore.InsertCar(context.Background(), car)
	if err != nil {
		log.Fatal(err)
	}
	return insertedCar
}

func seedReservation(userID, carID primitive.ObjectID, from, till time.Time) {
	reservation := &types.Reservation{
		UserID:   userID,
		CarID:    carID,
		FromDate: from,
		TillDate: till,
	}
	resp, err := reservationStore.InsertReservation(ctx, reservation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("booking: ", resp.ID)
}

func seedAgency(name, location string, rating int) *types.Agency {
	agency := types.Agency{
		Name:     name,
		Location: location,
		Cars:     []primitive.ObjectID{},
		Rating:   rating,
	}
	insertedAgency, err := agencyStore.InsertAgency(ctx, &agency)
	if err != nil {
		log.Fatal(err)
	}
	return insertedAgency
}

func main() {
	james := seedUser(false, "james", "foo", "james@foo.com", "supersafe")
	seedUser(true, "admin", "admin", "admin@admin.com", "adminsafe")

	agency1 := seedAgency("Driving Partner", "Rome", 3)
	seedAgency("Car Bank", "Milan", 5)
	seedAgency("Go voom voom", "Paris", 2)

	car1 := seedCar("sport", "petrol", "BMW", 2020, 100, agency1.ID)
	seedCar("economy", "hybrid", "toyota", 2017, 20, agency1.ID)
	seedCar("luxury", "petrol", "benz G-wagon", 2022, 200, agency1.ID)
	seedReservation(james.ID, car1.ID, time.Now(), time.Now().AddDate(0, 0, 2))
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	agencyStore = db.NewMongoAgencyStore(client)
	carStore = db.NewMongoCarStore(client, agencyStore)
	userStore = db.NewMongoUserStore(client)
	reservationStore = db.NewMongoReservationStore(client)
}
