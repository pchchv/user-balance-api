package main

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pchchv/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func depositToDB(user User) error {
	filter := bson.D{{Key: "_id", Value: user.Id}}
	update := bson.D{{Key: "balance", Value: user.Balance}}

	_, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func getUserFromDB(id uuid.UUID) (user User, err error) {
	res := usersCollection.FindOne(context.TODO(), bson.M{"id": id})
	err = res.Decode(user)
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, err
}

func db() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(getEnvValue("MONGO")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		golog.Fatal(err.Error())
	}

	golog.Info("Connected to MongoDB!")

	usersCollection = client.Database(getEnvValue("DATABASE")).Collection("users")
}
