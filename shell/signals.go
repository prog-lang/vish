package shell

import (
	"os/signal"
	"syscall"
)

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
