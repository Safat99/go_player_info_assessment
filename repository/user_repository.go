package repository

import (
	"context"
	"log"
	"player_info/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(model.User) (*model.User, error)
	FindByUserName(string) ([]bson.M, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, database string) *userRepository {
	collection := client.Database(database).Collection("users")
	return &userRepository{collection}
}

func (u *userRepository) CreateUser(user model.User) (string, error) {

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Fatal(err)
	}

	return user.ID.Hex(), nil
}
