package parser

import (
	"fmt"
)

type Command struct {
	Command  string
	Argument string
}

func ParseCommandLine(s string) ([]Command, error) {
	var res []Command
	tokens, err := Lex(s)
	if err != nil {
		return nil, fmt.Errorf("lexer error: %w", err)
	}
	_ = tokens
	return res, nil
}
