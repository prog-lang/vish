package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		REPL()
	}
}

func REPL() {
	input := Read()
	result, err := Eval(input)
	Abort(err)
	Print(result)
}

func Read() string {
	fmt.Print("VISH ❯❯❯ ")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}

func Eval(input string) (result []byte, err error) {
	split := strings.Split(input, " ")
	result, err = exec.Command(split[0], split[1:]...).Output()
	return
}

func Abort(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func Print(result []byte) {
	fmt.Println(string(result))
}
