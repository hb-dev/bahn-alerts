package schedule

import (
	"testing"
	"time"

	"github.com/hb-dev/bahn-alerts/bahn"
)

func TestDaysOfInterestMap(t *testing.T) {
	daysOfInterest := []string{"Monday", "Tuesday", "Friday"}
	daysOfInterestMap := daysOfInterestMap(daysOfInterest)

	if !(*daysOfInterestMap)["Monday"] {
		t.Fatalf("expect %v for Monday, got: %v", true, (*daysOfInterestMap)["Monday"])
	}

	if (*daysOfInterestMap)["Thursday"] {
		t.Fatalf("expect %v for Thursday, got: %v", false, (*daysOfInterestMap)["Thursday"])
	}
}

func TestTimeToDateString(t *testing.T) {
	ti, _ := time.Parse(time.RFC3339, "2018-08-03T15:04:05Z")
	timeToDateString := timeToDateString(ti)
	expected := "2018-08-03"

	if expected != timeToDateString {
		t.Fatalf("expected %s got %s", expected, timeToDateString)
	}
}

func TestDateOfInterest(t *testing.T) {
	daysOfInterestMap := map[string]bool{
		"Friday":   true,
		"Saturday": false,
	}

	ti, _ := time.Parse(time.RFC3339, "2018-08-03T15:04:05Z")

	dateOfInterest := dateOfInterest(ti, &daysOfInterestMap)
	if !dateOfInterest {
		t.Fatalf("expected %s to be of interest, got: %v", ti.Weekday(), dateOfInterest)
	}
}

func TestDateNotOfInterest(t *testing.T) {
	daysOfInterestMap := map[string]bool{
		"Friday":   false,
		"Saturday": true,
	}

	ti, _ := time.Parse(time.RFC3339, "2018-08-03T15:04:05Z")

	dateOfInterest := dateOfInterest(ti, &daysOfInterestMap)
	if dateOfInterest {
		t.Fatalf("expected %s to not be of interest, got: %v", ti.Weekday(), dateOfInterest)
	}
}

func TestDepartureTimeOfTrain(t *testing.T) {
	departures := bahn.DepartureCollection{
		bahn.Departure{
			TrainName: "ICE 321",
			DateTime:  "2018-08-03T15:04:05",
		},
		bahn.Departure{
			TrainName: "ICE 123",
			DateTime:  "2018-08-03T17:04:05",
		},
	}
	trainName := "ICE 123"

	expected := "17:04:05"

	departureTimeOfTrain := departureTimeOfTrain(&departures, trainName)

	if departureTimeOfTrain != expected {
		t.Fatalf("expected departure time to be %s, got %s", expected, departureTimeOfTrain)
	}
}

func TestDepartureTimeOfTrainNotFound(t *testing.T) {
	departures := bahn.DepartureCollection{
		bahn.Departure{
			TrainName: "ICE 321",
			DateTime:  "2018-08-03T15:04:05",
		},
		bahn.Departure{
			TrainName: "ICE 456",
			DateTime:  "2018-08-03T17:04:05",
		},
	}
	trainName := "ICE 123"

	expected := "No departure"

	departureTimeOfTrain := departureTimeOfTrain(&departures, trainName)

	if departureTimeOfTrain != expected {
		t.Fatalf("expected departure time to be %s, got %s", expected, departureTimeOfTrain)
	}
}
