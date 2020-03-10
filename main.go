package main

import (
	"Sputnik/repl"
	"fmt"
	"os"
	"os/user"
)

const ASCIILOGO = `
                                                                    
                                                      88 88         
                                    ,d                "" 88         
                                    88                   88         
,adPPYba, 8b,dPPYba,  88       88 MM88MMM 8b,dPPYba,  88 88   ,d8   
I8[    "" 88P'    "8a 88       88   88    88P'   '"8a 88 88 ,a8"    
 '"Y8ba,  88       d8 88       88   88    88       88 88 8888[      
aa    ]8I 88b,   ,a8" "8a,   ,a88   88,   88       88 88 88'"Yba,   
 '"YbbdP"' 88'YbbdP"'   'YbbdP'Y8   "Y888 88       88 88 88   'Y8a  
          88                                                        
          88
`

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Welcome to the Sputnik programming language!")
	fmt.Printf(ASCIILOGO)
	repl.Start(os.Stdin, os.Stdout)
}
