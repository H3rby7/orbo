package controllers

import (
	"log"

	"github.com/slack-go/slack"
)

// Get Value of BlockAction
// Returns an array of strings containing all values
func getValuesFromBlockAction(blockAction slack.BlockAction) (bool, string, []string) {
	switch blockAction.Type {
	case "plain_text_input":
		return false, blockAction.Value, nil
	case "checkboxes":
		values := make([]string, len(blockAction.SelectedOptions))
		for i, v := range blockAction.SelectedOptions {
			values[i] = v.Value
		}
		return true, "", values
	case "datepicker":
		return false, blockAction.SelectedDate, nil
	case "timepicker":
		return false, blockAction.SelectedTime, nil
	}
	log.Printf("WARNING: Unknown blockAction.Type: %v", blockAction.Type)
	return true, "", nil
}
