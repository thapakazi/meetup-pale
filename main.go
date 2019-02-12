package pale

import (
	"fmt"
)

type Event struct {
	name        string
	link        string
	location    string
	dateTime    string
	description string
	organizer   string
}
type Events []Event

type Meetup struct {
	Events
	Tags []string
}

func NewMeetup() *Meetup {
	return &Meetup{
		Events: []Event{},
		Tags:   []string{},
	}
}

type meetups interface {
	fetch() Meetup
}

func fetchMeetup(m meetups) {
	_meetup := m.fetch()
	fmt.Println(_meetup)
}
