package models

import "time"

type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	LoanID    string    `json:"loan_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
