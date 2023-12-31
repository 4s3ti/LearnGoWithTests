package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Error("wrong number of function calls, got", len(got), "want", 1)
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
