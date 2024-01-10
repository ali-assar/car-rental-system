package db

import (
	"context"

	"github.com/Ali-Assar/car-rental-system/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AgencyStore interface {
	InsertAgency(context.Context, *types.Agency) (*types.Agency, error)
	UpdateAgency(context.Context, Map, Map) error
	GetAgencies(context.Context, Map, *Pagination) ([]*types.Agency, error)
	GetAgencyByID(context.Context, string) (*types.Agency, error)
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

func (s *MongoAgencyStore) GetAgencyByID(ctx context.Context, id string) (*types.Agency, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var agency types.Agency
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&agency); err != nil {
		return nil, err
	}

	return &agency, nil
}

func (s *MongoAgencyStore) GetAgencies(ctx context.Context, filter Map, pag *Pagination) ([]*types.Agency, error) {
	opts := options.FindOptions{}
	opts.SetSkip((pag.Page - 1) * pag.Limit)
	opts.SetLimit(pag.Limit)
	resp, err := s.coll.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}
	var agencies []*types.Agency
	if err := resp.All(ctx, &agencies); err != nil {
		return nil, err
	}
	return agencies, err
}

func (s *MongoAgencyStore) UpdateAgency(ctx context.Context, filter, update Map) error {
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
