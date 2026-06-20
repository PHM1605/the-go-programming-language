package main

import (
	"fmt"

	"eval_proj/eval"
)

func main() {
	tests := []string{
		"x % 2",
		"math.Pi",
		"!true",
		`"hello"`,
		"log(10)",
		"sqrt(1,2)",
	}
	for _, test := range tests {
		fmt.Printf("%s\n", test)
		expr, err := eval.Parse(test)
		if err != nil {
			fmt.Printf("\t%s\n", err)
			continue
		}

		// Check the <expr>True just parsed
		vars := make(map[eval.Var]bool)
		if err := expr.Check(vars); err != nil {
			fmt.Printf("\t%s\n", err)
			continue
		}
	}

	// Other parts of main() after check....
	// ...
}
