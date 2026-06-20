package main

import (
	"exercise7_13/eval"
	"fmt"
)

func main() {
	tests := []string{
		"x",
		"x+y",
		"sqrt(A/pi)",
		"pow(x,3)+pow(y,3)",
		"sin(-x)*pow(1.5,-r)",
	}
	for _, s := range tests {
		expr, err := eval.Parse(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// convert expression tree to string
		text := expr.String()
		// parse back from string to expression
		expr_back, err_back := eval.Parse(text)
		if err_back != nil {
			fmt.Printf("reparse failed: %v\n", err)
			continue
		}

		fmt.Printf("original: %s\n", s)
		fmt.Printf("string: %s\n", text)
		fmt.Printf("reparsed: %v\n", expr_back)
		fmt.Println()
	}
}
