package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidAudioNoMimes = errors.New("openrtb: audio has no mimes")
)

// Audio The "Audio" object must be included directly in the impression object if the impression offered
// for auction is an in-stream Audio ad opportunity.
type Audio struct {
	Mimes         []string        `json:"mimes,omitempty"`         // Content MIME types supported.
	MinDuration   int             `json:"minduration,omitempty"`   // Minimum Audio ad duration in seconds
	MaxDuration   int             `json:"maxduration,omitempty"`   // Maximum audio ad duration in seconds
	Protocols     []int           `json:"protocols,omitempty"`     // audio bid response protocols
	StartDelay    int             `json:"startdelay,omitempty"`    // Indicates the start delay in seconds
	Sequence      int             `json:"sequence,omitempty"`      // If multiple ad impressions are offered in the same bid request,the sequence number will allow for the coordinated delivery of multiple creatives.
	BAttr         []int           `json:"battr,omitempty"`         // Blocked creative attributes
	MaxExtended   int             `json:"maxextended,omitempty"`   // Maximum extended audio ad duration
	MinBitrate    int             `json:"minbitrate,omitempty"`    // Minimum bit rate in Kbps
	MaxBitrate    int             `json:"maxbitrate,omitempty"`    // Maximum bit rate in Kbps
	Delivery      []int           `json:"delivery,omitempty"`      // List of supported delivery methods
	CompanionAd   []Banner        `json:"companionad,omitempty"`   // Array of Banner objects (Section 3.2.3) if companion ads are available.
	API           []int           `json:"api,omitempty"`           // List of supported API frameworks
	CompanionType []int           `json:"companiontype,omitempty"` // Supported DAAST companion ad types
	MaxSeq        int             `json:"maxseq,omitempty"`        // The maximum number of ads that can be played in an ad pod
	Feed          int             `json:"feed,omitempty"`          // Type of audio feed
	Stitched      int             `json:"stitched,omitempty"`      // Indicates if the audio is stitched, where 0 = no, 1 = yes.
	NVol          int             `json:"nvol,omitempty"`          // Volume normalization mode.
	Download      int             `json:"dl,omitempty"`            // Indicates if the audio file can be downloaded by the user,where 0 = no, 1 = yes.
	Ext           json.RawMessage `json:"ext,omitempty"`
}

// Validate Validates the object
func (v *Audio) Validate() error {
	if len(v.Mimes) == 0 {
		return ErrInvalidAudioNoMimes
	}
	return nil
}
