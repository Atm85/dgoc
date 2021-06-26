package dgoc

type Command interface {
	Execute(args []string)
}
