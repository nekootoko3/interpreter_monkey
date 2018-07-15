package repl

import (
	"bufio"
	"fmt"
	"interpreter_monkey/aa"
	"interpreter_monkey/evaluator"
	"interpreter_monkey/lexer"
	"interpreter_monkey/object"
	"interpreter_monkey/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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

		evaluated := evaluator.Eval(program, env)

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
