package dealmap

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"io/ioutil"
)

// Dealmap API information
const (
	ApiUrl = "http://api.thedealmap.com"
)

// DealMap class to used to communicate with the API
type DealMap struct {
	baseUrl    string       // The base URL to access the API
	httpClient *http.Client // The HTTP client to use
	key        string       // The developer key provided by TheDealMap
}

// Creates a new DealMap object using the given HTTP client and API key
func New(apiUrl string, client *http.Client, apiKey string) *DealMap {
	return &DealMap{apiUrl, client, apiKey}
}

// buildQuery returns a map to use as the query string in requests to the API.
// The given parameters are included if they are non-empty.
func (dm *DealMap) buildQuery(location string, query string, distanceMi int, startIndex int, pageSize int, activity int, capability int, expiration time.Time) url.Values {
	args := make(map[string][]string)
	if location != "" {
		args["l"] = []string{location}
	}
	if query != "" && query != "*" {
		args["q"] = []string{query}
	}
	if distanceMi > 0 {
		args["d"] = []string{strconv.Itoa(distanceMi)}
	}
	if startIndex >= 0 {
		args["si"] = []string{strconv.Itoa(startIndex)}
	}
	if pageSize > 0 {
		args["ps"] = []string{strconv.Itoa(pageSize)}
	}
	if activity > 0 {
		args["a"] = []string{strconv.Itoa(activity)}
	}
	if capability > 0 {
		args["c"] = []string{strconv.Itoa(capability)}
	}
	if expiration.Unix() > 0 {
		args["ed"] = []string{expiration.Format("2006-01-02")}
	}
	args["key"] = []string{dm.key}
	return args
}

// SearchDeals invokes the "search deals" API from TheDealMap and returns the response as a Deals object.
func (dm *DealMap) SearchDeals(location string, query string, distanceMi int, startIndex int, pageSize int, activity int, capability int, expirationDate time.Time) (*Deals, error) {
	args := dm.buildQuery(location, query, distanceMi, startIndex, pageSize, activity, capability, expirationDate)
	resp, err := dm.httpClient.Get(fmt.Sprintf("%s/search/deals/?%s", dm.baseUrl, args.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result Deals
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchBusinesses invokes the "search businesses" API from TheDealMap and returns the response as a Businesses object.
func (dm *DealMap) SearchBusinesses(location string, query string, distanceMi int, startIndex int, pageSize int, activity int) (*Businesses, error) {
	args := dm.buildQuery(location, query, distanceMi, startIndex, pageSize, activity, 0, time.Unix(0, 0))
	resp, err := dm.httpClient.Get(fmt.Sprintf("%s/search/businesses/?%s", dm.baseUrl, args.Encode()))

	if err != nil {
		return nil, err
	}
	var result Businesses
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DealDetails invokes the "deal details" API from TheDeapMap and returns the response as a Deal object.
func (dm *DealMap) DealDetails(id string) (*Deal, error) {
	args := dm.buildQuery("", "", 0, 0, 0, 0, 0, time.Unix(0, 0))
	resp, err := dm.httpClient.Get(fmt.Sprintf("%s/deals/%s/?%s", dm.baseUrl, id, args.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result Deal
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
