package command

import (
	"github.com/bwmarrin/discordgo"
)

var pingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Response with pong",
}

func pingHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
}
