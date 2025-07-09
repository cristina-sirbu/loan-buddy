package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cristina-sirbu/loan-buddy/internal/aggregator"
	"github.com/cristina-sirbu/loan-buddy/internal/provider"
	"github.com/cristina-sirbu/loan-buddy/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var Orders []models.Order

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

// POST /checkout
func Checkout(c echo.Context) error {
	type CheckoutRequest struct {
		UserID string `json:"user_id"`
		LoanID string `json:"loan_id"`
	}

	var req CheckoutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
	}
	if !isValidLoanID(req.LoanID) {
		log.Printf("Invalid loan_id requested: loan_id=%s\n", req.LoanID)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid loan_id"})
	}

	newOrder := models.Order{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		LoanID:    req.LoanID,
		Status:    "PENDING",
		CreatedAt: time.Now(),
	}

	newOrder.Status = simulateOrderStatus(newOrder)

	Orders = append(Orders, newOrder)

	writeOrderToFile(newOrder)

	log.Printf("New order placed: user_id=%s loan_id=%s status=%s\n", newOrder.UserID, newOrder.LoanID, newOrder.Status)
	return c.JSON(http.StatusCreated, newOrder)
}

func simulateOrderStatus(order models.Order) string {
	// statuses := []string{"APPROVED", "REJECTED"}

	payment := models.PaymentRequest{
		OrderID: order.ID,
		Amount:  getAmountFromLoanID(order.LoanID),
		Status:  order.Status,
	}

	jsonData, err := json.Marshal(payment)
	if err != nil {
		log.Printf("Error marshalling payment request: %v\n", err)
		return "REJECTED"
	}

	paymentServiceURL := os.Getenv("PAYMENT_SERVICE_URL")
	if paymentServiceURL == "" {
		log.Println("PAYMENT_SERVICE_URL not set, using default")
		paymentServiceURL = "http://localhost:8000/confirm-payment"
	}

	response, err := http.Post(paymentServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending payment request: %v\n", err)
		return "REJECTED"
	}
	defer response.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Printf("Error decoding payment response: %v\n", err)
		return "FAILED"
	}

	status, ok := result["status"]
	if !ok {
		log.Println("Payment response does not contain status")
		return "FAILED"
	}

	log.Printf("Payment response received: order_id=%s status=%s\n", order.ID, status)
	return status
}

func writeOrderToFile(order models.Order) {
	file, err := os.OpenFile("orders.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening orders file: %v\n", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(order); err != nil {
		log.Printf("Error writing order to file: %v\n", err)
		return
	}
	log.Printf("Order written to file: %s\n", order.ID)
}

func getAmountFromLoanID(loanID string) int {
	offers := aggregator.AggregateOffers()
	for _, offer := range offers {
		if offer.ID == loanID {
			return offer.Amount
		}
	}
	return 0
}

func isValidLoanID(id string) bool {
	offers := aggregator.AggregateOffers()
	for _, offer := range offers {
		if offer.ID == id {
			return true
		}
	}
	return false
}
