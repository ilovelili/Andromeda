package openrtb

import "encoding/json"

// User This object contains information known or derived about the human user of the device (i.e., the
// audience for advertising). The user id is an exchange artifact and may be subject to rotation or other
// privacy policies. However, this user ID must be stable long enough to serve reasonably as the basis for
// frequency capping and retargeting.
type User struct {
	ID          string          `json:"id,omitempty"`         // Unique consumer ID of this user on the exchange
	BuyerID     string          `json:"buyerid,omitempty"`    // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyeruid/buyerid or id is recommended. Valid for OpenRTB 2.3.
	BuyerUID    string          `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. Same as BuyerID but valid for OpenRTB 2.2.
	YearOfBirth int             `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender      string          `json:"gender,omitempty"`     // Gender ("M": male, "F" female, "O" Other)
	Keywords    string          `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent
	CustomData  string          `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange's cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include "escaped" quotation marks.
	Geo         *Geo            `json:"geo,omitempty"`
	Data        []Data          `json:"data,omitempty"`
	Ext         json.RawMessage `json:"ext,omitempty"`
}
