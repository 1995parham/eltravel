package main

import (
	"fmt"
	"log"
	"p1/alg"
	"p1/calc"
	"p1/io"
	"p1/passenger"
	"strconv"
	"strings"
	"time"
)

func main() {
	var reader io.Reader
	var writer io.Writer

	reader = io.FileReader{
		Name: "roads.txt",
	}
	writer = io.StandardWriter{}

	g, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	var src string
	var dst string
	if _, err := fmt.Scanf("%s", &src); err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Scanf("%s", &dst); err != nil {
		log.Fatal(err)
	}

	var rawDep string
	if _, err := fmt.Scanf("%s", &rawDep); err != nil {
		log.Fatal(err)
	}
	dep, err := time.Parse(time.Kitchen, rawDep)
	if err != nil {
		log.Fatal(err)
	}

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

	roads := alg.Route(g, src, dst)

	result := calc.Calculate(dep, roads, passengers)

	if err := writer.Write(result); err != nil {
		log.Fatal(err)
	}
}
