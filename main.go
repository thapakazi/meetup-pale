package pale

import (
	"fmt"
	"time"
)

type Event struct {
	name     string
	location string
	date     time.Duration
	details  string
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
