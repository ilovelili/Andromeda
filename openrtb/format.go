package openrtb

import (
	"encoding/json"
)

// Format This object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	Width  int             `json:"w,omitempty"` // Width
	Height int             `json:"h,omitempty"` // Height
	Ext    json.RawMessage `json:"ext,omitempty"`
}
