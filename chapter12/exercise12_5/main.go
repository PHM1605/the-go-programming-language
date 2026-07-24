package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

// NEW: helper to print indent
func writeIndent(buf *bytes.Buffer, n int) {
	for i := 0; i < n; i++ {
		buf.WriteByte(' ')
	}
}

// NEW: add "indent" information
func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "%g+%gi", real(c), imag(c))

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		if v.IsNil() {
			buf.WriteString("null")
		} else {
			return encode(buf, v.Elem(), indent) // "reflect.Value" of "Interface" or that pointed element
		}

		// ["xxx", "yyy", ...]
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		// insert newline for 1st element
		if v.Len() > 0 {
			buf.WriteByte('\n')
		}
		for i := 0; i < v.Len(); i++ {
			writeIndent(buf, indent+1)
			// "reflect.Value" of each element in that Slice
			if err := encode(buf, v.Index(i), indent+1); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteString(",\n")
			}
		}
		if v.Len() > 0 {
			buf.WriteByte('\n')
			writeIndent(buf, indent)
		}
		buf.WriteByte(']')

	// {"Title":"Dr. Strangelove", "xx":"yy"}...
	case reflect.Struct:
		buf.WriteByte('{')
		// i = 0,1,..,5
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
			}
			writeIndent(buf, indent+1)
			// Field(i): returns "reflect.Value" of the "value" of that Struct entry
			fmt.Fprintf(buf, "%q: ", v.Type().Field(i).Name)          // {"Title":
			if err := encode(buf, v.Field(i), indent+1); err != nil { // {"Title": "Dr. Strangelove"
				return err
			}
		}
		// last element: newline only
		if v.NumField() > 0 {
			buf.WriteByte('\n')
			writeIndent(buf, indent)
		}
		// {"Title": "Dr. Strangelove", "Subtitle": "xxx", ...}
		buf.WriteByte('}')

	// {"Dr. Strangelove": "Peter Sellers", {"xxx" "yyy"} ...}
	case reflect.Map:
		buf.WriteByte('{')
		keys := v.MapKeys()

		for i, key := range keys { // key: reflect.Value of "Dr. Strangelove"
			if i > 0 {
				buf.WriteString(",\n")
			}
			writeIndent(buf, indent+1)
			// {"Dr. Strangelove"
			fmt.Fprintf(buf, "%q: ", key.String())

			// {"Dr. Strangelove": "Peter Sellers"
			if err := encode(buf, v.MapIndex(key), indent+1); err != nil {
				return err
			}
		}
		// last element is handled separately
		if len(keys) > 0 {
			buf.WriteByte('\n')
			writeIndent(buf, indent)
		}
		buf.WriteByte('}')

	// v is "reflect.Value" of an Interface
	// v.Elem() returns its dynamic value
	// v.Elem().Type() returns its dynamic type
	case reflect.Interface: // (type value)
		if v.IsNil() {
			buf.WriteString("nil")
		} else {
			buf.WriteByte('(')
			// dynamic type
			fmt.Fprintf(buf, "%q ", v.Elem().Type().String())
			// dynamic value
			if err := encode(buf, v.Elem(), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
		}

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	// all good
	return nil
}

// Marshal: encodes a Go value in S-expression form
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	var strangelove = Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	data, err := Marshal(strangelove)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
