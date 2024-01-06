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
	ctx         = context.Background()
)

func seedAgency(name, location string, rating int) {
	agency := types.Agency{
		Name:     name,
		Location: location,
		Cars:     []primitive.ObjectID{},
		Rating:   rating,
	}
	cars := []types.Car{
		{Type: types.EconomyCarsType,
			BasePrice: 20,
		},
		{Type: types.SportCarsType,
			BasePrice: 500,
		},
		{Type: types.CargoVanType,
			BasePrice: 50,
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
}
