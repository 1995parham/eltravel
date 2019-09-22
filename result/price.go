package result

import "gitlab.com/1995parham/eltravel/passenger"

// Price represents the travel price
type Price struct {
	Total      int
	Passengers passenger.Passengers
}
