package dgoc

var CommandMap = make(map[string]interface{})

type Command interface {
	Execute(args []string)
}
