package schedule_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hb-dev/bahn-alerts/bahn"
	"github.com/hb-dev/bahn-alerts/schedule"
)

func TestSchedule(t *testing.T) {
	ts := exampleBahnAPIServer()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	trainName := "ICE 123"
	daysOfInterest := []string{"Monday", "Tuesday", "Friday"}
	limit := 3

	expected := []string{
		"2018-06-11T06:51",
		"2018-06-12T06:51",
		"2018-06-15T06:51",
	}

	schedule.TargetTime, _ = time.Parse(time.RFC3339, "2018-06-11T00:00:00Z")
	trainSchedule, err := schedule.Schedule(locationID, trainName, daysOfInterest, limit)
	if err != nil {
		t.Fatalf("schedule.Schedule() failed: %s", err)
	}

	if expected[0] != trainSchedule[0] {
		t.Fatalf("expected schedule item to be %s, got %s", expected[0], trainSchedule[0])
	}
}

func TestScheduleBahnAPIError(t *testing.T) {
	ts := exampleFailingBahnAPIServer()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	trainName := "ICE 123"
	daysOfInterest := []string{"Monday"}
	limit := 3

	_, err := schedule.Schedule(locationID, trainName, daysOfInterest, limit)
	if err == nil {
		t.Fatal("expected schedule.Schedule() to fail, but it didn't")
	}
}

func TestScheduleBahnAPIEmptyResponse(t *testing.T) {
	ts := exampleBahnAPIServerEmptyResonse()
	defer ts.Close()
	bahn.APIURL = ts.URL

	locationID := 87654321
	trainName := "ICE 123"
	daysOfInterest := []string{"Monday"}
	limit := 3

	expectedMessage := "No departure on 2018-06-11"

	schedule.TargetTime, _ = time.Parse(time.RFC3339, "2018-06-11T00:00:00Z")
	trainSchedule, err := schedule.Schedule(locationID, trainName, daysOfInterest, limit)
	if err != nil {
		t.Fatalf("schedule.Schedule() failed: %s", err)
	}

	if expectedMessage != trainSchedule[0] {
		t.Fatalf("expected schedule message to be %s, but got %s", expectedMessage, trainSchedule[0])
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
