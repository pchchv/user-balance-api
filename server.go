package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pchchv/golog"
)

// handlePing checks that the server is up and running
func handlePing(c echo.Context) error {
	message := "User balance API. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

// handleDeposit deposits funds to the user's balance
func handleDeposit(c echo.Context) error {
	var request struct {
		UserID uuid.UUID `json:"user_id"`
		Amount float64   `json:"amount"`
	}

	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	user, err := deposit(request.UserID, request.Amount)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// The declaration of all routes comes from it.
func routes(e *echo.Echo) {
	e.GET("/", handlePing)
	e.GET("/ping", handlePing)
	e.PATCH("/users/deposit", handleDeposit)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
