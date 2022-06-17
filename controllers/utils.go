package controllers

import (
	"github.com/slack-go/slack"
)

// Get Value of BlockAction
// Returns an array of strings containing all values
func getValuesFromBlockAction(blockAction slack.BlockAction) []string {
	var values []string
	switch blockAction.Type {
	case "plain_text_input":
		values = []string{blockAction.Value}
	case "checkboxes":
		values := make([]string, len(blockAction.SelectedOptions))
		for i, v := range blockAction.SelectedOptions {
			values[i] = v.Value
		}
	case "datepicker":
		values = []string{blockAction.SelectedDate}
	case "timepicker":
		values = []string{blockAction.SelectedTime}
	}
	return values
}
