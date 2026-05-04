// Echo3 prints command-line arguments using "strings" package
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// for debugging purpose => printing [Hello World]
	// fmt.Println(os.Args[1:], " ")
}
