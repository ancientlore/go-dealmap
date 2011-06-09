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

type Deals struct {
	Message      string
	Results      []Deal "Results>Deal"
	TotalResults int
}

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
	DiscountedValue                        int
	EffectiveTime                          string
	ExpirationTime                         string
	FaceValue                              int
	ID                                     string
	IconUrl                                string
	ImageUrl                               string
	IsSoldOut                              bool
	Keywords                               string
	Latitude                               float64
	Longitude                              float64
	MoreInfoLink                           string
	Phone                                  string
	Ratings                                int
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

func (d *Deal) ParseAdditionalDiscountCouponEffectiveTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.AdditionalDiscountCouponEffectiveTime)
}

func (d *Deal) ParseAdditionalDiscountCouponExpirationTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.AdditionalDiscountCouponExpirationTime)
}

func (d *Deal) ParseEffectiveTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.EffectiveTime)
}

func (d *Deal) ParseExpirationTime() (*time.Time, os.Error) {
	return time.Parse(time.RFC3339, d.ExpirationTime)
}

type Businesses struct {
	Message      string
	Results      []Business "Results>Business"
	TotalResults int
}

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
	Ratings      int
	Reviews      int
	State        string
	Title        string
}
