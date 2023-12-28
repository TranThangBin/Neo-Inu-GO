package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	Commands        = []*discordgo.ApplicationCommand{pingCommand}
	commandHandlers = map[string]func(s *discordgo.Session, e *discordgo.InteractionCreate){
		pingCommand.Name: pingHandler,
	}
)

func OnCommand(s *discordgo.Session, e *discordgo.InteractionCreate) {
	commandName := e.ApplicationCommandData().Name
	if h, ok := commandHandlers[commandName]; ok {
		h(s, e)
	} else {
		s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Command %s doesn't have a handler", commandName),
			},
		})
	}
}
