package checker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hb-dev/bahn-alerts/bahn"
	"github.com/hb-dev/bahn-alerts/checker"
	"github.com/hb-dev/bahn-alerts/schedule"
)

func TestCheck(t *testing.T) {
	ts := exampleBahnAPIServer()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	daysOfInterest := []string{"Monday", "Tuesday", "Friday"}
	departureTime := "06:50"
	trainName := "ICE 123"
	limit := 3

	expectedChangedDepartureTimes := map[string]string{
		"2018-06-11": "06:51",
		"2018-06-12": "06:51",
		"2018-06-15": "06:51",
	}
	expectedChanged := true

	schedule.TargetDate = "2018-06-11"
	changed, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err != nil {
		t.Fatalf("checker.Check() failed: %s", err)
	}

	if expectedChanged != changed {
		t.Fatal("expected changed departures, but they didn't")
	}

	if expectedChangedDepartureTimes["2018-06-11"] != changedDepartureTimes["2018-06-11"] {
		t.Fatalf("expected changed departure to be %s, got %s", expectedChangedDepartureTimes["2018-06-11"], changedDepartureTimes["2018-06-11"])
	}
}

func TestCheckBahnAPIError(t *testing.T) {
	ts := exampleFailingBahnAPIServer()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	daysOfInterest := []string{"Monday"}
	departureTime := "06:50"
	trainName := "ICE 123"
	limit := 3

	expectedChangedDepartureTimes := map[string]string{"2018-06-11": "No departure (API Error)"}

	schedule.TargetDate = "2018-06-11"
	_, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err != nil {
		t.Fatalf("checker.Check() failed: %s", err)
	}

	if expectedChangedDepartureTimes["2018-06-11"] != changedDepartureTimes["2018-06-11"] {
		t.Fatalf("expected changed departure to be %s, got %s", expectedChangedDepartureTimes["2018-06-11"], changedDepartureTimes["2018-06-11"])
	}
}

func TestCheckNoDepartures(t *testing.T) {
	ts := exampleBahnAPIServerEmptyResonse()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	daysOfInterest := []string{"Monday"}
	departureTime := "06:50"
	trainName := "ICE 123"
	limit := 3

	expectedChangedDepartureTimes := map[string]string{"2018-06-11": "No departure"}
	expectedChanged := true

	schedule.TargetDate = "2018-06-11"
	changed, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err != nil {
		t.Fatalf("checker.Check() failed: %s", err)
	}

	fmt.Println("CHANGED:", changedDepartureTimes)

	if expectedChanged != changed {
		t.Fatal("expected changed departures, but they didn't")
	}

	if expectedChangedDepartureTimes["2018-06-11"] != changedDepartureTimes["2018-06-11"] {
		t.Fatalf("expected changed departure to be %s, got %s", expectedChangedDepartureTimes["2018-06-11"], changedDepartureTimes["2018-06-11"])
	}
}

func exampleBahnAPIServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, departuresAPIResponses[r.URL.String()])
	}))
}

func exampleBahnAPIServerEmptyResonse() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "[]")
	}))
}

func exampleFailingBahnAPIServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
}

var departuresAPIResponses = map[string]string{
	"/departureBoard/87654321?date=2018-06-11T06:50": `
    [
      {
        "name": "ICE 123",
        "stopId": 8011956,
        "stopName": "Jena Paradies",
        "dateTime": "2018-06-11T06:51"
      },
      {
        "name": "IC 2063",
        "stopId": 8011956,
        "stopName": "Jena Paradies",
        "dateTime": "2018-06-11T13:00"
      }
		]
  `,
	"/departureBoard/87654321?date=2018-06-12T06:50": `
		[
			{
				"name": "ICE 123",
				"stopId": 8011956,
				"stopName": "Jena Paradies",
				"dateTime": "2018-06-12T06:51"
			},
			{
				"name": "IC 2063",
				"stopId": 8011956,
				"stopName": "Jena Paradies",
				"dateTime": "2018-06-12T13:00"
			}
		]
	`,
	"/departureBoard/87654321?date=2018-06-15T06:50": `
		[
			{
				"name": "ICE 123",
				"stopId": 8011956,
				"stopName": "Jena Paradies",
				"dateTime": "2018-06-15T06:51"
			},
			{
				"name": "IC 2063",
				"stopId": 8011956,
				"stopName": "Jena Paradies",
				"dateTime": "2018-06-15T13:00"
			}
		]
	`,
}
