package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

const (
	finalWord      = "Go!"
	countDownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}
