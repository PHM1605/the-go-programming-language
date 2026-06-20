package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3)+pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x,3)+pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
		{"5/9*(F-32)", Env{"F": -40}, "-40"},
		{"5/9*(F-32)", Env{"F": 32}, "0"},
		{"5/9*(F-32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		// print the original expression of the test only when it's different
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr) // turn a string => expression "binary add" (if e.g. "pow(x,3)+pow(y,3)")
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got) // print environment and output
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
