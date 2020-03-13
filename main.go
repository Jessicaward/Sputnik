package main

import (
	"Sputnik/printer"
	"Sputnik/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Welcome to the Sputnik programming language!")
	printer.PrintLogo()
	repl.Start(os.Stdin, os.Stdout)
}
