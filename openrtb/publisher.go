package openrtb

import "encoding/json"

// Publisher This object describes the publisher of the media in which the ad will be displayed. The publisher is typically the seller in an OpenRTB transaction.
type Publisher struct {
	ID       string          `json:"id,omitempty"`     // specific publisher ID.
	Name     string          `json:"name,omitempty"`   // Publisher name (may be aliased at the publisher’s request)
	Category []string        `json:"cat,omitempty"`    // Array of IAB content categories of the app.
	Domain   string          `json:"domain,omitempty"` // Highest level domain of the publisher (e.g., “publisher.com”)
	Ext      json.RawMessage `json:"ext,omitempty"`
}
