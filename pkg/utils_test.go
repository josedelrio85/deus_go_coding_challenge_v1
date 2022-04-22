package deus_cc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	assert := assert.New(t)

	events := map[Event]bool{
		{
			Url:  "http://web1.com/foo",
			Uuid: "4e9ca7cc-2e11-4b8d-adc4-c77e6606dbff",
		}: true,
		{
			Url:  "http://web1.com/bar",
			Uuid: "653dfcae-67f3-4d73-89b6-8d0dc94781e4",
		}: true,
		{
			Url:  "http://web1.com/foo",
			Uuid: "eda0c9fd-a5a0-48f3-96d6-7c0c8b3b8f8b",
		}: true,
	}
	SetTestEvents(events)

	tests := []struct {
		Description    string
		Event          Event
		ExpectedResult bool
	}{
		{
			Description: "when an event is not present, add it to events collector",
			Event: Event{
				Url:  "http://web1.com/new",
				Uuid: "eda0c9fd-2e11-4d73-adc4-8d0dc94781e4",
			},
			ExpectedResult: true,
		},
		{
			Description: "when an event is present, take no action",
			Event: Event{
				Url:  "http://web1.com/foo",
				Uuid: "4e9ca7cc-2e11-4b8d-adc4-c77e6606dbff",
			},
			ExpectedResult: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			result := Adder(test.Event)
			assert.Equal(test.ExpectedResult, result)
		})
	}
}

func TestGetter(t *testing.T) {
	assert := assert.New(t)

	events := map[Event]bool{
		{
			Url:  "http://web1.com/foo",
			Uuid: "4e9ca7cc-2e11-4b8d-adc4-c77e6606dbff",
		}: true,
		{
			Url:  "http://web1.com/bar",
			Uuid: "653dfcae-67f3-4d73-89b6-8d0dc94781e4",
		}: true,
		{
			Url:  "http://web1.com/foo",
			Uuid: "eda0c9fd-a5a0-48f3-96d6-7c0c8b3b8f8b",
		}: true,
	}
	SetTestEvents(events)

	tests := []struct {
		Description    string
		Url            string
		ExpectedResult bool
		ExpectedValue  int
	}{
		{
			Description:    "when an url has no visitors, return 0 value",
			Url:            "http://web1.com/new",
			ExpectedResult: false,
			ExpectedValue:  0,
		},
		{
			Description:    "when an url has visitors, return an integer greater than 0",
			Url:            "http://web1.com/foo",
			ExpectedResult: true,
			ExpectedValue:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			result := Getter(test.Url)
			if test.ExpectedResult {
				assert.Greater(result, 0)
			}
			assert.Equal(test.ExpectedValue, result)
		})
	}
}
