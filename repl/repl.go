package repl

import (
	"bufio"
	"fmt"
	"io"
	
	"sputnik/lexer"
	"sputnik/token"
)

//REPL - Read, Eval, Print, Loop
//Essentially this is like an interactive console for Sputnik

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
