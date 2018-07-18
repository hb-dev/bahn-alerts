package checker

import (
	"github.com/hb-dev/bahn-alerts/schedule"
)

func Check(locationID int, daysOfInterest []string, departureTime, trainName string, limit int) (bool, map[string]string, error) {
	changed := false
	changedDepartureTimes := make(map[string]string, 0)

	schedule, err := schedule.Schedule(locationID, trainName, departureTime, daysOfInterest, limit)
	if err != nil {
		return changed, changedDepartureTimes, err
	}

	if len(schedule) < 1 {
		changed = true
		return changed, map[string]string{"0": "No departures found"}, nil
	}

	for date, time := range schedule {
		if departureTime != time {
			changedDepartureTimes[date] = time
		}
	}

	if len(changedDepartureTimes) > 0 {
		changed = true
	}

	return changed, changedDepartureTimes, nil
}
