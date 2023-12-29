package command

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"Neo-Inu/command/fun"
	"Neo-Inu/command/info"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		fun.PingCommand,
		fun.DogCommand,
		info.MemberCommand,
		info.ServerCommand,
	}
	commandHandlers = map[string]func(s *discordgo.Session, e *discordgo.InteractionCreate){
		fun.PingCommand.Name:    fun.PingHandler,
		fun.DogCommand.Name:     fun.DogHandler,
		info.MemberCommand.Name: info.MemberHandler,
		info.ServerCommand.Name: info.ServerHandler,
	}
)

func OnCommand(s *discordgo.Session, e *discordgo.InteractionCreate) {
	commandName := e.ApplicationCommandData().Name
	log.Printf("%s use /%s", e.Member.User.Username, commandName)
	if h, ok := commandHandlers[commandName]; ok {
		h(s, e)
	}
}
