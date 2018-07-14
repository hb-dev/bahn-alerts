package bahn

import (
	"fmt"
)

type departure struct {
	TrainName string `json:"name"`
	BoardID   int    `json:"boardId"`
	StopID    int    `json:"stopId"`
	StopName  string `json:"stopName"`
	DateTime  string `json:"dateTime"`
	Track     string `json:"track"`
	DetailsID string `json:"detailsId"`
}

type departures []departure

func Departures(locationID int, date string) (*departures, error) {
	return getDepartures(locationID, date)
}

func getDepartures(locationID int, date string) (*departures, error) {
	path := fmt.Sprintf("departureBoard/%d?date=%s", locationID, date)

	d := new(departures)
	if err := getJSON(apiURL(path), d); err != nil {
		return nil, err
	}
	return d, nil
}
