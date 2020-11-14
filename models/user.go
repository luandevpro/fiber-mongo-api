package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email"`
	Name      string             `json:"name"`
	Age       int                `json:"age,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	Status    bool               `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
}
