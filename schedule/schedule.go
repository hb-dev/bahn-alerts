package schedule

import (
	"fmt"
	"time"

	"github.com/hb-dev/bahn-alerts/bahn"
)

var TargetTime = time.Now()

func Schedule(locationID int, trainName string, daysOfInterest []string, limit int) ([]string, error) {
	t := TargetTime
	daysOfInterestMap := daysOfInterestMap(daysOfInterest)
	schedule := make([]string, 0)

	count := 1
	for count <= limit {
		if dateOfInterest(t, daysOfInterestMap) {
			date := timeToDateString(t)
			departures, err := bahn.Departures(locationID, date)
			if err != nil {
				return schedule, err
			}
			if len(*departures) == 0 {
				schedule = append(schedule, fmt.Sprintf("No departure on %s", date))
				count++
			}
			departureTimeOfTrain := departureTimeOfTrain(departures, trainName)
			if departureTimeOfTrain != "" {
				if departureTimeOfTrain == "No departure found" {
					departureTimeOfTrain = fmt.Sprintf("No departure on %s", date)
				}
				schedule = append(schedule, departureTimeOfTrain)
				count++
			}
		}
		t = t.AddDate(0, 0, 1)
	}

	return schedule, nil
}

func departureTimeOfTrain(departures *bahn.DepartureCollection, trainName string) string {
	for _, departure := range *departures {
		if departure.TrainName == trainName {
			return departure.DateTime
		}
	}

	return fmt.Sprintf("No departure found")
}

func dateOfInterest(t time.Time, daysOfInterestMap *map[string]bool) bool {
	return (*daysOfInterestMap)[t.Weekday().String()]
}

func daysOfInterestMap(daysOfInterest []string) *map[string]bool {
	daysOfInterestMap := map[string]bool{
		"Monday":    false,
		"Tuesday":   false,
		"Wednesday": false,
		"Thursday":  false,
		"Friday":    false,
		"Saturday":  false,
		"Sunday":    false,
	}

	for _, day := range daysOfInterest {
		daysOfInterestMap[day] = true
	}

	return &daysOfInterestMap
}

func timeToDateString(t time.Time) string {
	return t.Local().Format("2006-01-02")
}
