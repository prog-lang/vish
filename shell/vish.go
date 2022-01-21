package vish

import (
	"fmt"
	"os"
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
	err = Eval(input)
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

func Eval(input string) (err error) {
	command := ParseCommand(input)
	if command == nil {
		return
	}
	if ran, err := RunSpecialCommand(command); ran {
		return err
	}
	ExecCommand(command)
	return
}
