package shell

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sharpvik/vish/parser"
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

	astree, err := parser.New().Parse(input)
	if err != nil {
		PrintError(err)
		return
	}

	err = Eval(astree)
	Alert(err)
}

func Read() (input string, err error) {
	fmt.Print(Prefix())
	input, err = NewReader(os.Stdin).Next()
	return
}
