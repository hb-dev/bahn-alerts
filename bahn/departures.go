package bahn

import (
	"fmt"
)

type Departure struct {
	TrainName string `json:"name"`
	BoardID   int    `json:"boardId"`
	StopID    int    `json:"stopId"`
	StopName  string `json:"stopName"`
	DateTime  string `json:"dateTime"`
	Track     string `json:"track"`
	DetailsID string `json:"detailsId"`
}

type DepartureCollection []Departure

func Departures(locationID int, date string) (*DepartureCollection, error) {
	return getDepartures(locationID, date)
}

func getDepartures(locationID int, date string) (*DepartureCollection, error) {
	path := fmt.Sprintf("departureBoard/%d?date=%s", locationID, date)

	departures := new(DepartureCollection)
	if err := getJSON(apiURL(path), departures); err != nil {
		return nil, err
	}
	return departures, nil
}
