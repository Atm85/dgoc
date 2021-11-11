package dgoc

var CommandMap = make(map[string]*Command)

type ICommand interface {
	Prepare()
	Execute(ctx *Context, args []string)
}

type Command struct {
	Name string
	Desc string

	Executor interface{}
}
