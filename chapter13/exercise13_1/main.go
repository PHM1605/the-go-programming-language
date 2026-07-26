package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

// NEW: 1_000_000_000 and 1_000_000_001 are considered equal
const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	if a == b {
		return true
	}
	// two very very closed numbers are also equal
	diff := math.Abs(a - b)
	if a == 0 || b == 0 {
		return diff < epsilon
	}
	return diff/math.Max(math.Abs(a), math.Abs(b)) < epsilon
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// we must use "reflect.Value" to allow "recursion of argument"
// i.e. to compare two "structs" as equal, we must be aware that "value" in {key:"value"} can also be "struct"
// seen: a "SET" that holds which 2 variables (addresses) that we have compared AND their type
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	// IsValid() is false if it's "nil" or "empty"
	// if both "nil"/"empty" => return "true"
	// if 1 of them "nil"/"empty" and the other is not => return "false"
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	// if both has value => compare their "Types"
	if x.Type() != y.Type() {
		return false
	}

	// NOTE: cycle check here; if this comparison has been done => terminate
	// we need to cache RAW ADDRESS of x & y to see if this comparison has been done
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // we are comparing 1 same variable
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] { // if both addresses and their type exist in the Set
			return true // => they are same, otherwise we have exited with "false" early in outer "equal" recursions
		}
		seen[c] = true
	}

	// if surely both has same type
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// NEW: add "almostEqual()" for ALL numeric cases
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return almostEqual(float64(x.Int()), float64(y.Int()))

	case reflect.Float32, reflect.Float64:
		return almostEqual(x.Float(), y.Float())

	case reflect.Complex64, reflect.Complex128:
		xc, yc := x.Complex(), y.Complex()
		return almostEqual(real(xc), real(yc)) && almostEqual(imag(xc), imag(yc))

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	// recursion check here
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		// empty map and "nil" map are considered equal
		// NOTE: we can NOT use x.IsValid() because "nil-map" is considered VALID
		if x.Len() == 0 && y.Len() == 0 {
			return true
		}
		// if 1 map is "nil" and the other has value
		if x.IsNil() != y.IsNil() {
			return false
		}
		// if 1 map has more fields than the other
		if x.Len() != y.Len() {
			return false
		}
		// loop over map
		for _, key := range x.MapKeys() {
			// value of that field
			xv := x.MapIndex(key)
			yv := y.MapIndex(key)
			// if key doesn't exist in map
			if !yv.IsValid() {
				return false
			}
			// recursive here
			if !equal(xv, yv, seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// NOTE: we don't expose "reflect.Value" in the API => we can't use "equal()" directly but wrapping here
func Equal(x, y interface{}) bool {
	// seen: a "SET" that holds which 2 variables (addresses) that we have compared AND their type
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func main() {
	fmt.Println(Equal(1000000000, 1000000001)) // "true"
	fmt.Println(Equal(1000000000, 1000000011)) // "false"
}
