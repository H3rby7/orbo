package model

import "testing"

func createTestInput() map[string]interface{} {
	testInput := map[string]interface{}{}
	testInput["type"] = []string{"show"}
	testInput["location"] = "a stage"
	testInput["date"] = "1990-04-28"
	testInput["time"] = "12:03"
	testInput["duration"] = "until midnight"
	testInput["customer"] = "a friend"
	testInput["occasion"] = "teamevent"
	testInput["wishes"] = "none"
	return testInput
}

func TestNewEventChannel(t *testing.T) {
	res, err := NewEventChannel(createTestInput())
	if err != nil {
		t.Error(err)
	}
	if res.EventType[0] != "show" {
		t.Error("Expected EventType to be 'show'")
	}
	if res.Location != "a stage" {
		t.Error("Expected Location to be 'a stage'")
	}
	if res.Duration != "until midnight" {
		t.Error("Expected Duration to be 'until midnight'")
	}
	if res.Customer != "a friend" {
		t.Error("Expected Customer to be 'a friend'")
	}
	if res.Occasion != "teamevent" {
		t.Error("Expected Occasion to be 'teamevent'")
	}
	if res.Wishes != "none" {
		t.Error("Expected Wishes to be 'none'")
	}
	// Date Tests
	y, m, d := res.Date.Date()
	if y != 1990 {
		t.Errorf("Wrong year, expected 1990, but got '%v'", y)
	}
	if m != 4 {
		t.Errorf("Wrong month, expected 4, but got '%v'", m)
	}
	if d != 28 {
		t.Errorf("Wrong day, expected 28, but got '%v'", d)
	}
	h, min, _ := res.Date.Clock()
	if h != 12 {
		t.Errorf("Wrong hour, expected 12, but got '%v'", h)
	}
	if min != 3 {
		t.Errorf("Wrong minute, expected 3, but got '%v'", min)
	}
}

func TestNewEventChannelMultipleTypes(t *testing.T) {
	input := createTestInput()
	input["type"] = []string{"show", "workshop"}
	res, err := NewEventChannel(input)
	if err != nil {
		t.Error(err)
	}
	if len(res.EventType) != 2 {
		t.Error("Expected EventType to be of length '2'")
	}
	if res.EventType[0] != "show" {
		t.Error("Expected EventType[0] to be 'show'")
	}
	if res.EventType[1] != "workshop" {
		t.Error("Expected EventType[1] to be 'workshop'")
	}
}
