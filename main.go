package main

import (
	"Sputnik/repl"
	"Sputnik/utilities"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	utilities.PrintError("test")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Welcome to the Sputnik programming language!")
	utilities.PrintLogo()
	repl.Start(os.Stdin, os.Stdout)
}
