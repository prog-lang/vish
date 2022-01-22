package shell

import (
	"os"
)

func changeWdToHomeDir() (err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	return os.Chdir(home)
}
