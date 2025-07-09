package models

type PaymentRequest struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
}
