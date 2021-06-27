package dgoc

import (
	"reflect"
	"strings"

	"github.com/bwmarrin/discordgo"
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

	v := reflect.ValueOf(command)
	context := v.Elem().FieldByName("Ctx")
	if context.IsValid() {
		ctx := &Context{}
		context.Set(reflect.ValueOf(ctx))
	}

	// run command
	command.(Command).Execute(append(args[:0], args[0+1:]...))
}
