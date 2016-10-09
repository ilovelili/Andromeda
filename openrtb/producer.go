package openrtb

import "encoding/json"

// Producer This object defines the producer of the content in which the ad will be shown. This is particularly useful
// when the content is syndicated and may be distributed through different publishers and thus when the
// producer and publisher are not necessarily the same entity.
type Producer struct {
	ID       string          `json:"id,omitempty"`     // ID uniquely identifying the producer.
	Name     string          `json:"name,omitempty"`   // Content producer or originator name (e.g., “Warner Bros”).
	Category []string        `json:"cat,omitempty"`    // Array of IAB content categories that describe the content producer
	Domain   string          `json:"domain,omitempty"` // Highest level domain of the content producer (e.g.,“producer.com”)
	Ext      json.RawMessage `json:"ext,omitempty"`
}
