package main

import (
	"context"
	"log"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	carStore    db.CarStore
	agencyStore db.AgencyStore
	userStore   db.UserStore
	ctx         = context.Background()
)

func seedUser(isAdmin bool, fname, lname, email, password string) {
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
	_, err = userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func seedAgency(name, location string, rating int) {
	agency := types.Agency{
		Name:     name,
		Location: location,
		Cars:     []primitive.ObjectID{},
		Rating:   rating,
	}
	cars := []types.Car{
		{Type: "muscle",
			Fuel:  "petrol",
			Year:  1999,
			Model: "ford mustang",
			Price: 200,
		},
		{Type: "economy",
			Fuel:  "hybrid",
			Year:  2005,
			Model: "toyota prius",
			Price: 20,
		},
		{Type: "luxury",
			Fuel:  "petrol",
			Year:  2023,
			Model: "mercedes benz G wagon",
			Price: 400,
		},
	}
	insertedAgency, err := agencyStore.InsertAgency(ctx, &agency)
	if err != nil {
		log.Fatal(err)
	}
	for _, car := range cars {
		car.AgencyID = insertedAgency.ID
		_, err := carStore.InsertCar(ctx, &car)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	seedAgency("Driving Partner", "Rome", 3)
	seedAgency("Car Bank", "Milan", 5)
	seedAgency("Go voom voom", "Paris", 2)
	seedUser(false, "james", "foo", "james@foo.com", "supersafe")
	seedUser(true, "admin", "admin", "admin@admin.com", "adminsafe")

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
}
