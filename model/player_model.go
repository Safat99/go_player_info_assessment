package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Pos string

// const (
// 	Goalkeeper Pos = "Goalkeeper"
// 	Defender   Pos = "Defender"
// 	Midfielder Pos = "Midfielder"
// 	Forward    Pos = "Forward"
// 	Manager    Pos = "Manager"
// )

type Player struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlayerName string             `json:"player_name" bson:"player_name" validate:"required"`
	Age        int64              `json:"age" bson:"age" validate:"required,gte=0,lte=150"`
	Position   string             `json:"position" bson:"position" validate:"required,oneof=Goalkeeper Defender Midfielder Forward Manager"`
	Country    string             `json:"country" bson:"country" validate:"required"`
	DOB        string             `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type UpdatePlayerDto struct {
	PlayerName string `json:"player_name,omitempty" bson:"player_name,omitempty"`
	Position   string `json:"position,omitempty" bson:"position,omitempty" validate:"required,oneof=Goalkeeper Defender Midfielder Forward Manager"`
}

// func validateDateFormat(f1 validator.FieldLevel)
