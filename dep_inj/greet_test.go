package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	s := strings.Builder{}
	Greet(&s, "Chris")

	got := s.String()
	want := "Porco, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
