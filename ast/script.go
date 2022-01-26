package ast

import (
	"os"
	"os/exec"
)

type Script struct {
	Name string
	Args []string
}

func NewScript(name string, args []string) *Script {
	return &Script{
		Name: name,
		Args: args,
	}
}

func (script *Script) Exec() (err error) {
	cmd := exec.Command(script.Name, script.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
