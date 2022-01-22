package vish

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/logrusorgru/aurora"
)

func ParseCommand(input string) *Command {
	split := strings.Fields(input)
	if len(split) == 0 {
		return nil
	}
	return NewCommand(split[0], split[1:])
}

func (shell *Vish) ExecCommand(command *Command) (err error) {
	cmd := exec.Command(command.Command, command.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (shell *Vish) manageSignals() {
	signal.Notify(shell.sigChan, syscall.SIGTERM, syscall.SIGINT)
	for {
		shell.manageCommand()
	}
}

func (shell *Vish) manageCommand() {
	defer ignorePanic()
	shell.cmd.Process.Signal(<-shell.sigChan)
}

func ignorePanic() {
	recover()
}

func Abort(err error) {
	if err != nil {
		PrintError(err)
		os.Exit(1)
	}
}

func Alert(err error) {
	if err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	fmt.Println(aurora.Red(err))
}
