package vish

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

func ParseCommand(input string) *Command {
	split := strings.Fields(input)
	if len(split) == 0 {
		return nil
	}
	return NewCommand(split[0], split[1:])
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
