package main

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/pchchv/env"
	"github.com/pchchv/golog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

var (
	testURL         string
	usersCollection *mongo.Collection
)

// init loadы values from .env into the system.
func init() {
	if err := env.Load(); err != nil {
		golog.Panic("No .env file found")
	}
}

// getEnvValue retrieves values from the environment (.env) file.
// Outputs a panic message if the value is missing.
func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		golog.Panic("Value %v does not exist", v)
	}

	return value
}

func deposit(id uuid.UUID, amount float64) (User, error) {
	u := User{}
	// TODO: Retrieve data from the database.
	// TODO: Update balance
	return u, nil
}

func withdraw(id uuid.UUID, amount float64) (User, error) {
	u := User{}
	// TODO: Retrieve data from the database.
	// TODO: Update balance
	return u, nil
}

func getBalance(id uuid.UUID) (user User, err error) {
	res := usersCollection.FindOne(context.TODO(), bson.M{"id": id})
	err = res.Decode(user)
	if err != nil {
		return user, errors.New("User not found")
	}

	return
}

func main() {
	server()
	db()
}
