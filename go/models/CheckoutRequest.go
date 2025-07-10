package models

type CheckoutRequest struct {
	UserID string `json:"user_id"`
	LoanID string `json:"loan_id"`
}