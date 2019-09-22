package result

import (
	"p1/road"
	"time"
)

// Path is a road that is used for traveling
type Path struct {
	Road          road.Road
	DepartureTime time.Time
	ArrivalTime   time.Time
	Duration      time.Duration
}
