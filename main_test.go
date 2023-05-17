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
		t.Fatal(errors.New("Error when receiving the user's balance, wrong balanceю"))
	}
}

func TestDeposit(t *testing.T) {
	var testAmount = 9.99

	user, err := getBalance(testID)
	if err != nil {
		t.Fatal(err)
	}

	resUser, err := deposit(testID, testAmount)
	if err != nil {
		t.Fatal(err)
	}

	if user.Balance+testAmount != resUser.Balance {
		t.Errorf("Error when topping up the balance. Expected: %v, Exist: %v", user.Balance+testAmount, resUser.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	var testAmount = 4.49

	user, err := getBalance(testID)
	if err != nil {
		t.Fatal(err)
	}

	resUser, err := withdraw(testID, testAmount)
	if err != nil {
		t.Fatal(err)
	}

	if user.Balance-testAmount != resUser.Balance {
		t.Errorf("Error when withdrawing funds from the balance. Expected: %v, Existing: %v", user.Balance+testAmount, resUser.Balance)
	}
}

func TestDeleteUser(t *testing.T) {
	user, err := deleteUser(testID)
	if err != nil {
		t.Fatal(err)
	}

	if user.Balance != 0 {
		t.Fatal(errors.New("Error when deleting a user. Incorrect balance was returned."))
	}
}
