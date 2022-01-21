package vish

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

func RunSpecialCommand(
	command string,
	args []string,
) (special bool, err error) {
	cmd, special := commands[command]
	if special {
		err = cmd(args)
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
	fmt.Println("Bye! :)")
	os.Exit(0)
	return
}
