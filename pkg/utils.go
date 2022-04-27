package deus_cc

import (
	"net/url"

	"github.com/google/uuid"
)

var events = make(map[Event]bool)

// SetTestEvents is used to "mock" the event collector. Used for testing only
func SetTestEvents(input map[Event]bool) {
	events = input
}

// Adder is the responsible to add (or not) an event to the event collector
// Use an Event struct as an input and return a boolean
func Adder(event Event) bool {
	if !events[event] {
		events[event] = true
		return true
	}
	return false
}

// Getter is the responsible to return the number of distinct visitors relate
// with the url provided as input param
// Use an string as an input and return a int
func Getter(url string) int {
	counter := 0
	for ev := range events {
		if ev.Url == url {
			counter++
		}
	}
	return counter
}

func ValidateData(event Event) (bool, error) {
	_, err := url.ParseRequestURI(event.Url)
	if err != nil {
		return false, err
	}

	_, err = uuid.Parse(event.Uuid)
	if err != nil {
		return false, err
	}
	return true, nil
}
