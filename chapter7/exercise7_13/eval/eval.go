package eval

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64 // convert Var to its value
	// report errors in this Expr and adds its Vars to the Set
	Check(vars map[Var]bool) error
	// NEW
	String() string
}

// 5 kinds of expressions
type Var string      // variable like "x"
type literal float64 // float like "3.21"
type unary struct {
	op rune // + -
	x  Expr
}
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}
type call struct {
	fn   string // function call "pow", "sin", "sqrt"
	args []Expr
}

// to evaluate a Var to value => we need a map from "string" to "float"
type Env map[Var]float64

// implementing the interfaces of Expr
func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("Unsupported binary operator: %q", b.op))
}
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("Unsupported function call: %s", c.fn))
}

// Implement the Check() method for the Interface
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}
func (literal) Check(vars map[Var]bool) error {
	return nil
}

// <unary> has an <op> and an <Expr>
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars) // <vars> Set will be updated
}

// <binary> has an <op> and 2 <Expr>
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars) // <vars> will be updated inside
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	// actual check variables
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

// Implement for Interface of Expr
func (v Var) String() string {
	return string(v) // "x"
}
func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l)) // 3.14 => "3.14"
}
func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String()) // "(+x)"
}
func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String()) // "(a + 3)"
}
func (c call) String() string {
	var args []string
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", ")) // "pow(x, 4)"
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
