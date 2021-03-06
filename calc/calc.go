package calc

import (
	"time"

	"gitlab.com/1995parham/eltravel/passenger"
	"gitlab.com/1995parham/eltravel/result"
	"gitlab.com/1995parham/eltravel/route"
)

// Calculate the price and path for the given routes
func Calculate(departure time.Time, routes []route.Route, passengers passenger.Passengers) []result.Result {
	results := make([]result.Result, len(routes))

	var currentTime = departure
	for i, r := range routes {
		results[i] = result.Result{
			Path: result.Path{
				Route:         r,
				DepartureTime: r.DepartureTime(currentTime),
				ArrivalTime:   r.ArrivalTime(currentTime),
				Duration:      r.Duration(),
			},
			Price: result.Price{
				Total:      r.Price(passengers),
				Passengers: passengers,
			},
		}
		currentTime = r.ArrivalTime(currentTime)
	}

	return results
}
