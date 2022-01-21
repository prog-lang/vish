package vish

import (
	"fmt"
	"os"
)

func Start() {
	fmt.Println(ShortInfo())
	fmt.Println(Welcome)
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
	fmt.Print(Prefix())
	input, err = NewReader(os.Stdin).Next()
	return
}

func Eval(input string) (err error) {
	command := ParseCommand(input)
	if command == nil {
		return
	}
	if ran, err := RunSpecialCommand(command); ran {
		return err
	}
	return ExecCommand(command)
}
