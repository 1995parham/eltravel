package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gitlab.com/1995parham/eltravel/alg"
	"gitlab.com/1995parham/eltravel/calc"
	"gitlab.com/1995parham/eltravel/io"
	"gitlab.com/1995parham/eltravel/passenger"
)

func main() {
	var reader io.Reader
	var writer io.Writer

	reader = io.FileReader{
		Name: "transportations.txt",
	}
	writer = io.StandardWriter{}

	g, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter the source:")
	var src string
	if _, err := fmt.Scanf("%s", &src); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter the destination:")
	var dst string
	if _, err := fmt.Scanf("%s", &dst); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter the departure time (eg. 08:00AM):")
	var rawDep string
	if _, err := fmt.Scanf("%s", &rawDep); err != nil {
		log.Fatal(err)
	}
	dep, err := time.Parse(time.Kitchen, rawDep)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter the passengers list that are separated by comma:")
	passengers := make([]passenger.Passenger, 0)
	var rawPassengers string
	if _, err := fmt.Scanf("%s", &rawPassengers); err != nil {
		log.Fatal(err)
	}
	for _, rp := range strings.Split(rawPassengers, ",") {
		age, err := strconv.Atoi(rp)
		if err != nil {
			log.Fatal(err)
		}
		passengers = append(passengers, passenger.Passenger{Age: age})
	}

	roads := alg.Route(g, src, dst, dep, passengers)

	result := calc.Calculate(dep, roads, passengers)

	if err := writer.Write(result); err != nil {
		log.Fatal(err)
	}
}
