package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// NEW: Cycle struct
type Cycle struct {
	Value int
	Tail  *Cycle
}

const maxDepth = 5

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
func display(path string, v reflect.Value, depth int) {
	if depth > maxDepth {
		fmt.Printf("%s = ...\n", path)
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			// 1st param: "e.args[0]" or "e.args[1]"
			// 2nd param: "reflect.Value" wrapper of interface "Expr" => will match "reflect.Interface" later
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), depth+1)
		}

	case reflect.Struct:
		// v.NumField() returns the number of fields in that struct wrapped by "v"
		// i = 0; 1
		// v.Type().Field(i): metadata of fn-Field/ args-Field
		// v.Field(i): "reflect.Value" wrapper of "fn-Field" and "args-Field"
		for i := 0; i < v.NumField(); i++ {
			// fieldPath: "e.fn"/"e.args"
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), depth+1)
		}

	case reflect.Map:
		// "key" is a "reflect.Value"
		// v.MapIndex(key) returns a "reflect.Value"
		for _, key := range v.MapKeys() {
			// NEW: display the "key" recursively
			display(path+".key", key, depth+1) // m1.key.X = 1; m1.key.Y = 2; m1.key.X = 3; m1.key.Y = 4;
			// display value recursively
			display(path+"[value]", v.MapIndex(key), depth+1) // m1[value] = "A"; m1[value] = "B"
		}

	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// .Elem() returns what "reflect.Value" that pointer is pointing to
			display(fmt.Sprintf("(*%s)", path), v.Elem(), depth+1)
		}

	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			// reflect.Type of this Interface
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			// Elem(): returns "reflect.Value" of this Interface
			display(path+".value", v.Elem(), depth+1)
		}

	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// export "Display"; hide the inner "display" function that does the real work
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

func main() {
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
}
