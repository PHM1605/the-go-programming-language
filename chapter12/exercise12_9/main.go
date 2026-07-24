package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/scanner"
)

type Token interface{}
type Symbol string
type String string
type Int int
type StartList struct{}
type EndList struct{}

// lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // type of word
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan() // type of word
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

// Decoder
type Decoder struct {
	lex *lexer // scanner and its holding token
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{
		scan: scanner.Scanner{
			Mode: scanner.GoTokens,
		},
	}
	lex.scan.Init(r)
	lex.next() // get 1st token
	return &Decoder{lex: lex}
}

// NEW: fetch token
func (d *Decoder) Token() (Token, error) {
	switch d.lex.token {
	case scanner.EOF:
		return nil, io.EOF

	case '(':
		d.lex.next()
		return StartList{}, nil

	case ')':
		d.lex.next()
		return EndList{}, nil

	case scanner.Ident: // string with no "" or ''
		sym := Symbol(d.lex.text())
		d.lex.next()
		return sym, nil

	case scanner.String: // string with "" or ''
		s, err := strconv.Unquote(d.lex.text())
		if err != nil {
			return nil, err
		}
		d.lex.next()
		return String(s), nil

	case scanner.Int:
		n, err := strconv.Atoi(d.lex.text())
		if err != nil {
			return nil, err
		}
		d.lex.next()
		return Int(n), nil

	default:
		return nil, fmt.Errorf("unexpected token %q at %s", d.lex.text(), d.lex.scan.Position)
	}
}

func main() {
	input := `
		((Title "Dr. Strangelove")
		(Year 1964)
		(Color nil))
	`
	dec := NewDecoder(strings.NewReader(input))
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T: %v\n", tok, tok)
	}
}
