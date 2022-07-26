package shell

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

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
