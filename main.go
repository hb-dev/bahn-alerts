package main

import (
	"fmt"
	"github.com/hb-dev/bahn-alerts/checker"
)

func main() {
	locationID := 8011956 // Jena Paradies
	time := "06:51"
	trainName := "IsCE 1526"
	daysOfInterest := []string{"Monday", "Tuesday", "Wednesday", "Thursday"}

	changed, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, time, trainName, 10)
	if err != nil {
		errMsg := err.Error()
		if errMsg == fmt.Sprintf("Train %s not found", trainName) {
			fmt.Println(err.Error())
			changedDepartureTimes = []string{errMsg}
		} else {
			panic(err)
		}
	}

	if changed {
		fmt.Println("Schedule changed:", changedDepartureTimes)
	}

	// define alert parameters
	// alert changed schedule
}
