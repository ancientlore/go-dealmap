package dealmap

import (
	"http"
	"fmt"
	"strconv"
	"xml"
	"time"
	"os"
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
func (dm *DealMap) buildQuery(location string, query string, distanceMi int, startIndex int, pageSize int, activity int, capability int, expiration *time.Time) map[string][]string {
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
	if expiration != nil {
		args["ed"] = []string{expiration.Format("2006-01-02")}
	}
	args["key"] = []string{dm.key}
	return args
}

// SearchDeals invokes the "search deals" API from TheDealMap and returns the response as a Deals object.
func (dm *DealMap) SearchDeals(location string, query string, distanceMi int, startIndex int, pageSize int, activity int, capability int, expirationDate *time.Time) (*Deals, os.Error) {
	args := dm.buildQuery(location, query, distanceMi, startIndex, pageSize, activity, capability, expirationDate)
	resp, _, err := dm.httpClient.Get(fmt.Sprintf("%s/search/deals/?%s", dm.baseUrl, http.EncodeQuery(args)))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result Deals
	err = xml.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchBusinesses invokes the "search businesses" API from TheDealMap and returns the response as a Businesses object.
func (dm *DealMap) SearchBusinesses(location string, query string, distanceMi int, startIndex int, pageSize int, activity int) (*Businesses, os.Error) {
	args := dm.buildQuery(location, query, distanceMi, startIndex, pageSize, activity, 0, nil)
	resp, _, err := dm.httpClient.Get(fmt.Sprintf("%s/search/businesses/?%s", dm.baseUrl, http.EncodeQuery(args)))
	if err != nil {
		return nil, err
	}
	var result Businesses
	err = xml.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DealDetails invokes the "deal details" API from TheDeapMap and returns the response as a Deal object.
func (dm *DealMap) DealDetails(id string) (*Deal, os.Error) {
	args := dm.buildQuery("", "", 0, 0, 0, 0, 0, nil)
	resp, _, err := dm.httpClient.Get(fmt.Sprintf("%s/deals/%s/?%s", dm.baseUrl, id, http.EncodeQuery(args)))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var result Deal
	err = xml.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
