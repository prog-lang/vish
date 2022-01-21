package vish

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Start() {
	for {
		REPL()
	}
}

func REPL() {
	input, err := Read()
	Abort(err)
	result, err := Eval(input)
	Abort(err)
	Print(result)
}

func Read() (input string, err error) {
	fmt.Print("VISH ❯❯❯ ")
	input, err = NewReader(os.Stdin).Next()
	return
}

func Eval(input string) (result []byte, err error) {
	split := strings.Fields(input)
	command := split[0]
	args := split[1:]

	if RunSpecialCommand(command, args) {
		return
	}

	result, err = exec.Command(command, args...).Output()

	return
}

func Abort(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func Print(result []byte) {
	fmt.Println(string(result))
}
