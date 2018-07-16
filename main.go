package main

import (
	"fmt"
	"github.com/hb-dev/bahn-alerts/checker"
)

func main() {
	locationID := 8011956 // Jena Paradies
	time := "06:51"
	trainName := "ICE 1526"
	daysOfInterest := []string{"Monday", "Tuesday", "Wednesday", "Thursday"}

	// limit > 9 -> 500 Response from bahn API
	_, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, time, trainName, 9)
	if err != nil {
		panic(err)
	}

	fmt.Println("Changed departure Times:", changedDepartureTimes)

	// define alert parameters
	// alert changed schedule
}
