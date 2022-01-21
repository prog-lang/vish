package vish

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

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
	Alert(err)
	Print(result)
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
	split := strings.Fields(input)
	if len(split) == 0 {
		return
	}

	command := split[0]
	args := split[1:]

	if ran, err := RunSpecialCommand(command, args); ran {
		return nil, err
	}
	result, err = exec.Command(command, args...).Output()
	return
}

func Print(result []byte) {
	fmt.Println(string(result))
}
