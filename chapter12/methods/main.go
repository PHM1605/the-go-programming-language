package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// x is a variable of type "time.Duration"
func Print(x interface{}) {
	// get both dynamic value & dynamic type of "x"
	v := reflect.ValueOf(x)
	t := v.Type() // dynamic type of "x" e.g. "time.Duration"
	fmt.Printf("type %s\n", t)

	// v.Method(i): that method => v.Method(i).Call(xxx) runs that
	// t.Method(i): that method's metadata => .Name() will print it's name
	// t.Method(i).String(): signature-string of the method
	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		// func (time.Duration) Hours() () float64
		fmt.Printf("func (%s) %s %s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	// type time.Duration
	// func (time.Duration) Abs () time.Duration
	// func (time.Duration) Hours () float64
	// func (time.Duration) Microseconds () int64
	// func (time.Duration) Milliseconds () int64
	// func (time.Duration) Minutes () float64
	// func (time.Duration) Nanoseconds () int64
	// func (time.Duration) Round (time.Duration) time.Duration
	// func (time.Duration) Seconds () float64
	// func (time.Duration) String () string
	// func (time.Duration) Truncate (time.Duration) time.Duration
	Print(time.Hour)

	// methods of *strings.Replacer
	Print(new(strings.Replacer))
}
