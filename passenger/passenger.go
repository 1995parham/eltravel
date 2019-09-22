package passenger

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
	}
	return Adult
}
