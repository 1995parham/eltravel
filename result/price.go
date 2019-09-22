package result

import "p1/passenger"

// Price represents the travel price
type Price struct {
	Total      int
	Passengers passenger.Passengers
}
