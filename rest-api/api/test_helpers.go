package api

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testdb) tearDown(t *testing.T) {
	dbName := os.Getenv("MONGO_DB_NAME")
	if err := tdb.client.Database(dbName).Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}

}

func setup(t *testing.T) *testdb {
	dburi := os.Getenv("MONGO_DB_URL_TEST")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
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

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}
