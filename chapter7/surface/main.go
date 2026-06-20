package main

import (
	"fmt"
	"math"
	"net/http"
	"surface/eval"
)

// convert "string" to "Expression tree"; checking error in the meantime
func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	// now we have the list of vars we must support
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	// after this r.Form will contain a "map[string][]string" like
	// {"expr": {"sin(-x)*pow(1.5,-r)"}}
	r.ParseForm()

	// NEW: we check for input expression (what User types on URL) first before running it for 10000 points
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)                              // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r}) // Evaluate a function with it's input variables
	})
}

// http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
func main() {
	http.HandleFunc("/", plot)
	fmt.Println("Listen on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
