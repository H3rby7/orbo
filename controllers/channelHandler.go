package controllers

import (
	"fmt"
	"h3rby7/orbo/model"
	"strings"
	"time"

	"github.com/slack-go/slack/socketmode"
)

type ChannelHandler struct {
	EventHandler *socketmode.SocketmodeHandler
}

func (h ChannelHandler) create(ev model.EventChannel, clt *socketmode.Client) {
	channelName := fmt.Sprintf("%v_%v_%v", channelDatePart(ev.Date), ev.Location, channelTypeParts(ev.eventType))
	clt.CreateConversation(channelName, false)
}

func channelDatePart(date time.Time) string {
	return date.Format("2006_01_02")
}

func channelTypeParts(eventType []string) string {
	if len(eventType) == 0 {
		return "tbd"
	}
	return strings.Join(eventType, "_")
}
