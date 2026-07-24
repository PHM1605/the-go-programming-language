package main

import (
	"chapter12/eval"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

// convert any basic type to its "string-representation"
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)

	// ...floating-point and complex number cases are omitted for brevity
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.String:
		return strconv.Quote(v.String())

	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map: // Slice, Ptr and Map are outside
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)

	default: // Array, Struct, Interface
		return v.Type().String() + " value"
	}
}

// v is wrapper "reflect.Value" over a "call" object
// => .Kind() returns a "reflect.Struct"
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			// 1st param: "e.args[0]" or "e.args[1]"
			// 2nd param: "reflect.Value" wrapper of interface "Expr" => will match "reflect.Interface" later
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}

	case reflect.Struct:
		// v.NumField() returns the number of fields in that struct wrapped by "v"
		// i = 0; 1
		// v.Type().Field(i): metadata of fn-Field/ args-Field
		// v.Field(i): "reflect.Value" wrapper of "fn-Field" and "args-Field"
		for i := 0; i < v.NumField(); i++ {
			// fieldPath: "e.fn"/"e.args"
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}

	case reflect.Map:
		// "key" is a "reflect.Value"
		// v.MapIndex(key) returns a "reflect.Value"
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}

	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// .Elem() returns what "reflect.Value" that pointer is pointing to
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}

	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// reflect.Type of this Interface
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			// Elem(): returns "reflect.Value" of this Interface
			display(path+".value", v.Elem())
		}

	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// export "Display"; hide the inner "display" function that does the real work
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// example to display
var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
}

func main() {
	// "e" is a "call" object that satisfies "eval.Expr"
	// "eval.Expr" is an Interface with "Eval(Env)" and "Check(map[Var]bool)" methods
	e, _ := eval.Parse("sqrt(A / pi)") // convert a string to an Expression named "e"
	Display("e", e)

	// Display movie
	fmt.Println()
	Display("strangelove", strangelove)

	// Display standard library
	fmt.Println()
	Display("os.Stderr", os.Stderr)

	// Display reflect.Value
	fmt.Println()
	Display("rV", reflect.ValueOf(os.Stderr))

	// Difference between "reflect.Value" returns by "reflect.ValueOf()" and ".Elem()"
	var i interface{} = 3
	fmt.Println()
	Display("i", i)   // returns an "int"
	Display("&i", &i) // return Interface via the ".Elem()" inside
}
