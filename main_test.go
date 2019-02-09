package pale

import (
	"testing"
)

func TestRun(t *testing.T) {

	twitterM := NewMeetupT()
	meetupDotComM := NewMeetupM()
	ms := []meetups{twitterM, meetupDotComM}
	for _, m := range ms {
		fetchMeetup(m)
	}
}
