package io

import (
	"fmt"
	"p1/result"
	"p1/route"
	"time"
)

// Writer writes the final result
type Writer interface {
	Write([]result.Result) error
}

// StandardWriter writes the final result on standard output
type StandardWriter struct{}

// Write the final result on standard output with some details
func (StandardWriter) Write(results []result.Result) error {
	var totalDuration int
	var totalPrice int

	fmt.Println("result for you query is:")
	for _, result := range results {
		path := result.Path
		price := result.Price
		children, adults, infants := price.Passengers.CountByType()

		fmt.Println("**************************************************")
		fmt.Printf("type: %s", path.Route.Type())
		fmt.Printf("reference: %s\n", path.Route.ID())
		fmt.Printf("route: %s(%s) --> %s(%s)\n", path.Route.Src(), path.DepartureTime.Format(time.Kitchen), path.Route.Dst(), path.ArrivalTime.Format(time.Kitchen))
		fmt.Printf("duration: %d minutes\n", int(path.Duration.Minutes()))

		fmt.Printf("price: infants(%d) children(%d) adults(%d)\n", infants, children, adults)
		switch r := path.Route.(type) {
		case route.Road:
			fmt.Printf("\t%d car(s) required\n", r.NumberOfCars(price.Passengers))
		case route.Flight:
			fmt.Printf("\t%d * %d\n", adults, r.PricePerPerson)
			fmt.Printf("\t%d * %d\n", children, r.PricePerPerson/2)
			fmt.Printf("\t%d * %d\n", infants, r.PricePerPerson/10)
		}
		fmt.Printf("\ttotal:  %d\n", price.Total)

		totalDuration += int(path.Duration.Minutes())
		totalPrice += price.Total
	}
	fmt.Printf("\n************************\n")
	fmt.Printf("Total duration: %d minutes\n", totalDuration)
	fmt.Println("Total price: ", totalPrice)
	return nil
}
