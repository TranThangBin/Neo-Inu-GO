package info

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var ServerCommand = &discordgo.ApplicationCommand{
	Name:        "server",
	Description: "Provide information about the server",
}

func ServerHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	guild, err := s.State.Guild(e.GuildID)
	if err != nil {
		log.Panicf("An error occured when trying to get guild info")
		return
	}
	guildId, err := strconv.ParseInt(guild.ID, 10, 64)
	if err != nil {
		log.Panicf("An error occured when trying to parse guild ID")
		return
	}
	createdAt := time.Unix((guildId>>22+int64(1420070400000))/1000, 0).Local()
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("`%s` is created at `%v` and has `%d members`",
				guild.Name,
				createdAt,
				guild.MemberCount,
			),
			Embeds: []*discordgo.MessageEmbed{
				{
					Image: &discordgo.MessageEmbedImage{
						URL: guild.IconURL(""),
					},
				},
			},
		},
	})
}
