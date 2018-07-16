package checker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	expectedChangedDepartureTimes := []string{
		"2018-06-11T06:51",
		"2018-06-12T06:51",
		"2018-06-15T06:51",
	}
	expectedChanged := true

	schedule.TargetTime, _ = time.Parse(time.RFC3339, "2018-06-11T00:00:00Z")
	changed, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err != nil {
		t.Fatalf("checker.Check() failed: %s", err)
	}

	if expectedChanged != changed {
		t.Fatal("expected changed departures, but they didn't")
	}

	if expectedChangedDepartureTimes[0] != changedDepartureTimes[0] {
		t.Fatalf("expected changed departure to be %s, got %s", expectedChangedDepartureTimes[0], changedDepartureTimes[0])
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

	schedule.TargetTime, _ = time.Parse(time.RFC3339, "2018-06-11T00:00:00Z")
	_, _, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err == nil {
		t.Fatal("expected checker.Check() to fail, but it didn't")
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

	expectedChangedDepartureTimes := []string{"No departure on 2018-06-11"}
	expectedChanged := true

	schedule.TargetTime, _ = time.Parse(time.RFC3339, "2018-06-11T00:00:00Z")
	changed, changedDepartureTimes, err := checker.Check(locationID, daysOfInterest, departureTime, trainName, limit)
	if err != nil {
		t.Fatalf("checker.Check() failed: %s", err)
	}

	if expectedChanged != changed {
		t.Fatal("expected changed departures, but they didn't")
	}

	if expectedChangedDepartureTimes[0] != changedDepartureTimes[0] {
		t.Fatalf("expected changed departure to be %s, got %s", expectedChangedDepartureTimes[0], changedDepartureTimes[0])
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
	"/departureBoard/87654321?date=2018-06-11": `
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
	"/departureBoard/87654321?date=2018-06-12": `
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
	"/departureBoard/87654321?date=2018-06-15": `
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
