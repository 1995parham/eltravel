package route

import (
	"gitlab.com/1995parham/eltravel/passenger"
	"time"
)

// Route is an abstraction to any path between cities
type Route interface {
	IsAvailable(time.Time, passenger.Passengers) bool
	ArrivalTime(time.Time) time.Time
	DepartureTime(time.Time) time.Time
	Price(passenger.Passengers) int
	Duration() time.Duration
	Type() string
	Dst() string
	Src() string
	ID() string
}
