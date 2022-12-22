package commands

import (
	"fmt"
	"log"

	dg "github.com/bwmarrin/discordgo"
	mtg "github.com/mccune1224/Sol-Ring/pkg/mtg"
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

		cardResponse, err := mtg.CardSearch(cardName)
		if err != nil {
			log.Fatal(err)
		}

		if len(*cardResponse) == 0 {
			SendChannelMessage(s, i, fmt.Sprintf("No Cards Found for Query '%s'", cardName))
			return
		}
		bestMatch := mtg.Card{}
		for _, item := range *cardResponse {
			bestMatch = item
			break
		}
		for _, item := range *cardResponse {
			if item.Name == cardName {
				bestMatch = item
			}
		}

		embed := &dg.MessageEmbed{
			Title: bestMatch.Name,
			URL:   bestMatch.ScryfallURI,
			// Description: fmt.Sprintf("%s %s", bestMatch.Rarity, bestMatch.ReleasedAt),
			Image: &dg.MessageEmbedImage{
				URL: bestMatch.ImageUris.Normal,
			},
			Fields: []*dg.MessageEmbedField{
				{
					Name:  "Price",
					Value: fmt.Sprintf(" [$%s - Tcgplayer](%s)", bestMatch.Prices.Usd, bestMatch.PurchaseUris.Tcgplayer),
				},
			},

			Footer: &dg.MessageEmbedFooter{},
		}

		// SendChannelMessage(s, i, cardDetails)
		SendChannelEmbed(s, i, embed)
	},
}
