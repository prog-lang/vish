package shell

import (
	"fmt"
	"os"
	"path"
)

var commands = map[string]func([]string) error{
	"cd":      cd,
	"version": version,
	"exit":    exit,
}

func RunSpecialCommand(cmd *Command) (special bool, err error) {
	command, special := commands[cmd.Command]
	if special {
		err = command(cmd.Args)
	}
	return
}

func cd(args []string) (err error) {
	if len(args) == 0 {
		return changeWdToHomeDir()
	}
	target := args[0]
	return os.Chdir(path.Clean(target))
}

func version(_ []string) (err error) {
	fmt.Println(ShortInfo())
	return
}

func exit(args []string) (err error) {
	fmt.Println(GoodBye)
	os.Exit(0)
	return
}
