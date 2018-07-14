package main

import (
	"fmt"
	"github.com/hb-dev/bahn-alerts/bahn"
)

func main() {

	// get data from bahn api
	locationID := 8011956 // Jena Paradies

	departures, err := bahn.Departures(locationID, "2018-06-12")
	if err != nil {
		panic(err)
	}

	for _, departure := range *departures {
		fmt.Printf("%s - %s\n", departure.TrainName, departure.DateTime)
	}

	// filter data
	// define schedule parameters
	// check for changed schedule
	// define alert parameters
	// alert changed schedule
}
