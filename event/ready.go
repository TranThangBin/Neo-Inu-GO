package event

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, e *discordgo.Ready) {
	log.Printf("Logged in as: %s#%s",
		s.State.User.Username,
		s.State.User.Discriminator)
}
