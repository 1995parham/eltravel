package io

import (
	"fmt"
	"p1/road"
	"time"
)

// Writer writes the final result
type Writer interface {
	Write([]road.Path) error
}

// StandardWriter writes the final result on standard output
type StandardWriter struct{}

// Write the final result on standard output with some details
func (StandardWriter) Write(roads []road.Path) error {
	var total int
	fmt.Println("result for you query is:")
	for _, road := range roads {
		fmt.Println("**************************************************")
		fmt.Println("type: Road")
		fmt.Printf("reference: %s\n", road.Road.Name)
		fmt.Printf("route: %s(%s) --> %s(%s)\n", road.Road.Source, road.DepartureTime.Format(time.Kitchen), road.Road.Destination, road.ArrivalTime.Format(time.Kitchen))
		fmt.Printf("duration: %d minutes\n", int(road.Duration.Minutes()))
		total += int(road.Duration.Minutes())
	}
	fmt.Printf("\n************************\n")
	fmt.Printf("Total duration: %d minutes\n", total)
	return nil
}
