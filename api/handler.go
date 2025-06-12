package api

import (
	"net/http"

	"github.com/cristina-sirbu/loan-buddy/internal/aggregator"
	"github.com/cristina-sirbu/loan-buddy/internal/provider"
	"github.com/labstack/echo/v4"
)

// GET /offers
func GetOffers(c echo.Context) error {
	offers := aggregator.AggregateOffers()
	if len(offers) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No offers available"})
	}

	return c.JSON(http.StatusOK, offers)
}

// POST /offers
func CreateOffer(c echo.Context) error {
	var offer provider.LoanOffer
	if err := c.Bind(&offer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
	}

	// For now, just echo it back
	return c.JSON(http.StatusCreated, offer)
}
