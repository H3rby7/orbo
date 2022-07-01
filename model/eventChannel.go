package model

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type EventChannel struct {
	EventType []string `mapstructure:"type"`
	Location  string
	Date      time.Time `mapstructure:"-"`
	Duration  string
	Customer  string
	Occasion  string
	Wishes    string
}

func NewEventChannel(data map[string]interface{}) (EventChannel, error) {
	e := EventChannel{}
	err := mapstructure.Decode(data, &e)
	if err != nil {
		return e, err
	}
	dayStr, ok := data["date"].(string)
	if !ok {
		return e, errors.New("Invalid 'date'")
	}
	timeStr, ok := data["time"].(string)
	if !ok {
		return e, errors.New("Invalid 'time'")
	}
	date, err := time.Parse("2006-01-02T15:04", fmt.Sprintf("%vT%v", dayStr, timeStr))
	if err != nil {
		return e, err
	}
	e.Date = date
	return e, nil
}
