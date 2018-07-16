package bahn

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJson(t *testing.T) {
	ts := exampleBahnAPIServer()
	defer ts.Close()

	expectedDeparture := Departure{
		TrainName: "ICE 1526",
		BoardID:   8011956,
		StopID:    8011956,
		StopName:  "Jena Paradies",
		DateTime:  "2018-07-16T06:51",
		Track:     "2",
		DetailsID: "308442%2F104334%2F294972%2F44672%2F80%3fstation_evaId%3D8011956",
	}

	departures := new(DepartureCollection)
	err := getJSON(ts.URL, departures)
	if err != nil {
		t.Fatalf("getJSON() failed: %s", err)
	}

	if expectedDeparture != (*departures)[0] {
		t.Fatalf("expected Departure to be %v, got: %v", expectedDeparture, (*departures)[0])
	}
}

func TestGetJsonWithAPIError(t *testing.T) {
	ts := exampleFailingBahnAPIServer()
	defer ts.Close()

	departures := new(DepartureCollection)
	err := getJSON(ts.URL, departures)
	if err == nil {
		t.Fatal("expected getJSON() to fail, but it didn't")
	}
}

func TestGetJsonWithErrorForEmptyResponse(t *testing.T) {
	ts := exampleBahnAPIServerEmptyResonse()
	defer ts.Close()

	departures := new(DepartureCollection)
	err := getJSON(ts.URL, departures)
	if err != nil {
		t.Fatalf("getJSON() failed: %s", err)
	}

	if len(*departures) > 0 {
		t.Fatalf("expected result to be 0 departures, got: %d", len(*departures))
	}
}

func exampleBahnAPIServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, departuresAPIResponse)
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

var departuresAPIResponse = `
[
  {
    "name": "ICE 1526",
    "boardId": 8011956,
    "stopId": 8011956,
    "stopName": "Jena Paradies",
    "dateTime": "2018-07-16T06:51",
    "track": "2",
    "detailsId": "308442%2F104334%2F294972%2F44672%2F80%3fstation_evaId%3D8011956"
  },
  {
    "name": "IC 2063",
    "type": "IC",
    "boardId": 8011956,
    "stopId": 8011956,
    "stopName": "Jena Paradies",
    "dateTime": "2018-07-16T13:00",
    "track": "2",
    "detailsId": "336453%2F116428%2F624222%2F199960%2F80%3fstation_evaId%3D8011956"
  },
  {
    "name": "IC 2060",
    "type": "IC",
    "boardId": 8011956,
    "stopId": 8011956,
    "stopName": "Jena Paradies",
    "dateTime": "2018-07-16T17:00",
    "track": "1",
    "detailsId": "694842%2F235865%2F326064%2F68582%2F80%3fstation_evaId%3D8011956"
  }
]
  `
