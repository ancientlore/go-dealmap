package dealmap

import (
	"time"
	"os"
)

// Centered activity Values
const (
	Kids      = 1
	Group     = 8
	Romantic  = 16
	Casual    = 64
	Fun       = 512
	LateNight = 16384
	Outdoor   = 32768
)

// Unknown value
const (
	Unknown = 0
)

// Deal types
const (
	GiftCertificate = 1
	BOGO            = 2
	PrintableCoupon = 4
	GroupBuy        = 8
	DailyDeal       = 16
	FreeDeal        = 32
)

// Deal capability values
const (
	Favorite           = 1  // Favorited by The Dealmap Deal Editors
	HasTransaction     = 2  // Involves some kind of purchase
	Featured           = 4  // This deal is featured on The Dealmap
	Exclusive          = 8  // This deal is exclusively available only at The Dealmap
	FiftyPercentOrMore = 16 // This deal has 50  or more % savings
	CanBePrinted       = 32 // This is a printable coupon
	Affiliate          = 64 // Affiliate
)

// Deal units
const (
	Price      = 1
	Percentage = 2
)

// Deal currency
const (
	USD = 1
	GBP = 2
	EUR = 3
)

// Deals is an API response containing a list of deals and an API message
type Deals struct {
	Message      string // Message from API, usually indicating a problem of some kind
	Results      []Deal `xml:"Results>Deal"` // List of Deal objects
	TotalResults int    // Total number of results
}

// A Deal contains all the details of a single deal. See TheDealMap API documentation for details.
type Deal struct {
	Activity                               int
	AddedBy                                string
	AdditionalDiscountCouponCode           string
	AdditionalDiscountCouponEffectiveTime  string
	AdditionalDiscountCouponExpirationTime string
	AdditionalDiscountDealUnit             int
	AdditionalDiscountedValue              int
	AddressLine                            string
	Affiliation                            string
	BDescription                           string
	BusinessID                             string
	BusinessName                           string
	Capability                             int
	Category                               string
	City                                   string
	Country                                string
	Currency                               int
	DealSource                             string
	DealType                               int
	DealUnit                               int
	Description                            string
	DiscountedValue                        float64
	EffectiveTime                          string
	ExpirationTime                         string
	FaceValue                              float64
	ID                                     string
	IconUrl                                string
	ImageUrl                               string
	IsSoldOut                              bool
	Keywords                               string
	Latitude                               float64
	Longitude                              float64
	MoreInfoLink                           string
	Phone                                  string
	Ratings                                float64
	ReviewCount                            int
	SoldCount                              int
	State                                  string
	Tags                                   string
	Terms                                  string
	Title                                  string
	TrackingUrl                            string
	TransactionUrl                         string
	YouSave                                string
	ZipCode                                string
}

// Attempts to parse the additional discount coupon's effective time
func (d *Deal) ParseAdditionalDiscountCouponEffectiveTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.AdditionalDiscountCouponEffectiveTime)
}

// Attempts to parse the additional discount coupon's expiration time
func (d *Deal) ParseAdditionalDiscountCouponExpirationTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.AdditionalDiscountCouponExpirationTime)
}

// Attempts to parse the deal's effective time
func (d *Deal) ParseEffectiveTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.EffectiveTime)
}

// Attempts to parse the deal's expiration time
func (d *Deal) ParseExpirationTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.ExpirationTime)
}

// Businesses is an API response containing a list of Business objects and a response message
type Businesses struct {
	Message      string     // Message from API, usually indicating a problem of some kind
	Results      []Business `xml:"Results>Business"` // List of Business objects
	TotalResults int        // Total number of results
}

// Business contains the details of a business in TheDealMap. See the API documentation for details.
type Business struct {
	Activity     int
	AddressLine  string
	Capability   int
	Category     string
	City         string
	Country      string
	Franchise    string
	ID           string
	Latitude     float64
	Longitude    float64
	MoreInfoLink string
	Phone        string
	Ratings      float64
	Reviews      int
	State        string
	Title        string
}
