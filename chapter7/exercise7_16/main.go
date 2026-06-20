package main

import (
	"exercise7_16/eval"
	"fmt"
	"html"
	"net/http"
	"strconv"
)

func calculator(w http.ResponseWriter, r *http.Request) {
	// When press "Evaluate", frontend will send GET request to "<address stored in Request>/?expr=x+2?x=4"
	fmt.Fprintln(w, `
		<html>
			<body>
				<h2>Calculator</h2>
				
				<form method="GET">
					Expression:
					<input name="expr" size="40">
					<br><br>
					
					x:
					<input name="x">
					
					y:
					<input name="y">
					
					A:
					<input name="A">
					
					pi:
					<input name="pi">
					<br><br>
					
					<input type="submit" value="Evaluate">
				</form>
	`)

	exprText := r.FormValue("expr")
	if exprText != "" {
		expr, err := eval.Parse(exprText)
		if err != nil {
			// display a <p> element of error message
			fmt.Fprintf(w, "<p>Parse error: %s</p>", html.EscapeString(err.Error())) // EscapeString means < or > is interpreted as those chars
			fmt.Fprintln(w, "</body></html>")
			return
		}
		// store variables in the inserted Expression (need to be supported by Form)
		var vars map[eval.Var]bool = make(map[eval.Var]bool)
		// Check error in Expressions; fill "vars" at the same time
		if err := expr.Check(vars); err != nil {
			fmt.Fprintf(w, "<p>Check error: %s</p>", html.EscapeString(err.Error()))
			fmt.Fprintf(w, "</body></html>")
			return
		}

		// environment variable
		env := make(eval.Env)
		// try getting variables in Expression from Request query (which are also what being filled in Form)
		// => update Environment dict
		for v := range vars {
			text := r.FormValue(string(v))
			if text == "" {
				continue
			}
			value, err := strconv.ParseFloat(text, 64)
			if err != nil {
				fmt.Fprintf(w, "<p>Bad value for %s</p>", v)
				fmt.Fprintln(w, "</body></html>")
				return
			}
			env[v] = value
		}
		// Calculate result and display
		result := expr.Eval(env)
		fmt.Fprintf(w, "<h3>Result = %g</h3>", result)
	}
	fmt.Fprintf(w, `
		</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", calculator)
	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
