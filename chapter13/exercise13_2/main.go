package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// to create  linked list
type Node struct {
	Value int
	Next  *Node
}

// which "address" of which "type" we have checked "cycle"
type visit struct {
	ptr unsafe.Pointer
	typ reflect.Type
}

func hasCycle(v reflect.Value, seen map[visit]bool) bool {
	// if nothing in the variable => no cycle
	if !v.IsValid() {
		return false
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return false
		}
		id := visit{
			ptr: unsafe.Pointer(v.Pointer()),
			typ: v.Type(),
		}
		// if this Node is already in "seen-map"
		if seen[id] {
			return true // cycle
		}
		// update "seen-map"
		seen[id] = true
		return hasCycle(v.Elem(), seen)

	case reflect.Interface:
		if v.IsNil() {
			return false
		}
		return hasCycle(v.Elem(), seen)

	// our Node case above
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if hasCycle(v.Field(i), seen) {
				return true
			}
		}
		return false // no cycle

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if hasCycle(v.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range v.MapKeys() {
			// check each Key in the Map
			if hasCycle(k, seen) {
				return true
			}
			// check each Value in the Map
			if hasCycle(v.MapIndex(k), seen) {
				return true
			}
		}
		return false // no cycle

	default:
		return false // for normal variable => no cycle
	}

}

func HasCycle(x interface{}) bool {
	// a Set to map which Address of which Type we have visited
	// if we visit a place twice => cycle
	seen := make(map[visit]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}

func main() {
	// no cycle
	a := &Node{Value: 1}
	b := &Node{Value: 2}
	c := &Node{Value: 3}
	a.Next = b
	b.Next = c
	fmt.Println(HasCycle(a)) // false

	// cycle
	x := &Node{Value: 1}
	y := &Node{Value: 2}
	z := &Node{Value: 3}
	x.Next = y
	y.Next = z
	z.Next = x
	fmt.Println(HasCycle(x)) // true
}
