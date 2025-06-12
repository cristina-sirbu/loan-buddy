package provider

type LoanOffer struct {
	Provider string  `json:"provider"`
	Rate     float64 `json:"rate"`
	Amount   int     `json:"amount"`
}

func GetMockOffers() []LoanOffer {
	return []LoanOffer{
		{Provider: "Provider A", Rate: 3.5, Amount: 10000},
		{Provider: "Provider B", Rate: 2.9, Amount: 12000},
		{Provider: "Provider C", Rate: 4.1, Amount: 9000},
	}
}
