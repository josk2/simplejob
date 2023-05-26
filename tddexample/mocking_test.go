package tddexample

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &spySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if want != got {
		t.Errorf("Got %q want %q", got, want)
	}

	if spySleeper.Calls != countdownStart {
		t.Errorf("not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.Calls)
	}

}
