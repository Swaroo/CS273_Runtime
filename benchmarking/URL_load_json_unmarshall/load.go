package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Station struct {
	ExecutionTime   string `json:"executionTime"`
	StationBeanList []struct {
		ID                    int     `json:"id"`
		StationName           string  `json:"stationName"`
		AvailableDocks        int     `json:"availableDocks"`
		TotalDocks            int     `json:"totalDocks"`
		Latitude              float64 `json:"latitude"`
		Longitude             float64 `json:"longitude"`
		StatusValue           string  `json:"statusValue"`
		StatusKey             int     `json:"statusKey"`
		AvailableBikes        int     `json:"availableBikes"`
		StAddress1            string  `json:"stAddress1"`
		StAddress2            string  `json:"stAddress2"`
		City                  string  `json:"city"`
		PostalCode            string  `json:"postalCode"`
		Location              string  `json:"location"`
		Altitude              string  `json:"altitude"`
		TestStation           bool    `json:"testStation"`
		LastCommunicationTime string  `json:"lastCommunicationTime"`
		LandMark              string  `json:"landMark"`
	} `json:"stationBeanList"`
}

func serial() {
	url := "http://feeds.citibikenyc.com/stations/stations.json"

	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	for i := 0; i < 50; i++ {

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, getErr := client.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		stationList := Station{}
		jsonErr := json.Unmarshal(body, &stationList)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		_ = stationList.ExecutionTime
	}
}

func doWork(ch chan string, url string, client http.Client) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	stationList := Station{}
	jsonErr := json.Unmarshal(body, &stationList)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ch <- stationList.ExecutionTime
}

func parallel() {
	ch := make(chan string, 50)
	url := "http://feeds.citibikenyc.com/stations/stations.json"

	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	for i := 0; i < 50; i++ {
		go doWork(ch, url, client)
	}
	for i := 0; i < 50; i++ {
		_ = <-ch
	}
}

func main() {

}
