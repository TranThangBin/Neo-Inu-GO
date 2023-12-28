package event

import (
	"github.com/bwmarrin/discordgo"
)

func OnMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.ID == s.State.User.ID {
		return
	}
	if e.Content == "ping" {
		s.ChannelMessageSend(e.ChannelID, "pong")
	}
}
