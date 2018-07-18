package schedule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hb-dev/bahn-alerts/bahn"
)

var TargetDate = timeToDateString(time.Now())

func Schedule(locationID int, trainName, departureTime string, daysOfInterest []string, limit int) (map[string]string, error) {
	t := stringToDateTime(TargetDate, departureTime)
	daysOfInterestMap := daysOfInterestMap(daysOfInterest)
	schedule := make(map[string]string, 0)

	count := 1
	for count <= limit {
		if dateOfInterest(t, daysOfInterestMap) {
			date := timeToDateString(t)
			departures, err := bahn.Departures(locationID, dateTimeStringWithTolerance(t))
			if err != nil {
				schedule[date] = "No departure (API Error)"
				count++
				t = t.AddDate(0, 0, 1)
				continue
			}
			if len(*departures) == 0 {
				schedule[date] = "No departure"
				count++
			}
			departureTimeOfTrain := departureTimeOfTrain(departures, trainName)
			if departureTimeOfTrain != "" {
				schedule[date] = departureTimeOfTrain
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
			return strings.Split(departure.DateTime, "T")[1]
		}
	}

	return fmt.Sprintf("No departure")
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

func timeToDateTimeString(t time.Time) string {
	return t.Local().Format("2006-01-02T15:04")
}

func stringToDateTime(d, t string) time.Time {
	dateTime, _ := time.Parse("2006-01-02T15:04", fmt.Sprintf("%sT%s", d, t))
	return dateTime
}

func dateTimeStringWithTolerance(t time.Time) string {
	return timeToDateTimeString(t.Add(time.Hour * -2))
}
