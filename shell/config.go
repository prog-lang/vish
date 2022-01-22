package shell

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/logrusorgru/aurora"
)

const (
	Welcome = "Hey there 😘"
	GoodBye = "Bye ❤️"
)

func Prefix() string {
	usr, _ := user.Current()
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s in %s %s ",
		aurora.Bold(aurora.Blue(usr.Username)),
		aurora.Bold(aurora.Green(path.Base(cwd))),
		aurora.Green("➜"))
}
