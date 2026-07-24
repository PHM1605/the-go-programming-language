package eval

import (
	"fmt"
	"math"
	"strings"
)

// 5 types of expressions
type Var string      // variable like "x"
type literal float64 // float like "3.12"
type unary struct {
	op rune // + -
	x  Expr
}
type binary struct {
	op   rune // one of "+", "-", "*", "/"
	x, y Expr
}
type call struct {
	fn   string // function call "pow", "sin", "sqrt"
	args []Expr
}

// to evaluate a Var to its value => map from "string" to "float"
type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64 // convert Var to its value
	// report errors in this Expr and add its Vars to the Set
	Check(vars map[Var]bool) error
}

// add method "Eval()" to make all types satisfies Expr
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

// Add method Check() to make all types of expressions satisfying Expr interface
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}
func (literal) Check(vars map[Var]bool) error {
	return nil // not defined
}

// unary has an "op" and an "Expr"
func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars) // "vars" Set will be updated
}

// binary has "op" and 2 "Expr"
func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil { // "vars" Set will be updated inside
		return err
	}
	return b.y.Check(vars) // "vars" Set will be updated inside
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), arity)
	}
	// actual check Variables
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
