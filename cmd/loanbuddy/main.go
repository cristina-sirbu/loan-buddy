package main

import (
	"log"

	"github.com/cristina-sirbu/loan-buddy/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Set up routes
	e.GET("/offers", api.GetOffers)
	e.POST("/offers", api.CreateOffer)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
