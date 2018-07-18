package main

import (
	"fmt"
	"github.com/hb-dev/bahn-alerts/checker"
)

func main() {
	locationID := 8011956 // Jena Paradies
	// locationID := 8010205 // Leipzig Hbf
	time := "06:50"
	// time := "15:46"
	trainName := "ICE 1526"
	// trainName := "IC 2060"
	daysOfInterest := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// limit > 9 -> 500 Response from bahn API
	_, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, time, trainName, 9)
	if err != nil {
		panic(err)
	}

	fmt.Println("Changed departure Times:", changedDepartureTimes)

	// define alert parameters
	// alert changed schedule
}
