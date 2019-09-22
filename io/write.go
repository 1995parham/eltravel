package io

import (
	"fmt"
	"p1/result"
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

		fmt.Println("**************************************************")
		fmt.Println("type: Road")
		fmt.Printf("reference: %s\n", path.Road.Name)
		fmt.Printf("route: %s(%s) --> %s(%s)\n", path.Road.Source, path.DepartureTime.Format(time.Kitchen), path.Road.Destination, path.ArrivalTime.Format(time.Kitchen))
		fmt.Printf("duration: %d minutes\n", int(path.Duration.Minutes()))

		fmt.Printf("price: infants(%d) children(%d) adults(%d)\n", price.Infants, price.Children, price.Adults)
		fmt.Printf("\t%d car(s) required\n", price.Cars)
		fmt.Printf("\ttotal:  %d\n", price.Total)

		totalDuration += int(path.Duration.Minutes())
		totalPrice += price.Total
	}
	fmt.Printf("\n************************\n")
	fmt.Printf("Total duration: %d minutes\n", totalDuration)
	fmt.Println("Total price: ", totalPrice)
	return nil
}
