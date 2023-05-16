package main

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func Test(t *testing.T) {
	testURL = "http://" + getEnvValue("HOST") + ":" + getEnvValue("PORT")
	testID = uuid.New()
}

func TestCreateUser(t *testing.T) {
	user, err := createUser(testID, 0)
	if err != nil {
		t.Fatal(err)
	}

	if user.Balance != 0 {
		t.Fatal(errors.New("Error when creating a user, incorrect balance."))
	}
}

func TestGetBalance(t *testing.T) {
	user, err := getBalance(testID)
	if err != nil {
		t.Fatal(err)
	}

	if user.Balance != 0 {
		t.Fatal(errors.New("Error when creating a user, incorrect balance."))
	}
}
