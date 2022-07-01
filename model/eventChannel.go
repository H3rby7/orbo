package model

import (
	"fmt"
	"time"
)

type EventChannel struct {
	EventType []string
	Location  string
	Date      time.Time
	Duration  string
	Customer  string
	Occasion  string
	Wishes    string
}

func NewEventChannel(data map[string]interface{}) (EventChannel, error) {
	e := EventChannel{}
	str, ok := data["occasion"].(string)
	if !ok {
		return e, createError("occasion", "string")
	}
	e.Occasion = str
	return e, nil
}

func createError(key string, expectedType string) error {
	return fmt.Errorf("field '%v' missing or not of type '%v'", key, expectedType)
}
