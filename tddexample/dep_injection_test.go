package tddexample

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kenny")

	got := buffer.String()
	want := "Hello Kenny"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
