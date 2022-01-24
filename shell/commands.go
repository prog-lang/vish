package shell

import (
	"fmt"
	"os"
	"path"

	"github.com/sharpvik/vish/ast"
)

var commands = map[string]func([]string) error{
	"cd":      cd,
	"version": version,
	"exit":    exit,
}

func RunSpecialCommand(cmd *ast.Command) (special bool, err error) {
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

// *** HELPERS ***

func changeWdToHomeDir() (err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	return os.Chdir(home)
}
