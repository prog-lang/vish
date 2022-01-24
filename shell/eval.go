package shell

import (
	"os"
	"os/exec"

	"github.com/sharpvik/vish/ast"
)

func Eval(astree ast.AST) (err error) {
	for _, cmd := range astree {
		if err = EvalCommand(cmd); err != nil {
			return
		}
	}
	return
}

func EvalCommand(command *ast.Command) (err error) {
	if ran, err := RunSpecialCommand(command); ran {
		return err
	}
	return ExecCommand(command)
}

func ExecCommand(command *ast.Command) (err error) {
	cmd := exec.Command(command.Command, command.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
