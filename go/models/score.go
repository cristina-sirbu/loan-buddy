package models

type ScoreResponse struct {
	LoanID              string  `json:"loan_id"`
	RiskScore           float64 `json:"risk_score"`
	ApprovalProbability float64 `json:"approval_probability"`
}