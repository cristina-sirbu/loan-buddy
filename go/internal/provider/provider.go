package provider

type LoanOffer struct {
	ID       string  `json:"id"`
	Provider string  `json:"provider"`
	Rate     float64 `json:"rate"`
	Amount   int     `json:"amount"`
}

func GetMockOffers() []LoanOffer {
	return []LoanOffer{
		{ID: "offer1", Provider: "Provider A", Rate: 3.5, Amount: 10000},
		{ID: "offer2", Provider: "Provider B", Rate: 2.9, Amount: 12000},
		{ID: "offer3", Provider: "Provider C", Rate: 4.1, Amount: 9000},
	}
}
