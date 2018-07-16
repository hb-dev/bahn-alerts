package checker

import (
	"strings"

	"github.com/hb-dev/bahn-alerts/schedule"
)

func Check(locationID int, daysOfInterest []string, departureTime, trainName string, limit int) (bool, []string, error) {
	changed := false
	changedDepartureTimes := make([]string, 0)

	schedule, err := schedule.Schedule(locationID, trainName, daysOfInterest, limit)
	if err != nil {
		return changed, changedDepartureTimes, err
	}

	if len(schedule) < 1 {
		changed = true
		return changed, []string{"No Departures found"}, nil
	}

	for _, s := range schedule {
		time := s
		if !strings.HasPrefix(s, "No departure on") {
			time = strings.Split(s, "T")[1]
		}
		if departureTime != time {
			changedDepartureTimes = append(changedDepartureTimes, s)
		}
	}

	if len(changedDepartureTimes) > 0 {
		changed = true
	}

	return changed, changedDepartureTimes, nil
}
