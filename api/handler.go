package api

import (
	"net/http"

	"github.com/cristina-sirbu/loan-buddy/internal/aggregator"
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
