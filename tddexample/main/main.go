package main

import (
	"os"

	"simplejob/tddexample"
)

func main() {
	sleeper := tddexample.DefaultSleeper{}
	tddexample.Countdown(os.Stdout, sleeper)
}
