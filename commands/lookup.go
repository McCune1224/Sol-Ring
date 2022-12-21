package commands

import (
	"encoding/json"
	"log"

	dg "github.com/bwmarrin/discordgo"
	mtg "github.com/mccune1224/Sol-Ring/mtg"
)

var CardLookup = SlashCommand{
	Feature: dg.ApplicationCommand{
		Name:        "card",
		Description: "Queries for specified MTG Card",
		Options: []*dg.ApplicationCommandOption{
			{
				Type:        dg.ApplicationCommandOptionString,
				Name:        "card_name",
				Description: "Desired Card to see",
				Required:    true,
			},
		},
	},
	Handler: func(s *dg.Session, i *dg.InteractionCreate) {
		cardName := i.ApplicationCommandData().Options[0].StringValue()

		cardResponse, err := mtg.CardRequest(cardName)
		if err != nil {
			log.Fatal(err)
		}

		foo, err := json.Marshal(cardResponse)
		if err != nil {
			SendChannelMessage(s, i, "Error Marshalling Card Response")
		}

		SendChannelMessage(s, i, string(foo))
	},
}
