package db

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarStore interface {
	InsertCar(context.Context, *types.Car) (*types.Car, error)
}

type MongoCarStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCarStore(client *mongo.Client, dbname string) *MongoCarStore {
	return &MongoCarStore{
		client: client,
		coll:   client.Database(dbname).Collection("cars"),
	}
}

func (s *MongoCarStore) InsertCar(ctx context.Context, car *types.Car) (*types.Car, error) {
	resp, err := s.coll.InsertOne(ctx, car)
	if err != nil {
		return nil, err
	}
	car.ID = resp.InsertedID.(primitive.ObjectID)

	//update the hotel with this room id
	filter := bson.M{"_id": car.AgencyID}
	update := bson.M{"$push": bson.M{"cars": car.ID}}
	return car, nil
}
