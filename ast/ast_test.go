package ast

import (
	"interpreter_monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "name"},
					Value: "name",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "nekootoko3"},
					Value: "nekootoko3",
				},
			},
		},
	}

	if program.String() != "let name = nekootoko3;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
