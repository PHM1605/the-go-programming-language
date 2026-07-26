package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		// do something
	}
}
