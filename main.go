package main

import (
	"fmt"
	"log"
	"p1/alg"
	"p1/io"
	"p1/road"
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

	roads := alg.Route(g, src, dst)

	path := make([]road.Path, len(roads))
	var total time.Duration
	for i, rd := range roads {
		dur := time.Duration(rd.Length*60/rd.SpeedLimit) * time.Minute

		path[i] = road.Path{
			Road:          rd,
			DepartureTime: dep.Add(total),
			ArrivalTime:   dep.Add(total).Add(dur),
			Duration:      dur,
		}

		total += dur
	}

	if err := writer.Write(path); err != nil {
		log.Fatal(err)
	}
}
