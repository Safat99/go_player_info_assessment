package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
