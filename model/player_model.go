package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pos int

const (
	Goalkeeper Pos = iota
	Defender
	Midfielder
	Forward
	Manager
)

type Player struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlayerName string             `json:"player_name,omitempty"`
	Age        int64              `json:"age,omitempty"`
	Position   Pos                `json:"position,omitempty"`
	Country    string             `json:"country,omitempty"`
	DOB        time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
