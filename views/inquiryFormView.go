package views

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/slack-go/slack"
)

var InquiryIds = map[string]string{
	"form_action": "inquiry_form_action",
	"type":        "inquiry_type",
	"location":    "inquiry_location",
	"date":        "inquiry_date",
	"time":        "inquiry_time",
	"duration":    "inquiry_duration",
	"customer":    "inquiry_customer",
	"occasion":    "inquiry_occasion",
	"wishes":      "inquiry_wishes",
}

//go:embed slackCommandAssets/*
var slashCommandAssets embed.FS

func InquiryForm() []slack.Block {

	tpl := renderTemplate(slashCommandAssets, "slackCommandAssets/inquiryForm.json", InquiryIds)

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}
