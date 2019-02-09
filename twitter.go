package pale

import "fmt"

type MeetupT struct {
	Meetup
	TwitterStuffs
}

type TwitterStuffs struct {
}

func NewMeetupT() *MeetupT {
	return &MeetupT{
		Meetup: *NewMeetup(),
	}
}

func (m *MeetupT) fetch() Meetup {
	fmt.Println("pulling tweets from twitter")
	return m.Meetup
}
