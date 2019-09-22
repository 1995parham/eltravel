package result

import (
	"gitlab.com/1995parham/eltravel/route"
	"time"
)

// Path is a road that is used for traveling
type Path struct {
	Route         route.Route
	DepartureTime time.Time
	ArrivalTime   time.Time
	Duration      time.Duration
}
