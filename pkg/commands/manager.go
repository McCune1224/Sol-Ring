package commands

import (
	"log"

	dg "github.com/bwmarrin/discordgo"
)

// Wrapper Structure to couple A ApplicationCommand with a dedicated Handler
type SlashCommand struct {
	// to avoid naming stuttering (SlashCommand.Command....) discordgo.ApplicationCommand is aliased to "Feature"
	Feature dg.ApplicationCommand

	// Slash Command Handler aka how the command should operate
	Handler func(s *dg.Session, i *dg.InteractionCreate)
}

// Slash Command Manager, binds all commands to a map, the command name now is used to reference a SlashCommand
type SlashCommandManager struct {
	SlashCommands map[string]SlashCommand
}

// Create New Instance of a Slash Command manager
func NewSlashCommandManager() *SlashCommandManager {
	scm := &SlashCommandManager{
		SlashCommands: make(map[string]SlashCommand),
	}
	return scm
}

func (scm SlashCommandManager) AddCommand(sc SlashCommand) {
	scm.SlashCommands[sc.Feature.Name] = sc
}

// Returns a tally of successful applications created for the session
func (scm SlashCommandManager) RegisterCommands(session *dg.Session) int {
	tally := 0
	for _, slashcmd := range scm.SlashCommands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, "", &slashcmd.Feature)
		if err != nil {
			log.Printf("Failed to add command %s\nError: %s", slashcmd.Feature.Name, err.Error())
			continue
		}
		tally++
	}
	return tally
}
