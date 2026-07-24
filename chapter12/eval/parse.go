package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type parser struct {
	// scan.Scan() => return type of the next token e.g. scanner.Ident or '(' or scanner.Int etc.
	// scan.TokenText() => show the string represents that token
	scan scanner.Scanner
	tok  rune
}

// every time we call p.next() we scan 1 more token from string
func (p *parser) next() {
	p.tok = p.scan.Scan() // p.tok = scanner.Ident
}

func (p *parser) parseExpr() Expr {
	return p.parseAddSub()
}

func (p *parser) parseAddSub() Expr {
	left := p.parseMulDiv()
	for p.tok == '+' || p.tok == '-' {
		op := p.tok
		p.next() // p.tok='pow' (right side)
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
	// "pow", "sin", "x" etc.
	case scanner.Ident:
		name := p.scan.TokenText() // "pow"
		p.next()                   // p.tok='('
		if p.tok != '(' {          // "x"
			return Var(name)
		}
		p.next() // fetch "x"
		var args []Expr
		if p.tok != ')' {
			for {
				args = append(args, p.parseExpr()) // [Var(x)]; [Var(x), Literal(3)]
				if p.tok != ',' {                  // ')'
					break
				}
				// if p.tok=',' then here we get '3'
				p.next()
			}
		}
		if p.tok != ')' {
			panic("missing")
		}
		p.next() // p.tok = '+'

		return call{fn: name, args: args} // left side; return pow(x,3)

	case '(':
		p.next()
		e := p.parseExpr()
		if p.tok != ')' {
			panic("missing )")
		}
		p.next() // to EOF
		return e
	}

	panic(fmt.Sprintf("unexpected '%s'", p.scan.TokenText()))
}

// e.g. input: pow(x,3)+pow(y,3) => return "binary Expr"
func Parse(input string) (_ Expr, err error) {
	// to make "panic" not crash
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("%v", x)
		}
	}()

	p := &parser{}
	p.scan.Init(strings.NewReader(input))
	p.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	p.next() // now "tok" = "pow"

	expr := p.parseExpr()

	// if not reach the end
	if p.tok != scanner.EOF {
		return nil, fmt.Errorf("unexpected '%s'", p.scan.TokenText())
	}
	return expr, nil
}
