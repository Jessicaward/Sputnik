package repl

import (
	"Sputnik/printer"
	"bufio"
	"fmt"
	"io"
	"os"

	"sputnik/lexer"
	"sputnik/token"
)

//REPL - Read, Eval, Print, Loop
//Essentially this is like an interactive console for Sputnik

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		printer.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "quit" || line == "exit" {
			os.Exit(3)
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			printer.PrintInformation(fmt.Sprintf("%+v\n", tok))
		}
	}
}
