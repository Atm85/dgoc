package dgoc

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	userid      string
	prefix      string
	defaultHelp bool = true
)

func SetPrefix(p string) {
	prefix = p
}

func UseDefaultHelp(value bool) {
	defaultHelp = value
}

type DGOC struct {
	Session    *discordgo.Session
	CommandMap map[string]*ICommand
}

// New initializes the dgoc command handler
func New(session *discordgo.Session) *DGOC {

	bot, _ := session.User("@me")
	userid = bot.ID

	session.AddHandler(OnMessage)
	return &DGOC{Session: session}
}

// AddCommand add a command(s) to the dgoc command map
func (dg *DGOC) AddCommand(commands ...interface{}) error {

	for _, command := range commands {

		i := reflect.TypeOf((*ICommand)(nil)).Elem()
		t := reflect.TypeOf(command)

		if t.Kind() != reflect.Ptr {
			return fmt.Errorf("error: command %v is not a pointer", t)
		}

		if !t.Implements(i) {
			return fmt.Errorf("error: command %v does not implement %v", t, i)
		}

		// t.String = *package.Name
		//
		// here we split 't.String()'
		// and extract the command Name
		// and convert it to lowercase
		name := strings.ToLower(strings.Split(t.String(), ".")[1])

		// get command description from ICommand interface object
		command.(ICommand).Prepare()

		v := reflect.ValueOf(command)
		f := v.Elem().FieldByName("Desc")
		desc := ""
		if f.IsValid() {
			desc = f.String()
		}

		CommandMap[name] = &Command{
			Name:     name,
			Desc:     desc,
			Executor: command,
		}
	}

	return nil
}
