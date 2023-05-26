package tddexample

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

var (
	countdownStart = 3
	finalWord      = `Go!`
)

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct {
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
		//time.Sleep(1 * time.Second)
	}

	fmt.Fprint(writer, finalWord)
}
