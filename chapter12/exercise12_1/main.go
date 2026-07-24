package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// NEW: key of a "map" can be a "struct"
type Point struct {
	X, Y int
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
			// NEW: display the "key" recursively
			display(path+".key", key) // m1.key.X = 1; m1.key.Y = 2; m1.key.X = 3; m1.key.Y = 4;
			// display value recursively
			display(path+"[value]", v.MapIndex(key)) // m1[value] = "A"; m1[value] = "B"
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

func main() {
	// keys are "struct"s
	m1 := map[Point]string{
		{1, 2}: "A",
		{3, 4}: "B",
	}
	// keys are "array"s
	m2 := map[[2]int]string{
		{5, 6}: "C",
		{7, 8}: "D",
	}

	Display("m1", m1)
	fmt.Println()
	Display("m2", m2)
}
