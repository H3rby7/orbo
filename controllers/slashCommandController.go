package controllers

import (
	"h3rby7/orbo/model"
	"h3rby7/orbo/util"
	"h3rby7/orbo/views"
	"log"
	"time"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

const (
	// Define Action_id as constant so we can refer to them in the controller
	InquiryCommand = "/anfrage"
)

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler   *socketmode.SocketmodeHandler
	ChannelHandler *ChannelHandler
}

func NewSlashCommandController(eventhandler *socketmode.SocketmodeHandler) SlashCommandController {
	// we need to cast our socketmode.Event into a SlashCommand
	c := SlashCommandController{
		EventHandler:   eventhandler,
		ChannelHandler: &ChannelHandler{},
	}

	c.EventHandler.HandleSlashCommand(
		InquiryCommand,
		c.inquiry,
	)

	// The form is sent back to us
	c.EventHandler.HandleInteractionBlockAction(
		views.InquiryIds["form_action"],
		c.handleInquiryForm,
	)

	// TEST commands
	c.EventHandler.HandleSlashCommand("/test", c.testRequest)

	return c

}

func (c SlashCommandController) inquiry(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	command, convErr := evt.Data.(slack.SlashCommand)

	if !convErr {
		log.Printf("ERROR converting event to Slash Command: %v", convErr)
	}

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	blocks := views.InquiryForm()

	// Post ephemeral message
	_, _, err := clt.PostMessage(
		command.ChannelID,
		slack.MsgOptionBlocks(blocks...),
		slack.MsgOptionResponseURL(command.ResponseURL, slack.ResponseTypeEphemeral),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for '%v': %v", InquiryCommand, err)
	}

}

func (c SlashCommandController) handleInquiryForm(evt *socketmode.Event, clt *socketmode.Client) {
	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)
	// TODO: This is the place to create channel, survey etc. - May want to add a microservice Survey App using Slack message metadata instead of a DB

	// we need to cast our socketmode.Event into a Slash Command
	interaction := evt.Data.(slack.InteractionCallback)
	rawInputs := interaction.BlockActionState.Values

	inputs := map[string]interface{}{}
	ids := util.MapReverse(views.InquiryIds)

	for _, entry := range rawInputs {
		for actionId, value := range entry {
			isMulti, single, multi := getValuesFromBlockAction(value)
			if isMulti {
				inputs[ids[actionId]] = multi
			} else {
				inputs[ids[actionId]] = single
			}
		}
	}

	eventData, conversionErr := model.NewEventChannel(inputs)

	if conversionErr != nil {
		log.Printf("ERROR while reading form data from FORM '%v' -> %v", InquiryCommand, conversionErr)
		// TODO: more Error handling?
	}
	// TODO: continue here with eventData -> create Channel, post some mesasgies.

	_, slackError := c.ChannelHandler.create(eventData, clt)

	if slackError != nil {
		log.Printf("ERROR creating channel from FORM '%v' -> %v", InquiryCommand, slackError)
		// TODO: more Error handling?
	}

}

func (c SlashCommandController) testRequest(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	command, convErr := evt.Data.(slack.SlashCommand)

	if !convErr {
		log.Printf("ERROR converting event to Slash Command: %v", convErr)
	}

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// Post message
	_, _, err := clt.PostMessage(
		command.ChannelID,
		slack.MsgOptionText("hello boyyyy, it is "+time.Now().Format(time.RFC3339), false),
		slack.MsgOptionUsername("jokahl"),
		slack.MsgOptionAsUser(true),
		slack.MsgOptionMetadata(slack.SlackMetadata{
			EventType: "aaa",
			EventPayload: map[string]interface{}{
				"test": "text",
			},
		}),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for 'test': %v", err)
	}

}
