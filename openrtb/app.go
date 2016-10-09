package openrtb

import "encoding/json"

// App This object should be included if the ad supported content is a non-browser application (typically inmobile) as opposed to a website
// A bid request must not contain both an App and a Site object. At aminimum, it is useful to provide an App ID or bundle, but this is not strictly required
type App struct {
	ID              string          `json:"id,omitempty"`            // recommended Exchange-specific app ID.
	Name            string          `json:"name,omitempty"`          // App name (may be aliased at the publisher’s request).
	Bundle          string          `json:"bundle,omitempty"`        // A platform-specific application identifier intended to be unique to the app and independent of the exchange
	Domain          string          `json:"domain,omitempty"`        // Domain of the app (e.g., “mygame.foo.com”)
	StoreURL        string          `json:"storeurl,omitempty"`      // App store URL for an installed app; for IQG 2.1 compliance.
	Category        []string        `json:"cat,omitempty"`           // Array of IAB content categories of the app.
	SectionCategory []string        `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current section of the app
	PageCategory    []string        `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the app
	Version         string          `json:"ver,omitempty"`           // Application version
	PrivacyPolicy   int             `json:"privacypolicy,omitempty"` // Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	Paid            int             `json:"paid,omitempty"`          // 0 = app is free, 1 = the app is a paid version
	Publisher       *Publisher      `json:"publisher,omitempty"`     // Details about the Publisher
	Content         *Content        `json:"content,omitempty"`       // Details about the Content
	Keywords        string          `json:"keywords,omitempty"`      // Comma separated list of keywords about the app
	Ext             json.RawMessage `json:"ext,omitempty"`
}
