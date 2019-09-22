package route

import (
	"p1/passenger"
	"strings"
	"time"
)

// KitchenTime is a time in kitchen format
type KitchenTime time.Time

// UnmarshalJSON unmarshals json data to kitchen format
func (kt *KitchenTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.Kitchen, s)
	if err != nil {
		return err
	}
	*kt = KitchenTime(t)
	return nil
}

// Flight represents a flight path
type Flight struct {
	Name           string      `json:"flight_name"`
	Source         string      `json:"source"`
	Destination    string      `json:"destination"`
	TakeoffTime    KitchenTime `json:"take_off_time"`
	LandingTime    KitchenTime `json:"landing_time"`
	PricePerPerson int         `json:"price_per_person"`
	Capacity       int         `json:"capacity"`
}

// IsAvailable check the flight constraints for given departure time and passengers
func (f Flight) IsAvailable(dep time.Time, ps passenger.Passengers) bool {
	if ps.Total() > f.Capacity {
		return false
	}
	if dep.Add(30 * time.Minute).After(time.Time(f.TakeoffTime)) {
		return false
	}

	return true
}

// ArrivalTime calculates the arrival time based on given departure time
func (f Flight) ArrivalTime(dep time.Time) time.Time {
	return time.Time(f.LandingTime)
}

// Price calculates the flight price for given passengers
func (f Flight) Price(ps passenger.Passengers) int {
	children, adults, infants := ps.CountByType()

	return adults*f.PricePerPerson + children*(f.PricePerPerson/2) + infants*(f.PricePerPerson/10)
}

// Type return route Type
func (Flight) Type() string {
	return "Flight"
}

// Src return source of route
func (f Flight) Src() string {
	return f.Source
}

// Dst return destination of route
func (f Flight) Dst() string {
	return f.Destination
}

// DepartureTime returns the actual departure time
func (f Flight) DepartureTime(time.Time) time.Time {
	return time.Time(f.TakeoffTime)
}

// Duration returns the trip duration
func (f Flight) Duration() time.Duration {
	return time.Time(f.LandingTime).Sub(time.Time(f.TakeoffTime))
}

// ID returns the route ID
func (f Flight) ID() string {
	return f.Name
}
