package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/pchchv/env"
	"github.com/pchchv/golog"
)

type User struct {
	Id      uuid.UUID
	Name    string
	Balance float64
}

var testURL string

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

func main() {
	server()
}
