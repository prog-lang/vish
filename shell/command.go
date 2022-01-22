package shell

type Command struct {
	Command string
	Args    []string
}

func NewCommand(cmd string, args []string) *Command {
	return &Command{
		Command: cmd,
		Args:    args,
	}
}
