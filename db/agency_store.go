package db

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AgencyStore interface {
	InsertAgency(context.Context, *types.Agency) (*types.Agency, error)
	UpdateAgency(context.Context, bson.M, bson.M) error
}

type MongoAgencyStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoAgencyStore(client *mongo.Client, dbname string) *MongoAgencyStore {
	return &MongoAgencyStore{
		client: client,
		coll:   client.Database(dbname).Collection("agency"),
	}
}

func (s *MongoAgencyStore) UpdateAgency(ctx context.Context, filter, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoAgencyStore) InsertAgency(ctx context.Context, agency *types.Agency) (*types.Agency, error) {
	resp, err := s.coll.InsertOne(ctx, agency)
	if err != nil {
		return nil, err
	}
	agency.ID = resp.InsertedID.(primitive.ObjectID)
	return agency, nil
}
