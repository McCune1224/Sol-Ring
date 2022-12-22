package main

import (
	"log"
	"os"
	"os/signal"

	dg "github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/Sol-Ring/pkg/commands"
)

// Create Bot Session
var (
	discordBot *dg.Session
)

func init() {
	var err error
	token := os.Getenv("DISCORD_TOKEN")
	discordBot, err = dg.New("Bot " + token)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Create Command Handler and Register Slash Commands
var SCM *commands.SlashCommandManager

func init() {
	SCM = commands.NewSlashCommandManager()
	SCM.AddCommand(commands.Ping)
	SCM.AddCommand(commands.CardLookup)
	// Create handlers for interactionCreation Commands (Message & Componetns)
	discordBot.AddHandler(func(s *dg.Session, i *dg.InteractionCreate) {
        switch i.Type{
        case dg.InteractionApplicationCommand:
    }

		if cmd, ok := SCM.SlashCommands[i.ApplicationCommandData().Name]; ok {
			cmd.Handler(s, i)
		}

	})
}

// Create Websocket Connection, Register Commands, and run until signal close
func main() {
	var err error

	err = discordBot.Open()
	if err != nil {
		log.Fatal(err.Error())
	}
	regCommandTally := SCM.RegisterCommands(discordBot)
	log.Printf("Successfully registered %d/%d commands", len(SCM.SlashCommands), regCommandTally)
	defer discordBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Printf("%s Listening for Commands", discordBot.State.User.Username)
	<-stop
}
