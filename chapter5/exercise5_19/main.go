package main

import "fmt"

// Note: this function doesn't have a "return" statement
// but still, "result" has a value
// => trick: cause a "panic" then create "defer-recover"

func panicFunction() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = 42 // we change returned value here
		}
	}()

	panic("boom")
}

func main() {
	fmt.Println(panicFunction())
}
