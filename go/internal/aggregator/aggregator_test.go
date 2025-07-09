package aggregator

import "testing"

func TestAggregateOffers(t *testing.T) {
	offers := AggregateOffers()

	if len(offers) != 3 {
		t.Errorf("Expected 3 offers, got %d", len(offers))
	}

	for i := 0; i < len(offers)-1; i++ {
		if offers[i].Rate > offers[i+1].Rate {
			t.Errorf("Offers are not sorted by rate: %v < %v", offers[i].Rate, offers[i+1].Rate)
		}
	}
}
