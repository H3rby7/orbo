package controllers

import "github.com/slack-go/slack/socketmode"

type InviteHandler struct {
	EventHandler *socketmode.SocketmodeHandler
}

func (h InviteHandler) inviteAllMembers(clt *socketmode.Client) {

}
