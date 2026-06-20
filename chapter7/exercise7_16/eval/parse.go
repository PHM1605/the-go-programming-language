package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type parser struct {
	// scan.Scan() => fetch 1 more token
	// scan.TokenText() => get that token out
	scan scanner.Scanner
	tok  rune
}

// every time we call "p.next()" we scan 1 more token from string
func (p *parser) next() {
	p.tok = p.scan.Scan() // p.tok = "pow" (type scanner.Ident)
}

// input: pow(x,3)+pow(y,3)
func Parse(input string) (_ Expr, err error) {
	// to make <panic> not crash
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	p := &parser{}
	p.scan.Init(strings.NewReader(input))
	p.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	p.next() // now "tok"="pow"

	expr := p.parseExpr()

	// if not reach the end
	if p.tok != scanner.EOF {
		return nil, fmt.Errorf("unexpected %q", p.scan.TokenText())
	}
	return expr, nil
}

func (p *parser) parseExpr() Expr {
	return p.parseAddSub()
}

func (p *parser) parseAddSub() Expr {
	left := p.parseMulDiv()
	for p.tok == '+' || p.tok == '-' {
		op := p.tok
		p.next() // p.tok = "pow"
		right := p.parseMulDiv()
		left = binary{
			op: op,
			x:  left,
			y:  right,
		}
	}
	return left
}

func (p *parser) parseMulDiv() Expr {
	left := p.parseUnary()
	for p.tok == '*' || p.tok == '/' {
		op := p.tok
		p.next()
		right := p.parseUnary()
		left = binary{
			op: op,
			x:  left,
			y:  right,
		}
	}
	return left
}

func (p *parser) parseUnary() Expr {
	if p.tok == '+' || p.tok == '-' {
		op := p.tok
		p.next() // p.tok = "x"
		return unary{
			op: op,
			x:  p.parseUnary(),
		}
	}
	return p.parsePrimary()
}

func (p *parser) parsePrimary() Expr {
	switch p.tok {
	case scanner.Int, scanner.Float:
		v, _ := strconv.ParseFloat(p.scan.TokenText(), 64)
		p.next()
		return literal(v)

	// "pow", "sin" etc. is of type "scanner.Ident"
	case scanner.Ident:
		name := p.scan.TokenText() // "pow"
		p.next()                   // fetch "("
		if p.tok != '(' {
			return Var(name) // "x"
		}
		p.next() // fetch "x" => p.tok = "x"
		var args []Expr
		if p.tok != ')' {
			for {
				args = append(args, p.parseExpr()) // args = [Var("x")], [Var("x"),Literal("3")] etc.
				if p.tok != ',' {                  // "," then ")"
					break
				}
				p.next() // p.tok = "3"
			}
		}
		if p.tok != ')' {
			panic("missing")
		}
		p.next() // p.tok = "+"

		return call{fn: name, args: args} // left side

	case '(':
		p.next()
		e := p.parseExpr()
		if p.tok != ')' {
			panic("missing )")
		}
		p.next() // to EOF
		return e
	}

	panic(fmt.Sprintf("unexpected token %q", p.scan.TokenText()))
}
