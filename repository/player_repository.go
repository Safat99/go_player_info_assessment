package repository

import (
	"context"
	"log"
	"player_info/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRepository struct {
	Collection *mongo.Collection
}

func NewPlayerRepository(db *mongo.Database) *PlayerRepository {
	return &PlayerRepository{
		Collection: db.Collection("players"),
	}
}

func (p *PlayerRepository) CreatePlayer(player *model.Player) (string, error) {

	player.ID = primitive.NewObjectID()
	player.CreatedAt = time.Now()
	player.UpdatedAt = time.Now()

	_, err := p.Collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Fatal(err)
	}

	return player.ID.Hex(), nil
}

func (p *PlayerRepository) FindById(playerId string) (*model.Player, error) {

	objId, err := primitive.ObjectIDFromHex(playerId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	player := &model.Player{}

	err = p.Collection.FindOne(context.Background(), filter).Decode(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (p *PlayerRepository) UpdatePlayer(playerId string, player *model.UpdatePlayerDto) (*model.Player, error) {

	objId, err := primitive.ObjectIDFromHex(playerId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{
		"$set": bson.M{
			"player_name": player.PlayerName,
			"position":    player.Position,
			"updated_at":  time.Now(),
		},
	}

	var updatedObject *model.Player
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = p.Collection.FindOneAndUpdate(context.Background(), filter, update, options).Decode(&updatedObject)
	if err != nil {
		return nil, err
	}
	return updatedObject, nil
}
