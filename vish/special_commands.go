package vish

import (
	"fmt"
	"os"
)

var commands = map[string]func(){
	"exit": exit,
}

func RunSpecialCommand(command string, args []string) (special bool) {
	cmd, special := commands[command]
	if special {
		cmd()
	}
	return
}

func exit() {
	fmt.Println("Bye! :)")
	os.Exit(0)
}
