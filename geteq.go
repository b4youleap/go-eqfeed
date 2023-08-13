package main

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type EQ struct {
	Type     string    `json:"type"`
	Metadata Metadata  `json:"metadata"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	ID         string     `json:"id"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Mag     float64     `json:"mag"`
	Place   string      `json:"place"`
	Time    int64       `json:"time"`
	Updated int64       `json:"updated"`
	Tz      interface{} `json:"tz"`
	URL     string      `json:"url"`
	Detail  string      `json:"detail"`
	Felt    int64       `json:"felt"`
	Cdi     float64     `json:"cdi"`
	MMI     float64     `json:"mmi"`
	Alert   interface{} `json:"alert"`
	Status  string      `json:"status"`
	Tsunami int64       `json:"tsunami"`
	Sig     int64       `json:"sig"`
	Net     string      `json:"net"`
	Code    string      `json:"code"`
	IDS     string      `json:"ids"`
	Sources string      `json:"sources"`
	Types   string      `json:"types"`
	Nst     int64       `json:"nst"`
	Dmin    float64     `json:"dmin"`
	RMS     float64     `json:"rms"`
	Gap     int64       `json:"gap"`
	MagType string      `json:"magType"`
	Type    string      `json:"type"`
	Title   string      `json:"title"`
}

type Metadata struct {
	Generated int64  `json:"generated"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Status    int64  `json:"status"`
	API       string `json:"api"`
	Count     int64  `json:"count"`
}

func main() {

	url := "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var eqData EQ
	err = json.Unmarshal(body, &eqData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("EQ data: %+v\n", eqData)

}
