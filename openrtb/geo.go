package openrtb

import "encoding/json"

// Geo Note that the Geo Object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Latitude      float64         `json:"lat,omitempty"`           // Latitude from -90 to 90
	Longitude     float64         `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          int             `json:"type,omitempty"`          // Source of location data; recommended when passing lat / lon.
	Accuracy      int             `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services (i.e., type = 1).
	LastFix       int             `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int             `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable (i.e., type = 2)
	Country       string          `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string          `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string          `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string          `json:"metro,omitempty"`
	City          string          `json:"city,omitempty"`
	Zip           string          `json:"zip,omitempty"`
	UTCOffset     int             `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           json.RawMessage `json:"ext,omitempty"`
}
