package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "tree"}, "one\ttwo\ttree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"a", "b", "c"}, "a:b:c"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"}, // wrong expectation
	}

	for _, test := range tests {
		desc := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)
		// "out" is from the "echo.go" file
		out = new(bytes.Buffer)
		// start writing into "out"
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s is failed: %v", desc, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", desc, got, test.want)
		}
	}
}
