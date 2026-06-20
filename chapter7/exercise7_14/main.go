package main

import (
	"exercise7_14/eval"
	"fmt"
)

// This is Expr interface; calculate "min(x,y)"
type min struct {
	x, y eval.Expr
}

func (m min) Eval(env eval.Env) float64 {
	a := m.x.Eval(env)
	b := m.y.Eval(env)
	if a < b {
		return a
	}
	return b
}

func (m min) Check(vars map[eval.Var]bool) error {
	if err := m.x.Check(vars); err != nil {
		return err
	}
	return m.y.Check(vars)
}

func (m min) String() string {
	return fmt.Sprintf("min(%s, %s)", m.x, m.y)
}

func main() {
	// create a test <min> expression
	expr := min{
		x: eval.Var("x"),
		y: eval.Var("y"),
	}
	env := eval.Env{
		"x": 10,
		"y": 3,
	}
	fmt.Println(expr.Eval(env))
}
