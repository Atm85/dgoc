package dgoc

import "github.com/bwmarrin/discordgo"

type DGOC struct {
	Session *discordgo.Session
}

// New initializes the dgoc command handler
func New(session *discordgo.Session) *DGOC {
	return &DGOC{Session: session}
}

// AddCommand add a command(s) to the dgoc command map
func (dg *DGOC) AddCommand(commands ...interface{}) {
	for _, command := range commands {
		_ = command
	}
}
