package cmn4j_test

import (
	"testing"

	"time"

	"github.com/kc1116/cmool-events/events"
)

var db = events.Db

func TestCreateEventNode(t *testing.T) {
	var testEvent events.Event
	testEvent.Properties.Name = "Test"
	testEvent.Properties.DateCreated = time.Now()
	testEvent.Properties.Description = "This is a test event."
	testEvent.Properties.Keywords = []string{"key", "words"}
	testEvent.Properties.Rating = 3.5
	testEvent.Properties.TypeOfEvent = "Just a Test"
	testEvent.Properties.Emblem = "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRjWFifmZY2WER6nCFNZOCtF2WSRm2vkDr3erTHUdTWFI8tCoQDJaXNJ5c"
	testEvent.Properties.Location.City = "Harrison"
	testEvent.Properties.Location.State = "New Jersey"
	testEvent.Properties.Location.StreetAddress = "1 Harrison ave"
	testEvent.Properties.Location.ZipCode = "07029"

	event, err := events.CreateEventNode(testEvent)
	if err != nil {
		t.Error("Expected an test event got an error:", err.Error())
	} else {
		t.Logf("TestCreateEventNode:%+v\n", event.Properties.UniqueID)
	}

}

func TestGetEventNode(t *testing.T) {
	uuid := "e17adaa4-a6d5-47c3-a8ee-13a59f4957d7"

	event, err := events.GetEventNode(uuid)
	if err != nil {
		t.Error("Expected an test event got an error:", err.Error())
	} else {
		t.Logf("TestGetEventNode:%+v\n", event.Properties)
	}

}
