package openrtb

import "encoding/json"

// Pmp Private Marketplace Object
type Pmp struct {
	PrivateAuction int             `json:"private_auction,omitempty"`
	Deals          []Deal          `json:"deals,omitempty"`
	Ext            json.RawMessage `json:"ext,omitempty"`
}
