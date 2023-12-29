package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"Neo-Inu/command"
	"Neo-Inu/event"
)

var s *discordgo.Session

var (
	guildId        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	botToken       = flag.String("token", "", "Token of your discord bot application")
	removeCommands = flag.Bool("rmcmd", true, "Remove all command after shutdown")
)

func init() {
	flag.Parse()
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + *botToken)
	if err != nil {
		log.Fatalf("Invalid bot parameter %s", err.Error())
	}
}

func init() {
	s.Identify.Intents = discordgo.IntentGuildMessages
}

func init() {
	s.AddHandler(event.OnReady)
	s.AddHandler(event.OnMessage)
	s.AddHandler(command.OnCommand)
}

func main() {
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session %s", err.Error())
	}
	registeredCommands := make([]*discordgo.ApplicationCommand, len(command.Commands))
	for i, command := range command.Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *guildId, command)
		if err != nil {
			log.Panicf("Unable to create command %s: %v", command.Name, err)
		}
		log.Printf("Successfully load command %s", cmd.Name)
		registeredCommands[i] = cmd
	}
	defer s.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGINT, syscall.SIGTERM,
		os.Interrupt, os.Kill)
	log.Println("Press Ctrl+C to exit")
	<-stop
	if *removeCommands {
		for _, command := range registeredCommands {
			err = s.ApplicationCommandDelete(s.State.User.ID, *guildId, command.ID)
			if err != nil {
				log.Panicf("Unable to delete command %s: %v", command.Name, err)
			}
			log.Printf("Successfully delete command %s", command.Name)
		}
	}
	log.Println("Gracefully shutting down.")
}
