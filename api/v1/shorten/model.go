package shorten

type RequestCreate struct {
	URL       string `json:"url"`
	Shortcode string `json:"shortcode"`
}

type ResponseCreate struct {
	Shortcode string `json:"shortcode"`
}

type ResponseStatus struct {
	StartDate     string `json:"startDate"`
	LastSeenDate  string `json:"lastSeenDate,omitempty"`
	RedirectCount int64  `json:"redirectCount"`
}

type ResponseList struct {
	Items []Shorten `json:"items"`
	Total int       `json:"total"`
}
type Shorten struct {
	ID            int    `json:"id"`
	Shortcode     string `json:"shortcode"`
	Url           string `json:"url"`
	RedirectCount int64  `json:"redirectCount"`
	LastSeenDate  string `json:"lastSeenDate"`
	StartDate     string `json:"startDate"`
}
