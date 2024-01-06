package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Agency struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Cars     []primitive.ObjectID `bson:"cars" json:"cars"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type CarType int

const (
	_ CarType = iota
	EconomyCarsType
	CompactCarsType
	SUVCarsType
	SportCarsType
	MinivanCarsType
	CargoVanType
	LuxuryCarsType
	FullSizeCarsType
	MuscleCar
)

type Car struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type     string             `bson:"type" json:"type"`
	Fuel     string             `bson:"fuel" json:"fuel"` //petrol, hybrid,gas
	Year     int                `bson:"year" json:"year"`
	Model    string             `bson:"model" json:"model"`
	Price    float64            `bson:"price" json:"price"`
	AgencyID primitive.ObjectID `bson:"agencyID" json:"agencyID"`
}
