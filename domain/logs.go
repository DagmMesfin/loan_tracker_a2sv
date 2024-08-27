package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Activity  string             `json:"activity" bson:"activity"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
