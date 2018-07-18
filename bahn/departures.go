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

func Departures(locationID int, dateTime string) (*DepartureCollection, error) {
	return getDepartures(locationID, dateTime)
}

func getDepartures(locationID int, dateTime string) (*DepartureCollection, error) {
	path := fmt.Sprintf("departureBoard/%d?date=%s", locationID, dateTime)

	departures := new(DepartureCollection)
	if err := getJSON(apiURL(path), departures); err != nil {
		return nil, err
	}
	return departures, nil
}
