package info

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var MemberCommand = &discordgo.ApplicationCommand{
	Name:        "member",
	Description: "Get info of a member of the server",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "username",
			Description: "Username of the member",
			Required:    true,
		},
	},
}

func MemberHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	userId := e.ApplicationCommandData().Options[0].Value.(string)
	guild, err := s.State.Guild(e.GuildID)
	if err != nil {
		log.Panicf("Error getting guild info %v", err)
		return
	}
	member, err := s.GuildMember(guild.ID, userId)
	if err != nil {
		log.Panicf("Error getting member info %v", err)
		return
	}
	s.InteractionRespond(e.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
				Content: fmt.
					Sprintf("`%s` joined `%s` at `%v`",
						member.User.Username,
						guild.Name,
						member.JoinedAt.Local(),
					),
				Embeds: []*discordgo.MessageEmbed{
					{
						Image: &discordgo.MessageEmbedImage{
							URL: member.AvatarURL(""),
						},
					},
				},
			},
		})
}
