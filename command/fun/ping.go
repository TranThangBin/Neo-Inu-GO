package fun

import (
	"github.com/bwmarrin/discordgo"
)

var PingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Response with pong",
}

func PingHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	s.InteractionRespond(e.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "pong",
			},
		})
}
