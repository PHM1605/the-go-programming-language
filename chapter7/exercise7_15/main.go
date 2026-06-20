package main

import (
	"bufio"
	"exercise7_15/eval"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Store environment for many commands
	env := make(eval.Env)

	fmt.Println("Expression REPL")
	fmt.Println("Type 'q', 'quit', or 'exit' to leave.")
	for {
		fmt.Print("\nexp> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error: ", err)
			return
		}
		// when user types nothing but Enter
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// when user wants to quit
		if line == "q" || line == "quit" || line == "exit" {
			fmt.Println("bye")
			return
		}
		// real parsing
		expr, err := eval.Parse(line)
		if err != nil {
			fmt.Println("Parse error: ", err)
			continue
		}
		// keep track of the vars we must supply
		vars := make(map[eval.Var]bool)
		// check if ALL EXPRESSIONS are correct; fill "vars" in the meantime
		if err := expr.Check(vars); err != nil {
			fmt.Println("Check error: ", err)
			continue
		}
		// Ask for variables we don't know yet
		for v := range vars {
			if _, ok := env[v]; ok {
				continue
			}
			var value float64
			fmt.Printf("%s = ", v)
			if _, err := fmt.Scan(&value); err != nil {
				fmt.Println("Invalid number")
				return
			}
			env[v] = value
		}
		result := expr.Eval(env)
		fmt.Printf("Result = %g\n", result)
	}
}
