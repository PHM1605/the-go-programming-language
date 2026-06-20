// we need Type that satisfies:
// package flag
//
//	type Value interface {
//		String() string
//		Set(string) error
//	}
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// celsiusFlag must satisfy "Value interface" => must have "String()" and "Set()"
type celsiusFlag struct {
	Celsius // with "String()" inside; which is both f.String() and (&f).String()
}

// <s> will be filled with what follows "-temp" in command line e.g. s = "273.15K"
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // split "273.15K" to 2 variables

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// name: "temp", which is in the command line "-temp"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// NOTE: Var() needs flag.Value() INTERFACE here => we pass <f> which must satisfy <Value> iterface
	flag.CommandLine.Var(&f, name, usage) // <f> String() and Set(s string) functions will be called here; <s> will be filled with what follows "-temp" in command line
	return &f.Celsius
}
