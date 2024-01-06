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
	GetCars(context.Context, bson.M) ([]*types.Car, error)
}

type MongoCarStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	AgencyStore
}

func NewMongoCarStore(client *mongo.Client, agencyStore AgencyStore) *MongoCarStore {
	return &MongoCarStore{
		client:      client,
		coll:        client.Database(DBNAME).Collection("cars"),
		AgencyStore: agencyStore,
	}
}

func (s *MongoCarStore) GetCars(ctx context.Context, filter bson.M) ([]*types.Car, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var cars []*types.Car
	if err := resp.All(ctx, &cars); err != nil {
		return nil, err
	}
	return cars, nil

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
	if err := s.AgencyStore.UpdateAgency(ctx, filter, update); err != nil {
		return nil, err
	}
	return car, nil
}
