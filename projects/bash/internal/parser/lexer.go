package parser

import (
	"errors"
	"strings"
	"unicode"
)

const (
	word               = -1
	singleQuotedString = -2
	doubleQuotedString = -3
	dollarSign         = -4
)

type token struct {
	val string
	typ int
}

func Lex(s string) ([]token, error) {
	var res []token
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\'':
			q := strings.IndexByte(s[i+1:], '\'')
			if q == -1 {
				return nil, errors.New("single quote isn't closed")
			}
			res = append(res, token{
				val: s[i+1 : i+q],
				typ: singleQuotedString,
			})
		case '"':
			q := strings.IndexByte(s[i+1:], '"')
			if q == -1 {
				return nil, errors.New("double quote isn't closed")
			}
			res = append(res, token{
				val: s[i+1 : i+q],
				typ: doubleQuotedString,
			})
		case '$':
			res = append(res, token{val: s[i : i+1], typ: dollarSign})
		default:
			for ; unicode.IsSpace(rune(s[i])); i++ {
			}
			res = append(res, token{val: s[:i], typ: word})
		}
	}
	return res, nil
}
