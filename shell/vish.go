package shell

import (
	"fmt"
	"os"
	"os/exec"
)

type Vish struct {
	sigChan chan os.Signal
	cmd     *exec.Cmd
}

func New() *Vish {
	return &Vish{
		sigChan: make(chan os.Signal, 1),
	}
}

func (vish *Vish) Start() {
	fmt.Println(ShortInfo())
	fmt.Println(Welcome)
	go vish.manageSignals()
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
