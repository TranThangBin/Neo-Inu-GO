package command

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		pingCommand,
		memberCommand,
	}
	commandHandlers = map[string]func(s *discordgo.Session, e *discordgo.InteractionCreate){
		pingCommand.Name:   pingHandler,
		memberCommand.Name: memberHandler,
	}
)

func OnCommand(s *discordgo.Session, e *discordgo.InteractionCreate) {
	commandName := e.ApplicationCommandData().Name
	if h, ok := commandHandlers[commandName]; ok {
		h(s, e)
	}
}
