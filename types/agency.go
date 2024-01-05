package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Agency struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Cars     []primitive.ObjectID `bson:"cars" json:"cars"`
	//maybe adding ratings here
}

type CarType int

const (
	_ CarType = iota
	EconomyCarsType
	CompactCarsType
	SUVCarsType
	SportCarsType
	MinivanCarsType
	HybridCarsType
	ElectricCarsType
	CargoVanType
	LuxuryCarsType
	FullSizeCarsType
)

type Car struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type      CarType            `bson:"type" json:"type"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"price" json:"price"`
	AgencyID  primitive.ObjectID `bson:"agencyID" json:"agencyID"`
}
