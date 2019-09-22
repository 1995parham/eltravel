package road

import "time"

// Path is a road that is used for traveling
type Path struct {
	Road          Road
	DepartureTime time.Time
	ArrivalTime   time.Time
	Duration      time.Duration
}
