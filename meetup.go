package pale

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
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
		"key":     os.Getenv("MEETUP_API_KEY"),
		"country": os.Getenv("MEETUP_API_COUNTRY"),
		"radius":  os.Getenv("MEETUP_API_RADIUS"),
	}
	urlpath := []string{"/find/groups?upcoming_events=true"}
	for k, v := range params {
		param := k + "=" + v
		urlpath = append(urlpath, param)
	}
	urlParams := strings.Join(urlpath, "&")
	body, _ := m.get_url(urlParams)

	fmt.Println(gjson.Get(string(body), "#.name"))
	events = append(events, Event{
		name: ee["name"].(string), // .(map[string]string)["name"].(string)
		// description: ee["description"].(string),
	})
	// 	// group_name := body["group_name"]
	// 	// group_description := body["group_description"]
	// 	// group_link := body["group_link"]
	// 	// lat := body["lat"]
	// 	// lon := body["lon"]
	// 	// members := body["members"]
	// 	// next_event_id := body["next_event_id"]
	// 	// next_event_name := body["next_event_name"]
	// 	// next_event_time := body["next_event_time "]
	// 	// organizer_name := body["organizer_name"]
	// }

	return m.m
}
