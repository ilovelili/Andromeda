package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidVideoNoMimes = errors.New("openrtb: video has no mimes")
)

// Video The "video" object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	Mimes          []string        `json:"mimes,omitempty"`          // Content MIME types supported.
	MinDuration    int             `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration    int             `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Protocol       int             `json:"protocol,omitempty"`       // Video bid response protocols
	Protocols      []int           `json:"protocols,omitempty"`      // Video bid response protocols
	Width          int             `json:"w,omitempty"`              // Width of the player in pixels
	Height         int             `json:"h,omitempty"`              // Height of the player in pixels
	StartDelay     int             `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	Linearity      int             `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Skip           int             `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped,where 0 = no, 1 = yes
	SkipMin        int             `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable;
	SkipAfter      int             `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence       int             `json:"sequence,omitempty"`       // If multiple ad impressions are offered in the same bid request,the sequence number will allow for the coordinated delivery of multiple creatives.
	BAttr          []int           `json:"battr,omitempty"`          // Blocked creative attributes
	MaxExtended    int             `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	MinBitrate     int             `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	MaxBitrate     int             `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	BoxingAllowed  int             `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	PlaybackMethod []int           `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery       []int           `json:"delivery,omitempty"`       // List of supported delivery methods
	Pos            int             `json:"pos,omitempty"`            // Ad Position
	CompanionAd    []Banner        `json:"companionad,omitempty"`    // Array of Banner objects (Section 3.2.3) if companion ads are available.
	API            []int           `json:"api,omitempty"`            // List of supported API frameworks
	CompanionType  []int           `json:"companiontype,omitempty"`
	Ext            json.RawMessage `json:"ext,omitempty"`
}

// Validate Validates the object
func (v *Video) Validate() error {
	if len(v.Mimes) == 0 {
		return ErrInvalidVideoNoMimes
	}

	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}
