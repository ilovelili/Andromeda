package openrtb

import "encoding/json"

// Site This object should be included if the ad supported content is a website as opposed to a non-browser application
// A bid request must not contain both a Site and an App object. At a minimum, it is useful to provide a site ID or page URL, but this is not strictly required.
type Site struct {
	ID              string          `json:"id,omitempty"`            // recommended Exchange-specific site ID.
	Name            string          `json:"name,omitempty"`          // Site name (may be aliased at the publisher’s request).
	Domain          string          `json:"domain,omitempty"`        // Domain of the site (e.g., “mysite.foo.com”).
	Category        []string        `json:"cat,omitempty"`           // Array of IAB content categories of the site.
	SectionCategory []string        `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current section of the site
	PageCategory    []string        `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the site
	Page            string          `json:"page,omitempty"`          // URL of the page where the impression will be shown.
	Referrence      string          `json:"ref,omitempty"`           // Referrer URL that caused navigation to the current page.
	Search          string          `json:"search,omitempty"`        // Search string that caused navigation to the current page.
	Mobile          int             `json:"mobile,omitempty"`        // Indicates if the site has been programmed to optimize layout when viewed on mobile devices, where 0 = no, 1 = yes.
	PrivacyPolicy   int             `json:"privacypolicy,omitempty"` // Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	Publisher       *Publisher      `json:"publisher,omitempty"`     // Details about the Publisher of the site.
	Content         *Content        `json:"content,omitempty"`       // Details about the Content within the site.
	Keywords        string          `json:"keywords,omitempty"`      // Comma separated list of keywords about the site
	Ext             json.RawMessage `json:"ext,omitempty"`
}
