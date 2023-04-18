package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminUser struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
