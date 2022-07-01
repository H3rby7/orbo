package model

import "testing"

func createTestInput() map[string]interface{} {
	testInput := map[string]interface{}{}
	testInput["occasion"] = "teamevent"
	return testInput
}

func TestNewEventChannel(t *testing.T) {
	_, err := NewEventChannel(createTestInput())
	if err != nil {
		t.Error(err)
	}
}
