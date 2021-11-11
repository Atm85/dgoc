package dgoc

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// do nothing if the bot sends a message
	if m.Author.ID == userid {
		return
	}

	// check if command identifier is present in message
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// separate message into args
	args := strings.Split(m.Content, " ")
	name := strings.Split(args[0], ".")[1]

	// get command from map
	command, ok := CommandMap[name]
	if !ok {
		return
	}

	// set command context
	ctx := &Context{

		Name: name,

		Session: s,
		Message: m.Message,
	}

	// run command
	command.Executor.(ICommand).Execute(ctx, append(args[:0], args[0+1:]...))
}
