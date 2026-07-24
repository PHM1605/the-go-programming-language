package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type Demo struct {
	B bool
	F float64
	C complex128
	I interface{}
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%g %g)", real(c), imag(c))

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem()) // "reflect.Value" of "Interface" or that pointed element

		// ("xxx" "yyy" ...)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			// "reflect.Value" of each element in that Slice
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	// ((Title "Dr. Strangelove")...)
	case reflect.Struct:
		buf.WriteByte('(')
		// i = 0,1,..,5
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			// Field(i): returns "reflect.Value" of the "value" of that Struct entry
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name) // ((Title
			if err := encode(buf, v.Field(i)); err != nil {  // ((Title "Dr. Strangelove"
				return err
			}
			buf.WriteByte(')') //((Title "Dr. Strangelove")
		}
		// ((Title "Dr. Strangelove") (Subtitle "xxx"), ...)
		buf.WriteByte(')')

	// (("Dr. Strangelove" "Peter Sellers") ("xxx" "yyy") ...)
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() { // key: reflect.Value of "Dr. Strangelove"
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			// (("Dr. Strangelove"
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			// (("Dr. Strangelove" "Peter Sellers"
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

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
			if err := encode(buf, v.Elem()); err != nil {
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
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	var x interface{} = []int{1, 2, 3} // an Interface has dynamic-type="[]int" and dynamic-value={1,2,3}
	// a Struct with various types to check
	d := Demo{
		B: true,
		F: 3.14,
		C: complex(1, 2),
		I: x,
	}

	data, err := Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
