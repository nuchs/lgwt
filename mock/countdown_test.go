package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	ss := &SpySleeper{}

	Countdown(buffer, ss)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if ss.Calls != 3 {
		t.Errorf("Not enough calls to sleeper, want 3 got %d", ss.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
