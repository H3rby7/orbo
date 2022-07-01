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
	if res.Occasion != "teamevent" {
		t.Error("Expected Occasion to be 'teamevent'")
	}
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
