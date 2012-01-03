package main

import (
	"../_obj/dealmap"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"strings"
	"time"
)

var key = flag.String("key", "", "API key")
var action = flag.String("action", "SearchDeals", "Action to take")
var location = flag.String("l", "Washington, DC", "Location")
var query = flag.String("q", "indian food", "Query")
var distanceMi = flag.Int("d", 0, "Distance in miles")
var startIndex = flag.Int("si", 0, "Start index")
var pageSize = flag.Int("ps", 0, "Page size")
var activity = flag.Int("a", 0, "Activity")
var capability = flag.Int("c", 0, "Capability")
var expirationDate = flag.String("ed", "", "Expiration date")
var id = flag.String("id", "", "Deal ID")

func main() {
	flag.Parse()
	var client http.Client
	dm := dealmap.New(dealmap.ApiUrl, &client, *key)
	var ed time.Time
	var err error
	if *expirationDate != "" {
		ed, err = time.Parse("01-02-2006", *expirationDate)
		if err != nil {
			log.Fatal(err)
		}
	}
	var b []byte
	switch strings.ToLower(*action) {
	case "searchdeals":
		deals, err := dm.SearchDeals(*location, *query, *distanceMi, *startIndex, *pageSize, *activity, *capability, ed)
		if err != nil {
			log.Fatal(err)
		}
		b, err = json.MarshalIndent(deals, "", "    ")
	case "searchbusinesses":
		businesses, err := dm.SearchBusinesses(*location, *query, *distanceMi, *startIndex, *pageSize, *activity)
		if err != nil {
			log.Fatal(err)
		}
		b, err = json.MarshalIndent(businesses, "", "    ")
	case "dealdetails":
		deals, err := dm.DealDetails(*id)
		if err != nil {
			log.Fatal(err)
		}
		b, err = json.MarshalIndent(deals, "", "    ")
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(b))
}
