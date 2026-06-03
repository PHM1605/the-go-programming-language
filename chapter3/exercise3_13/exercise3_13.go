// Write const declaration for KB, MB up through YB as compactly as you can
// go run exercise3_13.go
package main

import "fmt"

// NOTE: we can't use int64 as it's over-ranged
const (
	KB float64 = 1000
	MB         = KB * 1000
	GB         = MB * 1000
	TB         = GB * 1000
	PB         = TB * 1000
	EB         = PB * 1000
	ZB         = EB * 1000
	YB         = ZB * 1000
)

func main() {
	fmt.Printf("KB = %.0f\n", KB)
	fmt.Printf("MB = %.0f\n", MB)
	fmt.Printf("GB = %.0f\n", GB)
	fmt.Printf("TB = %.0f\n", TB)
	fmt.Printf("PB = %.0f\n", PB)
	fmt.Printf("EB = %.0f\n", EB)
	fmt.Printf("ZB = %.0f\n", ZB)
	fmt.Printf("YB = %.0f\n", YB)
}
