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
	FindByUserNameAndRole(username string, role string) ([]bson.M, error)
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

	_, err := u.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	return user.ID.Hex(), nil
}

func (u *userRepository) FindByUserNameAndRole(username string, role string) ([]bson.M, error) {

	filter := bson.M{"name": username, "role": role}

	cursor, err := u.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
