package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEcho() *echo.Echo {
	e := echo.New()

	// Set up routes
	e.GET("/offers", GetOffers)
	e.POST("/offers", CreateOffer)

	return e
}

func TestGetOffers(t *testing.T) {
	e := setupEcho()

	req := httptest.NewRequest(http.MethodGet, "/offers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetOffers(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var offers []map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &offers)
	assert.NoError(t, err)
	assert.True(t, len(offers) >= 1)
}

func TestCreateOffer(t *testing.T) {
	e := setupEcho()

	offer := `{"provider": "TestProvider", "rate": 3.9, "amount": 15000}`
	req := httptest.NewRequest(http.MethodPost, "/offers", bytes.NewBufferString(offer))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CreateOffer(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var created map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, "TestProvider", created["provider"])
	assert.Equal(t, 3.9, created["rate"])
	assert.Equal(t, float64(15000), created["amount"])
}
