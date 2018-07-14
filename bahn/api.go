package bahn

import (
	"encoding/json"
	"net/http"
	"time"
)

const APIURL = "https://api.deutschebahn.com/freeplan/v1"

var bahnClient = &http.Client{Timeout: 10 * time.Second}

func apiURL(path string) string {
	return APIURL + "/" + path
}

func getJSON(url string, target interface{}) error {
	r, err := bahnClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
