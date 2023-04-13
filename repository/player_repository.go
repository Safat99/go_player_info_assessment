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

type PlayerRepository interface {
	CreatePlayer(model.Player) (string, error)
	FindAll() ([]primitive.M, error)
	FindById(string) (*model.Player, error)
	FindByPlayerName(string) ([]bson.M, error)
	FindByCountry(string) ([]bson.M, error)

	UpdatePlayer(string, *model.UpdatePlayerDto) (*model.Player, error)
	DeletePlayerById(string) error
}

type playerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(client *mongo.Client, database string) *playerRepository {
	collection := client.Database(database).Collection("players")
	return &playerRepository{collection}
}

func (p *playerRepository) CreatePlayer(player model.Player) (string, error) {

	player.ID = primitive.NewObjectID()
	player.CreatedAt = time.Now()
	player.UpdatedAt = time.Now()

	_, err := p.collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Fatal(err)
	}

	return player.ID.Hex(), nil
}

func (p *playerRepository) FindById(playerId string) (*model.Player, error) {

	objId, err := primitive.ObjectIDFromHex(playerId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	player := &model.Player{}

	err = p.collection.FindOne(context.Background(), filter).Decode(player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (p *playerRepository) FindByPlayerName(playerName string) ([]bson.M, error) {

	filter := bson.M{"player_name": playerName}

	cursor, err := p.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *playerRepository) FindByCountry(country string) ([]bson.M, error) {

	filter := bson.M{"country": country}

	cursor, err := p.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *playerRepository) FindAll() ([]primitive.M, error) {

	filter := bson.D{{}}

	cursor, err := p.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []primitive.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *playerRepository) UpdatePlayer(playerId string, player *model.UpdatePlayerDto) (*model.Player, error) {

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
	err = p.collection.FindOneAndUpdate(context.Background(), filter, update, options).Decode(&updatedObject)
	if err != nil {
		return nil, err
	}
	return updatedObject, nil
}

func (p *playerRepository) DeletePlayerById(playerId string) error {
	objId, err := primitive.ObjectIDFromHex(playerId)
	if err != nil {
		return nil
	}

	filter := bson.M{"_id": objId}
	_, err = p.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
