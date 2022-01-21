package vish

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/logrusorgru/aurora"
)

func Start() {
	for {
		REPL()
	}
}

func REPL() {
	input, err := Read()
	Alert(err)
	result, err := Eval(input)
	Print(result)
	Alert(err)
}

func Read() (input string, err error) {
	PrintPrefix()
	input, err = NewReader(os.Stdin).Next()
	return
}

func PrintPrefix() {
	cwd, err := os.Getwd()
	Alert(err)
	fmt.Printf("%s %s ",
		aurora.Bold(aurora.Green(path.Base(cwd))),
		aurora.Green("➜"))
}

func Eval(input string) (result []byte, err error) {
	command := ParseCommand(input)
	if command == nil {
		return
	}
	if ran, err := RunSpecialCommand(command); ran {
		return nil, err
	}
	result, err = exec.Command(
		command.Command, command.Args...).CombinedOutput()
	return
}

func Print(result []byte) {
	fmt.Println(string(result))
}
