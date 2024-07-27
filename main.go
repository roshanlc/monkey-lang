package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/roshanlc/interpreter-in-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\nType :quit or :exit to exit\n")
	repl.Start(os.Stdin, os.Stdout)
}
