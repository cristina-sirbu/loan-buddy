package aggregator

import (
	"sort"

	"github.com/cristina-sirbu/loan-buddy/internal/provider"
)

func AggregateOffers() []provider.LoanOffer {
	offers := provider.GetMockOffers()

	sort.Slice(offers, func(i, j int) bool {
		return offers[i].Rate < offers[j].Rate
	})

	return offers
}
