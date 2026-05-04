// Echo prints command-line arguments with "index" and "value"
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b strings.Builder
	var res string
	for i, arg := range os.Args[1:] {
		b.WriteString(strconv.Itoa(i) + "/ " + arg + "\n")
	}
	res = b.String()
	fmt.Println(res)
}
