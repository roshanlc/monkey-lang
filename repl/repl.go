package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/roshanlc/interpreter-in-go/lexer"
	"github.com/roshanlc/interpreter-in-go/token"
)

// REPL = Read Evaluate Print Loop

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Exit codes
		if line == ":quit" || line == ":exit" {
			fmt.Fprint(out, "Bye!!!\n")
			return
		}
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n\n", tok)
		}
	}
}
