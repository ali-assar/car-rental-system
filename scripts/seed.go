package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	agencyStore := db.NewMongoAgencyStore(client, db.DBNAME)
	carStore := db.NewMongoCarStore(client, db.DBNAME)

	agency := types.Agency{
		Name:     "DrivingPartner",
		Location: "Rome",
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
		insertedCar, err := carStore.InsertCar(ctx, &car)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedCar)
	}

}
