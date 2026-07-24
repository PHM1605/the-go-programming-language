package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// format any input value into "string"
func Any(value interface{}) string {
	// reflect.ValueOf(xxx): wrapping "interface{}" to "reflect.Value" (which has Kind() method)
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	// "reflect.Value" has Kind() method
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

	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		// reflect.Value calls "Pointer()" method that returns "reference-address-of-chan" for example
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)

	default: // reflect.Array, reflect.Struct, reflect.Interface
		// convert from "reflect.Value" to "reflect.Type" first; then use String() method
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}
