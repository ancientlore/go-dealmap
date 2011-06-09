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
	baseUrl    string
	httpClient *http.Client
	key        string
}

// Creates a new TripIt object using the given HTTP client and authorization object
func New(apiUrl string, client *http.Client, apiKey string) *DealMap {
	return &DealMap{apiUrl, client, apiKey}
}

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

func (dm *DealMap) SearchBusinesses(location string, query string, distanceMi int, startIndex int, pageSize int, activity int) (interface{}, os.Error) {
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
