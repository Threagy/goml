package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUser() *mongo.Collection {
	return mongoClient.Database(dbName).Collection("user")
}

// GetUser ...
func GetUser(userID string, password string) (*UserModel, error) {
	filter := bson.D{primitive.E{Key: "user_id", Value: userID}}
	user := &UserModel{}
	err := getUser().FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("cannot find from db: %v", err)
	}

	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}

	return user, nil
}

// GetUsers ...
func GetUsers() ([]*UserModel, error) {
	var users []*UserModel
	ctx := context.Background()
	cursor, err := getUser().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for {
		eof := cursor.Next(ctx)
		if !eof {
			break
		}

		user := &UserModel{}
		cursor.Decode(user)
		users = append(users, user)
	}

	return users, nil
}

// NewUser ...
func NewUser(user UserModel) error {
	_, err := getUser().InsertOne(context.Background(), user)

	if err != nil {
		return fmt.Errorf("cannot create the user: %v", err)
	}

	return nil
}
