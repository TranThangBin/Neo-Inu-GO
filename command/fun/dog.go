package fun

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type dogResponse struct {
	Message string
	Status  string
}

var DogCommand = &discordgo.ApplicationCommand{
	Name:        "dog",
	Description: "Get a random dog image",
}

func DogHandler(s *discordgo.Session, e *discordgo.InteractionCreate) {
	res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Panicf("An error occured when trying to fetch dog image. Error: %v", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panicf("An error occured when trying to read the response body. Error: %v", err)
		return
	}
	var dog dogResponse
	err = json.Unmarshal(body, &dog)
	if err != nil {
		log.Panicf("An error occured when trying to parse the response body. Error: %v", err)
		return
	}
	s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Image: &discordgo.MessageEmbedImage{
						URL: dog.Message,
					},
				},
			},
		},
	})
}
