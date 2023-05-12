package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pchchv/golog"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "User balance API. Version 0.0.1"
	return c.String(http.StatusOK, message)
}

// The declaration of all routes comes from it.
func routes(e *echo.Echo) {
	e.GET("/ping", pingHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	golog.Fatal(e.Start(":" + getEnvValue("PORT")).Error())
}
