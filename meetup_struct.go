package pale

type MeetupEvent struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	DateInSeriesPattern bool   `json:"date_in_series_pattern,omitempty"`
	Status              string `json:"status,omitempty"`
	Time                int64  `json:"time"`
	LocalDate           string `json:"local_date"`
	LocalTime           string `json:"local_time"`
	WaitlistCount       int    `json:"waitlist_count"`
	YesRsvpCount        int    `json:"yes_rsvp_count"`
	Venue               struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Lat      float64 `json:"lat"`
		Lon      float64 `json:"lon"`
		Address1 string  `json:"address_1"`
		City     string  `json:"city"`
		Country  string  `json:"country"`
	} `json:"venue,omitempty"`
	Group struct {
		Name    string `json:"name"`
		ID      int    `json:"id"`
		Urlname string `json:"urlname"`
	} `json:"group"`
	Link        string `json:"link"`
	Description string `json:"description,omitempty"`
	RsvpLimit   int    `json:"rsvp_limit,omitempty"`
	HowToFindUs string `json:"how_to_find_us,omitempty"`
}

type MeetupEvents struct {
	Events []MeetupEvent `json:"events"`
}
