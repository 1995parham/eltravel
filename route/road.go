package route

import (
	"gitlab.com/1995parham/eltravel/passenger"
	"time"
)

// Road represents a route in country with its details
type Road struct {
	Name        string `json:"road_name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Length      int    `json:"length"`
	SpeedLimit  int    `json:"speed_limit"`
}

// IsAvailable check the road constraints for given departure time and passengers
func (Road) IsAvailable(time.Time, passenger.Passengers) bool {
	return true
}

// ArrivalTime calculates the arrival time based on given departure time
func (r Road) ArrivalTime(dep time.Time) time.Time {
	dur := r.Duration()
	return dep.Add(dur)
}

// NumberOfCars return number of cars for the road
func (r Road) NumberOfCars(ps passenger.Passengers) int {
	children, adults, _ := ps.CountByType()

	cars := (children + adults) / 4
	if (children+adults)%4 != 0 {
		cars++
	}

	return cars
}

// Price calculates the road price for given passengers
func (r Road) Price(ps passenger.Passengers) int {
	cars := r.NumberOfCars(ps)

	var price int
	if r.Length > 50 {
		price = 400_000 + (r.Length-50)*6_000
	} else {
		price = 400_000
	}

	return price * cars
}

// Type return route Type
func (Road) Type() string {
	return "Road"
}

// Src return source of route
func (r Road) Src() string {
	return r.Source
}

// Dst return destination of route
func (r Road) Dst() string {
	return r.Destination
}

// DepartureTime returns the actual departure time
func (Road) DepartureTime(dep time.Time) time.Time {
	return dep
}

// Duration returns the trip duration
func (r Road) Duration() time.Duration {
	return time.Duration(r.Length*60/r.SpeedLimit) * time.Minute
}

// ID returns the route ID
func (r Road) ID() string {
	return r.Name
}
