package pale

import "fmt"

type MeetupM struct {
	Meetup
	MeetupDotComStuffs
}

type MeetupDotComStuffs struct {
}

func NewMeetupM() *MeetupM {
	return &MeetupM{
		Meetup: *NewMeetup(),
	}
}

func (m *MeetupM) fetch() Meetup {
	fmt.Println("pulling tweets from Meetup.com")
	return m.Meetup
}
