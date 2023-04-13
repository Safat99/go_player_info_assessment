package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pos string

const (
	Goalkeeper Pos = "Goalkeeper"
	Defender   Pos = "Defender"
	Midfielder Pos = "Midfielder"
	Forward    Pos = "Forward"
	Manager    Pos = "Manager"
)

type Player struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlayerName string             `json:"player_name,omitempty" bson:"player_name"`
	Age        int64              `json:"age,omitempty" bson:"age"`
	Position   Pos                `json:"position,omitempty" bson:"position"`
	Country    string             `json:"country,omitempty" bson:"country"`
	DOB        time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type UpdatePlayerDto struct {
	PlayerName string `json:"player_name,omitempty" bson:"player_name,omitempty"`
	Position   Pos    `json:"position,omitempty" bson:"position,omitempty"`
}
