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
	GetAgencies(context.Context, bson.M) ([]*types.Agency, error)
	GetAgencyByID(context.Context, primitive.ObjectID) (*types.Agency, error)
}

type MongoAgencyStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoAgencyStore(client *mongo.Client) *MongoAgencyStore {
	return &MongoAgencyStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("agency"),
	}
}

func (s *MongoAgencyStore) GetAgencyByID(ctx context.Context, id primitive.ObjectID) (*types.Agency, error) {
	var agency types.Agency
	if err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&agency); err != nil {
		return nil, err
	}

	return &agency, nil
}

func (s *MongoAgencyStore) GetAgencies(ctx context.Context, filter bson.M) ([]*types.Agency, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var agencies []*types.Agency
	if err := resp.All(ctx, &agencies); err != nil {
		return nil, err
	}
	return agencies, err
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
