package passenger

// Passengers is an array of passengers
type Passengers []Passenger

// Passenger represents a single passenger
type Passenger struct {
	Age int
}

// Type based on his/her age
type Type int

const (
	// Infant Passenger
	Infant Type = iota
	// Child Passenger
	Child Type = iota
	// Adult Passenger
	Adult Type = iota
)

// Type returns passenger type based on his/her age
func (p Passenger) Type() Type {
	if p.Age < 2 {
		return Infant
	} else if p.Age >= 2 && p.Age < 12 {
		return Child
	}
	return Adult
}

// CountByType returns a number of children, adults and infants
func (ps Passengers) CountByType() (int, int, int) {
	children := 0
	infants := 0
	adults := 0

	for _, p := range ps {
		switch p.Type() {
		case Child:
			children++
		case Adult:
			adults++
		case Infant:
			infants++
		}
	}

	return children, adults, infants
}

// Total returns the total number of passengers
func (ps Passengers) Total() int {
	return len(ps)
}
