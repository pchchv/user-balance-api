package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/pchchv/env"
	"github.com/pchchv/golog"
)

type User struct {
	Id      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

var (
	testURL string
)

func init() {
	// Load values from .env into the system
	if err := env.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value
	// Outputs a panic if the value is missing
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

func main() {
	server()
}
