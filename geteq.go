package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/feeds"
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
		log.Fatal("Error fetching data:", err) // https://pkg.go.dev/log#Fatal
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	var eqData EQ
	err = json.Unmarshal(body, &eqData)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	resp.Body.Close()

	feed := &feeds.Feed{
		Title:       eqData.Metadata.Title,
		Link:        &feeds.Link{Href: url},
		Description: "USGS Earthquake Hazards Program, responsible for monitoring, reporting, and researching earthquakes and earthquake hazards",
		Author:      &feeds.Author{Name: "USGS", Email: "eq_questions@usgs.gov"},
		Created:     time.Now(),
	}

	for _, earthquake := range eqData.Features {

		eqDate := time.UnixMilli(earthquake.Properties.Updated) // https://pkg.go.dev/time#Parse

		item := &feeds.Item{
			Title:   earthquake.Properties.Title,
			Link:    &feeds.Link{Href: earthquake.Properties.URL},
			Created: eqDate,
		}

		feed.Items = append(feed.Items, item)

	}

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rss)

}
