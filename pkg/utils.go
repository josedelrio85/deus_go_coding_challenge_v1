package deus_cc

import (
	"net/url"
	"sync"

	"github.com/google/uuid"
)

type Storer struct {
	events  map[Event]bool
	counter map[string]int
	mutex   *sync.Mutex
}

var singleton *Storer
var once sync.Once

func GetStorer() *Storer {
	once.Do(func() {
		singleton = &Storer{
			events:  make(map[Event]bool),
			counter: make(map[string]int),
			mutex:   &sync.Mutex{},
		}
	})
	return singleton
}

// SetTestEvents is used to "mock" the event collector. Used for testing only
func (s *Storer) SetTestEvents(events []Event) {
	for _, event := range events {
		s.Adder(event)
	}
}

// Adder is the responsible to add (or not) an event to the event collector
// Use an Event struct as an input and return a boolean
func (s *Storer) Adder(event Event) bool {
	if !s.events[event] {
		s.mutex.Lock()
		s.events[event] = true
		s.counter[event.Url]++
		s.mutex.Unlock()
		return true
	}
	return false
}

// Getter is the responsible to return the number of distinct visitors relate
// with the url provided as input param
// Use an string as an input and return a int
func (s *Storer) Getter(url string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.counter[url]
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
