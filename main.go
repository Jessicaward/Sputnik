package main

import (
	"Sputnik/printer"
	"Sputnik/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	printer.Println("Hello %s!\n", user.Username)
	printer.Println("Welcome to the Sputnik programming language!")
	printer.PrintLogo()
	repl.Start(os.Stdin, os.Stdout)
}
