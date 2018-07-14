package bahn_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hb-dev/bahn-alerts/bahn"
)

func TestGetDepartures(t *testing.T) {
	ts := exampleBahnApiServer()
	defer ts.Close()
	bahn.ApiURL = ts.URL

	departures, err := bahn.Departures(123, "date")
	if err != nil {
		t.Fatalf("bahn.Departures() failed: %s", err)
	}

	if len(*departures) != 3 {
		t.Fatalf("expected 3 departures, got: %d", len(*departures))
	}

	if (*departures)[0].TrainName != "ICE 1526" {
		t.Fatalf("expect train name to be ICE 1526, got: %s", (*departures)[0].TrainName)
	}
}

func exampleBahnApiServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, departuresApiResponse)
	}))
}

var departuresApiResponse = `
    [
      {
        "name": "ICE 1526",
        "boardId": 8011956,
        "stopId": 8011956,
        "stopName": "Jena Paradies",
        "dateTime": "2018-06-12T06:51",
        "track": "2",
        "detailsId": "514122%2F172894%2F792392%2F224822%2F80%3fstation_evaId%3D8011956"
      },
      {
        "name": "IC 2063",
        "type": "IC",
        "boardId": 8011956,
        "stopId": 8011956,
        "stopName": "Jena Paradies",
        "dateTime": "2018-06-12T13:00",
        "track": "2",
        "detailsId": "475920%2F162924%2F364760%2F23740%2F80%3fstation_evaId%3D8011956"
      },
      {
        "name": "IC 2060",
        "type": "IC",
        "boardId": 8011956,
        "stopId": 8011956,
        "stopName": "Jena Paradies",
        "dateTime": "2018-06-12T17:00",
        "track": "1",
        "detailsId": "193857%2F68878%2F100874%2F14182%2F80%3fstation_evaId%3D8011956"
      }
    ]
  `
