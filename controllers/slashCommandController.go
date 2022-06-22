package controllers

import (
	"h3rby7/orbo/views"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

const (
	// Define Action_id as constant so we can refer to them in the controller
	ShowRequestCommand = "/anfrage"
)

// We create a sctucture to let us use dependency injection
type SlashCommandController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewSlashCommandController(eventhandler *socketmode.SocketmodeHandler) SlashCommandController {
	// we need to cast our socketmode.Event into a SlashCommand
	c := SlashCommandController{
		EventHandler: eventhandler,
	}

	c.EventHandler.HandleSlashCommand(
		ShowRequestCommand,
		c.showRequest,
	)

	// The form is sent back to us
	c.EventHandler.HandleInteractionBlockAction(
		views.IDs["form_action"],
		c.handleShowRequestForm,
	)

	return c

}

func (c SlashCommandController) showRequest(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	command, convErr := evt.Data.(slack.SlashCommand)

	if convErr != true {
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
		log.Printf("ERROR while sending message for '%v': %v", ShowRequestCommand, err)
	}

}

func (c SlashCommandController) handleShowRequestForm(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into a Slash Command
	interaction := evt.Data.(slack.InteractionCallback)
	rawInputs := interaction.BlockActionState.Values

	inputs := map[string]interface{}{}

	for _, entry := range rawInputs {
		for actionId, value := range entry {
			inputs[actionId] = getValuesFromBlockAction(value)
		}
	}

	// Make sure to respond to the server to avoid an error
	clt.Ack(*evt.Request)

	// create the view using block-kit
	// blocks := views.LaunchRocket()

	_, _, err := clt.PostMessage(
		interaction.Container.ChannelID,
		slack.MsgOptionBlocks(),
		slack.MsgOptionResponseURL(interaction.ResponseURL, slack.ResponseTypeInChannel),
		slack.MsgOptionReplaceOriginal(interaction.ResponseURL),
	)

	// Handle errors
	if err != nil {
		log.Printf("ERROR while sending message for /rocket: %v", err)
	}

}
