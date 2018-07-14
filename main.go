package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const APIURL = "https://api.deutschebahn.com/freeplan/v1"

var bahnClient = &http.Client{Timeout: 10 * time.Second}

type departure struct {
	TrainName string `json:"name"`
	BoardID   int    `json:"boardId"`
	StopID    int    `json:"stopId"`
	StopName  string `json:"stopName"`
	DateTime  string `json:"dateTime"`
	Track     string `json:"track"`
	DetailsID string `json:"detailsId"`
}

type Departures []departure

func main() {

	// get data from bahn api
	locationID := 8011956 // Jena Paradies

	departures, err := getDepartures(locationID, "2018-06-12")
	if err != nil {
		panic(err)
	}

	for _, departure := range *departures {
		fmt.Printf("%s - %s\n", departure.TrainName, departure.DateTime)
	}

	// filter data
	// define schedule parameters
	// check for changed schedule
	// define alert parameters
	// alert changed schedule
}

func getDepartures(locationID int, date string) (*Departures, error) {
	path := fmt.Sprintf("departureBoard/%d?date=%s", locationID, date)

	departures := new(Departures)
	if err := getJSON(apiURL(path), departures); err != nil {
		return nil, err
	}
	return departures, nil
}

func getJSON(url string, target interface{}) error {
	r, err := bahnClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func apiURL(path string) string {
	return APIURL + "/" + path
}
