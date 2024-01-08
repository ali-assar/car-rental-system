package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	CarID    primitive.ObjectID `bson:"carID,omitempty" json:"carID,omitempty"`
	FromDate time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
}
