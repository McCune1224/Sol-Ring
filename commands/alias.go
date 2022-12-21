package commands

import dg "github.com/bwmarrin/discordgo"

func SendChannelMessage(s *dg.Session, i *dg.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionResponseData{
			Content: content,
		},
	})
}
