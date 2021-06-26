package dgoc

import (
	"fmt"
	"reflect"

	"github.com/bwmarrin/discordgo"
)

type DGOC struct {
	Session *discordgo.Session
}

// New initializes the dgoc command handler
func New(session *discordgo.Session) *DGOC {
	return &DGOC{Session: session}
}

// AddCommand add a command(s) to the dgoc command map
func (dg *DGOC) AddCommand(commands ...interface{}) error {
	for _, command := range commands {
		i := reflect.TypeOf((*Command)(nil)).Elem()
		t := reflect.TypeOf(command)
		value := t.Implements(i)
		if !value {
			return fmt.Errorf("error: command %s does not implement %s", t, i)
		}
	}

	return nil
}
