package main

import (
	"fmt"
	"time"

	"github.com/hb-dev/bahn-alerts/checker"
)

type setting struct {
	LocationID     int
	TrainName      string
	DepartureTime  string
	DaysOfInterest []string
	Limit          int
}

type settings []setting

func main() {

	sets := settings{
		setting{
			LocationID:     8011956, // Jena Paradies
			TrainName:      "ICE 1526",
			DepartureTime:  "06:52",
			DaysOfInterest: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"},
			Limit:          4, // limit > 9 -> 500 Response from bahn API
		},
		setting{
			LocationID:     8010205, // Leipzig Hbf
			TrainName:      "IC 2060",
			DepartureTime:  "15:46",
			DaysOfInterest: []string{"Monday", "Tuesday", "Wednesday", "Thursday"},
			Limit:          4, // limit > 9 -> 500 Response from bahn API
		},
	}

	for {
		for _, set := range sets {
			_, times, err := checker.Check(set.LocationID, set.DaysOfInterest, set.DepartureTime, set.TrainName, set.Limit)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Changed departure Times for %s: %v\n", set.TrainName, times)
		}
		time.Sleep(1 * time.Minute)
	}

	// define alert parameters
	// alert changed schedule
}
