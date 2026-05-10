// Print command line arguments with package "flag"
// $ go build echo4.go
// $ ./echo4 -s / a bc def => print "a/bc/def\n"
// $ ./echo4 -n a bc def => print "a bc def" (NO \n)
// $ ./echo4 -help => print message (3rd argument in flag.Bool or flag.String)
package main

import (
	"flag"
	"fmt"
	"strings"
)

// create a POINTER VARiable of type "bool", flag name "-n", default value "false", message "xxx" when error or "-help"
var n = flag.Bool("n", false, "omit trailing newline")

// create a POINTER VARiable of type "bool", flag name "-s", default value "false", message "xxx" when error or "-help"
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // update "n" and "sep"
	// flag.Args() => non-flag variables
	// Join non-flag variables in command-line with separator "*sep"
	fmt.Print(strings.Join(flag.Args(), *sep))
	// if "-n" not found in command line => print newline
	if !*n {
		fmt.Println()
	}
}
