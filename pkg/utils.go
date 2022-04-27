package deus_cc

import (
	"log"
	"net/url"
	"sync"

	"github.com/google/uuid"
)

type Storer struct {
	events  map[Event]bool
	counter map[string]int
}

var singleton *Storer
var once sync.Once

func GetStorer() *Storer {
	once.Do(func() {
		singleton = &Storer{
			events:  make(map[Event]bool),
			counter: make(map[string]int),
		}
	})
	return singleton
}

// SetTestEvents is used to "mock" the event collector. Used for testing only
func SetTestEvents(events []Event) {
	for _, event := range events {
		Adder(event)
	}
}

// Adder is the responsible to add (or not) an event to the event collector
// Use an Event struct as an input and return a boolean
func Adder(event Event) bool {
	storer := GetStorer()
	if !storer.events[event] {
		storer.events[event] = true
		storer.counter[event.Url]++
		return true
	}
	return false
}

// Getter is the responsible to return the number of distinct visitors relate
// with the url provided as input param
// Use an string as an input and return a int
func Getter(url string) int {
	storer := GetStorer()
	log.Println(url)
	log.Println(storer.counter[url])
	return storer.counter[url]
	// log.Println(storer.events)
	// log.Println(storer.counter[url])
	// for ev := range storer.events {
	// 	if ev.Url == url {
	// 		counter++
	// 	}
	// }
	// return counter
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
