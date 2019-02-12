package pale

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type MeetupM struct {
	m      Meetup
	client *http.Client
}

func NewMeetupM() *MeetupM {

	return &MeetupM{
		m: *NewMeetup(),
		client: &http.Client{
			// TODO: move it to env var
			Timeout: time.Second * 15,
		},
	}
}

func (m *MeetupM) fetch() Meetup {
	fmt.Println("pulling tweets from Meetup.com")

	params := map[string]string{
		"key":    os.Getenv("MEETUP_API_KEY"),
		"lat":    os.Getenv("MEETUP_API_LAT"),
		"lon":    os.Getenv("MEETUP_API_LON"),
		"radius": os.Getenv("MEETUP_API_RADIUS"),
	}
	urlpath := []string{"/find/upcoming_events?"}
	for k, v := range params {
		param := k + "=" + v
		urlpath = append(urlpath, param)
	}
	urlParams := strings.Join(urlpath, "&")
	body, _ := m.get_url(urlParams)

	var meetupEvents MeetupEvents
	err := json.Unmarshal(body, &meetupEvents)
	if err != nil {
		fmt.Println("Erro unmarshaling, Details:", err)
	}

	var events Events
	for _, v := range meetupEvents.Events {
		events = append(events, Event{
			name:        v.Name,
			description: v.Description[:len(v.Description)%15] + "...",
			link:        v.Link,
			location:    v.Venue.Name,
			dateTime:    v.LocalDate + "@" + v.LocalTime,
			organizer:   v.Group.Name,
		})
	}
	m.m.Events = events
	return m.m
}
