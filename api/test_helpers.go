package api

import (
	"context"
	"log"
	"testing"

	"github.com/Ali-Assar/car-rental-system/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

type testdb struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testdb) tearDown(t *testing.T) {
	if err := tdb.client.Database(db.DBNAME).Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}
	agencyStore := db.NewMongoAgencyStore(client)
	return &testdb{
		client: client,
		Store: &db.Store{
			Agency:      agencyStore,
			User:        db.NewMongoUserStore(client),
			Car:         db.NewMongoCarStore(client, agencyStore),
			Reservation: db.NewMongoReservationStore(client),
		},
	}
}
