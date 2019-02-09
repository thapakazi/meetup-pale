package pale

import (
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestRun(t *testing.T) {

	// twitterM := NewMeetupT()
	meetupDotComM := NewMeetupM()
	ms := []meetups{meetupDotComM}
	for _, m := range ms {
		fetchMeetup(m)
	}
}
