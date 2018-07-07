package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/aa"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

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

		// print token ------------------------------------
		//		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		//			fmt.Printf("%+v\n", tok)
		//		}
		// ---------------------------------------------

		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	aa := aa.New()
	for _, msg := range errors {
		io.WriteString(out, aa.PrintAa())
		io.WriteString(out, "\t"+msg+"\n")
	}
}
