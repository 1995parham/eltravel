package calc

import (
	"p1/passenger"
	"p1/result"
	"p1/road"
	"time"
)

// Calculate the price and path for the given roads
func Calculate(departure time.Time, roads []road.Road, passengers []passenger.Passenger) []result.Result {
	children := 0
	infants := 0
	adults := 0

	for _, p := range passengers {
		switch p.Type() {
		case passenger.Child:
			children++
		case passenger.Adult:
			adults++
		case passenger.Infant:
			infants++
		}
	}

	cars := (children + adults) / 4
	if (children+adults)%4 != 0 {
		cars++
	}

	results := make([]result.Result, len(roads))
	var total time.Duration
	for i, rd := range roads {
		dur := time.Duration(rd.Length*60/rd.SpeedLimit) * time.Minute

		var price int
		if rd.Length > 50 {
			price = 400_000 + (rd.Length-50)*6_000
		} else {
			price = 400_000
		}

		results[i] = result.Result{
			Path: result.Path{
				Road:          rd,
				DepartureTime: departure.Add(total),
				ArrivalTime:   departure.Add(total).Add(dur),
				Duration:      dur,
			},
			Price: result.Price{
				Total:    price * cars,
				Cars:     cars,
				Children: children,
				Adults:   adults,
				Infants:  infants,
			},
		}

		total += dur
	}

	return results
}
