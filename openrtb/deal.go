package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidDealNoID = errors.New("openrtb: deal has no ID")
)

// Deal PMP Deal
type Deal struct {
	ID               string          `json:"id"` // Unique deal ID
	BidFloor         float64         `json:"bidfloor,omitempty"`
	BidFloorCurrency string          `json:"bidfloorcur,omitempty"` // Currency of bid floor
	AuctionType      int             `json:"at,omitempty"`          // Optional override of the overall auction type of the bid request.
	WSeat            []string        `json:"wseat,omitempty"`       // Array of buyer seats allowed to bid on this Direct Deal.
	WAdvDomain       []string        `json:"wadomain,omitempty"`    // Array of advertiser domains allowed to bid on this Direct Deal
	Ext              json.RawMessage `json:"ext,omitempty"`
}

// Validate Validates the object
func (d *Deal) Validate() error {
	if len(d.ID) == 0 {
		return ErrInvalidVideoNoMimes
	}

	return nil
}
